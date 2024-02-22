package discord

import (
	"fmt"
	"metabot/internal/jup"
	"metabot/internal/metadao"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/dustin/go-humanize"
)

func UpdateChannelNames(s *discordgo.Session) {
	updateAllGuilds(s)

	ticker := time.NewTicker(time.Minute)
	for {
		select {
		case <-ticker.C:
			fmt.Println("Updating proposals")
			updateAllGuilds(s)
		}
	}
}

// updateAllGuilds encapsulates the logic to update all guilds, making it reusable
func updateAllGuilds(s *discordgo.Session) {
	start := time.Now()
	guilds, err := s.UserGuilds(0, "", "")
	if err != nil {
		fmt.Println("Error fetching guilds:", err)
		return
	}

	for _, guild := range guilds {
		updateChannelsInGuild(s, guild.ID)
	}
	fmt.Printf("Took %s to update \n", time.Since(start))
}

func updateChannelsInGuild(s *discordgo.Session, guildID string) {
	channels, err := s.GuildChannels(guildID)
	if err != nil {
		fmt.Println("Error fetching channels for guild:", err)
		return
	}

	metabotCategoryID := os.Getenv("CATEGORY_ID")

	proposals := metadao.GetProposals()
	existingChannels := make(map[int]bool) // Track existing channels by proposal ID

	for _, channel := range channels {
		if channel.ParentID == metabotCategoryID && channel.Name != "METABOT" {
			proposalIDWithPrefix := strings.Split(channel.Name, ":")[0]
			proposalID, err := strconv.Atoi(strings.TrimPrefix(proposalIDWithPrefix, "P"))

			// Skip over any channels that don't match the format
			if err != nil {
				continue
			}

			if newName, ok := proposals[proposalID]; ok {
				// Update the channel if it exists in proposals
				_, err := s.ChannelEdit(channel.ID, &discordgo.ChannelEdit{
					Name:     newName,
					Position: channel.Position,
				})
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
			ch, err := s.GuildChannelCreate(guildID, name, discordgo.ChannelTypeGuildVoice)
			if err != nil {
				fmt.Println("Error creating channel:", err)
			}
			_, err = s.ChannelEdit(ch.ID, &discordgo.ChannelEdit{
				ParentID: metabotCategoryID,
			})
			if err != nil {
				fmt.Println("Error setting parent channel:", err)
			}
		}
	}

	reorderChannels(s, guildID, metabotCategoryID, proposals)

	updateSpotPrice(s, guildID, metabotCategoryID)
}

func updateSpotPrice(s *discordgo.Session, guildID string, categoryID string) {
	price := jup.GetMETAPrice()
	prettyPrice := fmt.Sprintf("Spot: $%s", humanize.CommafWithDigits(price, 2))

	channels, err := s.GuildChannels(guildID)
	if err != nil {
		fmt.Println("Error fetching channels for guild:", err)
		return
	}

	for _, channel := range channels {
		if channel.ParentID == categoryID && strings.HasPrefix(channel.Name, "Spot:") {
			_, err := s.ChannelEdit(channel.ID, &discordgo.ChannelEdit{
				Name:     prettyPrice,
				Position: 0,
			})
			if err != nil {
				fmt.Println("Error updating channel name:", err)
			}
			break
		}
	}
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

	for index, proposalID := range sortedProposalIDs {
		for _, channel := range categoryChannels {
			proposalIDWithPrefix := strings.Split(channel.Name, ":")[0]
			id, err := strconv.Atoi(strings.TrimPrefix(proposalIDWithPrefix, "P"))
			if err != nil {
				continue
			}
			if id == proposalID && channel.Position != index+1 {
				_, err := s.ChannelEditComplex(channel.ID, &discordgo.ChannelEdit{
					Position: index + 1,
				})
				if err != nil {
					fmt.Println("Error reordering channel:", err)
				}
				break
			}
		}
	}
}
