package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"metabot/internal/autocrat_v0"
	"metabot/internal/openbook_twap"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"

	bin "github.com/gagliardetto/binary"

	"github.com/bwmarrin/discordgo"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/joho/godotenv"
)

var (
	client *rpc.Client
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	Token := os.Getenv("DISCORD_TOKEN")
	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	fmt.Println("Bot is now running. Press CTRL+C to exit.")
	go updateChannelNames(dg)
	select {} // Keep the program running until CTRL+C is pressed.
}

func updateChannelNames(s *discordgo.Session) {
	updateAllGuilds(s)

	ticker := time.NewTicker(time.Minute)
	for {
		select {
		case <-ticker.C:
			updateAllGuilds(s)
		}
	}
}

// updateAllGuilds encapsulates the logic to update all guilds, making it reusable
func updateAllGuilds(s *discordgo.Session) {
	guilds, err := s.UserGuilds(0, "", "")
	if err != nil {
		fmt.Println("Error fetching guilds:", err)
		return
	}

	for _, guild := range guilds {
		updateChannelsInGuild(s, guild.ID)
	}
}

func updateChannelsInGuild(s *discordgo.Session, guildID string) {
	channels, err := s.GuildChannels(guildID)
	if err != nil {
		fmt.Println("Error fetching channels for guild:", err)
		return
	}

	var metabotCategoryID string
	for _, channel := range channels {
		if channel.Type == discordgo.ChannelTypeGuildCategory && channel.Name == "METABOT" {
			metabotCategoryID = channel.ID
			break
		}
	}

	if metabotCategoryID == "" {
		return // "metabot" category not found
	}

	proposals := GetProposals()
	existingChannels := make(map[int]bool) // Track existing channels by proposal ID

	for _, channel := range channels {
		if channel.ParentID == metabotCategoryID && channel.Name != "METABOT" {
			proposalIDWithPrefix := strings.Split(channel.Name, ":")[0]
			proposalID, err := strconv.Atoi(strings.TrimPrefix(proposalIDWithPrefix, "P"))
			if err != nil {
				continue // Skip if the channel name does not conform to expected format
			}

			if newName, ok := proposals[proposalID]; ok {
				// Update the channel if it exists in proposals
				_, err := s.ChannelEdit(channel.ID, &discordgo.ChannelEdit{Name: newName})
				if err != nil {
					fmt.Println("Error updating channel name:", err)
				}
				existingChannels[proposalID] = true
			} else {
				// Delete the channel if it does not exist in proposals
				_, err := s.ChannelDelete(channel.ID)
				if err != nil {
					fmt.Println("Error deleting channel:", err)
				}
			}
		}
	}

	// Create channels for proposals that do not have an existing channel
	for proposalID, name := range proposals {
		if !existingChannels[proposalID] {
			_, err := s.GuildChannelCreate(guildID, name, discordgo.ChannelTypeGuildVoice)
			if err != nil {
				fmt.Println("Error creating channel:", err)
			}
		}
	}

	reorderChannels(s, guildID, metabotCategoryID, proposals)
}

func reorderChannels(s *discordgo.Session, guildID string, metabotCategoryID string, proposals map[int]string) {
	channels, err := s.GuildChannels(guildID)
	if err != nil {
		fmt.Println("Error fetching channels for guild:", err)
		return
	}

	var sortedProposalIDs []int
	for id := range proposals {
		sortedProposalIDs = append(sortedProposalIDs, id)
	}
	sort.Ints(sortedProposalIDs)

	// Filter channels by the metabotCategoryID and sort them by position
	var categoryChannels []*discordgo.Channel
	for _, channel := range channels {
		if channel.ParentID == metabotCategoryID {
			categoryChannels = append(categoryChannels, channel)
		}
	}
	sort.Slice(categoryChannels, func(i, j int) bool {
		return categoryChannels[i].Position < categoryChannels[j].Position
	})

	// Assuming sortedProposalIDs is already populated as shown above
	for index, proposalID := range sortedProposalIDs {
		for _, channel := range categoryChannels {
			proposalIDWithPrefix := strings.Split(channel.Name, ":")[0]
			id, err := strconv.Atoi(strings.TrimPrefix(proposalIDWithPrefix, "P"))
			if err != nil {
				continue
			}
			if id == proposalID {
				_, err := s.ChannelEditComplex(channel.ID, &discordgo.ChannelEdit{
					Position: index,
				})
				if err != nil {
					fmt.Println("Error reordering channel:", err)
				}
				break
			}
		}
	}
}

