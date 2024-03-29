// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package openbook_v2

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// Remove orders from the book when the market is expired (only
// [`close_market_admin`](crate::state::Market::close_market_admin)).
type PruneOrders struct {
	Limit *uint8

	// [0] = [SIGNER] closeMarketAdmin
	//
	// [1] = [WRITE] openOrdersAccount
	//
	// [2] = [] market
	//
	// [3] = [WRITE] bids
	//
	// [4] = [WRITE] asks
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewPruneOrdersInstructionBuilder creates a new `PruneOrders` instruction builder.
func NewPruneOrdersInstructionBuilder() *PruneOrders {
	nd := &PruneOrders{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 5),
	}
	return nd
}

// SetLimit sets the "limit" parameter.
func (inst *PruneOrders) SetLimit(limit uint8) *PruneOrders {
	inst.Limit = &limit
	return inst
}

// SetCloseMarketAdminAccount sets the "closeMarketAdmin" account.
func (inst *PruneOrders) SetCloseMarketAdminAccount(closeMarketAdmin ag_solanago.PublicKey) *PruneOrders {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(closeMarketAdmin).SIGNER()
	return inst
}

// GetCloseMarketAdminAccount gets the "closeMarketAdmin" account.
func (inst *PruneOrders) GetCloseMarketAdminAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetOpenOrdersAccountAccount sets the "openOrdersAccount" account.
func (inst *PruneOrders) SetOpenOrdersAccountAccount(openOrdersAccount ag_solanago.PublicKey) *PruneOrders {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(openOrdersAccount).WRITE()
	return inst
}

// GetOpenOrdersAccountAccount gets the "openOrdersAccount" account.
func (inst *PruneOrders) GetOpenOrdersAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetMarketAccount sets the "market" account.
func (inst *PruneOrders) SetMarketAccount(market ag_solanago.PublicKey) *PruneOrders {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(market)
	return inst
}

// GetMarketAccount gets the "market" account.
func (inst *PruneOrders) GetMarketAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetBidsAccount sets the "bids" account.
func (inst *PruneOrders) SetBidsAccount(bids ag_solanago.PublicKey) *PruneOrders {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(bids).WRITE()
	return inst
}

// GetBidsAccount gets the "bids" account.
func (inst *PruneOrders) GetBidsAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetAsksAccount sets the "asks" account.
func (inst *PruneOrders) SetAsksAccount(asks ag_solanago.PublicKey) *PruneOrders {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(asks).WRITE()
	return inst
}

// GetAsksAccount gets the "asks" account.
func (inst *PruneOrders) GetAsksAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

func (inst PruneOrders) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_PruneOrders,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst PruneOrders) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *PruneOrders) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.Limit == nil {
			return errors.New("Limit parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.CloseMarketAdmin is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.OpenOrdersAccount is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.Market is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.Bids is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.Asks is not set")
		}
	}
	return nil
}

func (inst *PruneOrders) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("PruneOrders")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("Limit", *inst.Limit))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=5]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("closeMarketAdmin", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("      openOrders", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("          market", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("            bids", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("            asks", inst.AccountMetaSlice.Get(4)))
					})
				})
		})
}

func (obj PruneOrders) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Limit` param:
	err = encoder.Encode(obj.Limit)
	if err != nil {
		return err
	}
	return nil
}
func (obj *PruneOrders) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Limit`:
	err = decoder.Decode(&obj.Limit)
	if err != nil {
		return err
	}
	return nil
}

// NewPruneOrdersInstruction declares a new PruneOrders instruction with the provided parameters and accounts.
func NewPruneOrdersInstruction(
	// Parameters:
	limit uint8,
	// Accounts:
	closeMarketAdmin ag_solanago.PublicKey,
	openOrdersAccount ag_solanago.PublicKey,
	market ag_solanago.PublicKey,
	bids ag_solanago.PublicKey,
	asks ag_solanago.PublicKey) *PruneOrders {
	return NewPruneOrdersInstructionBuilder().
		SetLimit(limit).
		SetCloseMarketAdminAccount(closeMarketAdmin).
		SetOpenOrdersAccountAccount(openOrdersAccount).
		SetMarketAccount(market).
		SetBidsAccount(bids).
		SetAsksAccount(asks)
}
