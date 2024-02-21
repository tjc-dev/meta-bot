package metadao

import (
	"context"
	"fmt"
	bin "github.com/gagliardetto/binary"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"math/big"
	"metabot/internal/autocrat_v0"
	"metabot/internal/openbook_twap"
	"os"
)

var (
	client *rpc.Client
)

type MarketData struct {
	Pass             float64
	LastObservedPass float64
	Fail             float64
	LastObservedFail float64
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
			proposals[int(prop.Number)] = fmt.Sprintf("P%d: $%.2fp > $%.2ff", prop.Number, details.Pass, details.Fail)
		} else {
			proposals[int(prop.Number)] = fmt.Sprintf("P%d: $%.2fp < $%.2ff", prop.Number, details.Pass, details.Fail)
		}
	}

	return proposals
}

func GetTWAP(proposal autocrat_v0.Proposal) MarketData {
	passMarket := GetTWAPMarket(proposal.OpenbookTwapPassMarket)
	failMarket := GetTWAPMarket(proposal.OpenbookTwapFailMarket)

	return MarketData{
		Pass:             GetDollarValue(passMarket.TwapOracle),
		Fail:             GetDollarValue(failMarket.TwapOracle),
		LastObservedPass: GetHiddenDollarValue(passMarket.TwapOracle),
		LastObservedFail: GetHiddenDollarValue(failMarket.TwapOracle),
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