func GetProposals() (proposals map[int]string) {
	endpoint := os.Getenv("RPC_URL")
	client = rpc.New(endpoint)
	proposals = make(map[int]string)

	filter := rpc.RPCFilter{
		Memcmp: &rpc.RPCFilterMemcmp{Bytes: solana.Base58(autocrat_v0.ProposalDiscriminator[:]), Offset: 0},
	}

	out, err := client.GetProgramAccountsWithOpts(
		context.TODO(),
		solana.MustPublicKeyFromBase58("metaX99LHn3A7Gr7VAcCfXhpfocvpMpqQ3eyp3PGUUq"),
		&rpc.GetProgramAccountsOpts{
			Filters: []rpc.RPCFilter{filter},
		},
	)
	if err != nil {
		panic(err)
	}

	for _, proposal := range out {
		var prop autocrat_v0.Proposal
		// Account{}.Data.GetBinary() returns the *decoded* binary data
		// regardless the original encoding (it can handle them all).
		err = bin.NewBorshDecoder(proposal.Account.Data.GetBinary()).Decode(&prop)
		if err != nil {
			panic(err)
		}
		if prop.State != autocrat_v0.ProposalStatePending {
			continue
		}

		details := GetTWAP(prop)
		if details.Pass > details.Fail {
			proposals[int(prop.Number)] = fmt.Sprintf("P%d: $%.2fp > $%.2ff \n", prop.Number, details.Pass, details.Fail)
		} else {
			proposals[int(prop.Number)] = fmt.Sprintf("P%d: $%.2fp < $%.2ff \n", prop.Number, details.Pass, details.Fail)
		}
	}

	return proposals
}

func GetTWAP(proposal autocrat_v0.Proposal) TWAP {
	passMarket := GetTWAPMarket(proposal.OpenbookTwapPassMarket)
	failMarket := GetTWAPMarket(proposal.OpenbookTwapFailMarket)

	return TWAP{
		Pass:       GetDollarValue(passMarket.TwapOracle),
		Fail:       GetDollarValue(failMarket.TwapOracle),
		HiddenPass: GetHiddenDollarValue(passMarket.TwapOracle),
		HiddenFail: GetHiddenDollarValue(failMarket.TwapOracle),
	}
}

func GetTWAPMarket(market solana.PublicKey) openbook_twap.TWAPMarket {
	twap, err := client.GetAccountInfo(
		context.TODO(),
		market)

	if err != nil {
		panic(err)
	}

	var twapMarket openbook_twap.TWAPMarket
	err = bin.NewBorshDecoder(twap.GetBinary()).Decode(&twapMarket)

	if err != nil {
		panic(err)
	}

	return twapMarket
}

func GetDollarValue(oracle openbook_twap.TWAPOracle) float64 {

	slotsPassed := oracle.LastUpdatedSlot - oracle.InitialSlot
	slotsPassedBig := new(big.Int).SetUint64(slotsPassed)
	twapValue := new(big.Int).Div(oracle.ObservationAggregator.BigInt(), slotsPassedBig)

	twapValueFloat := new(big.Float).SetInt(twapValue)
	divisor := new(big.Float).SetFloat64(10000.0) // 4 decimal places means dividing by 10000
	// Divide to get the actual dollar value
	dollarValue, _ := new(big.Float).Quo(twapValueFloat, divisor).Float64()

	return dollarValue

}

func GetHiddenDollarValue(oracle openbook_twap.TWAPOracle) float64 {

	dollarValue := float64(oracle.LastObservation) / 10000.0

	return dollarValue
}

type TWAP struct {
	Pass       float64
	HiddenPass float64
	Fail       float64
	HiddenFail float64
}
