// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package openbook_twap

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// GetBestBidAndAsk is the `getBestBidAndAsk` instruction.
type GetBestBidAndAsk struct {

	// [0] = [] market
	//
	// [1] = [] bids
	//
	// [2] = [] asks
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewGetBestBidAndAskInstructionBuilder creates a new `GetBestBidAndAsk` instruction builder.
func NewGetBestBidAndAskInstructionBuilder() *GetBestBidAndAsk {
	nd := &GetBestBidAndAsk{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 3),
	}
	return nd
}

// SetMarketAccount sets the "market" account.
func (inst *GetBestBidAndAsk) SetMarketAccount(market ag_solanago.PublicKey) *GetBestBidAndAsk {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(market)
	return inst
}

// GetMarketAccount gets the "market" account.
func (inst *GetBestBidAndAsk) GetMarketAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetBidsAccount sets the "bids" account.
func (inst *GetBestBidAndAsk) SetBidsAccount(bids ag_solanago.PublicKey) *GetBestBidAndAsk {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(bids)
	return inst
}

// GetBidsAccount gets the "bids" account.
func (inst *GetBestBidAndAsk) GetBidsAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetAsksAccount sets the "asks" account.
func (inst *GetBestBidAndAsk) SetAsksAccount(asks ag_solanago.PublicKey) *GetBestBidAndAsk {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(asks)
	return inst
}

// GetAsksAccount gets the "asks" account.
func (inst *GetBestBidAndAsk) GetAsksAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

func (inst GetBestBidAndAsk) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_GetBestBidAndAsk,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst GetBestBidAndAsk) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *GetBestBidAndAsk) Validate() error {
	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Market is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.Bids is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.Asks is not set")
		}
	}
	return nil
}

func (inst *GetBestBidAndAsk) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("GetBestBidAndAsk")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=0]").ParentFunc(func(paramsBranch ag_treeout.Branches) {})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=3]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("market", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("  bids", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("  asks", inst.AccountMetaSlice.Get(2)))
					})
				})
		})
}

func (obj GetBestBidAndAsk) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	return nil
}
func (obj *GetBestBidAndAsk) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	return nil
}

// NewGetBestBidAndAskInstruction declares a new GetBestBidAndAsk instruction with the provided parameters and accounts.
func NewGetBestBidAndAskInstruction(
	// Accounts:
	market ag_solanago.PublicKey,
	bids ag_solanago.PublicKey,
	asks ag_solanago.PublicKey) *GetBestBidAndAsk {
	return NewGetBestBidAndAskInstructionBuilder().
		SetMarketAccount(market).
		SetBidsAccount(bids).
		SetAsksAccount(asks)
}
