// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package openbook_v2

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// Update the [`delegate`](crate::state::OpenOrdersAccount::delegate) of an open orders account.
type SetDelegate struct {

	// [0] = [WRITE, SIGNER] owner
	//
	// [1] = [WRITE] openOrdersAccount
	//
	// [2] = [] delegateAccount
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewSetDelegateInstructionBuilder creates a new `SetDelegate` instruction builder.
func NewSetDelegateInstructionBuilder() *SetDelegate {
	nd := &SetDelegate{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 3),
	}
	return nd
}

// SetOwnerAccount sets the "owner" account.
func (inst *SetDelegate) SetOwnerAccount(owner ag_solanago.PublicKey) *SetDelegate {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(owner).WRITE().SIGNER()
	return inst
}

// GetOwnerAccount gets the "owner" account.
func (inst *SetDelegate) GetOwnerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetOpenOrdersAccountAccount sets the "openOrdersAccount" account.
func (inst *SetDelegate) SetOpenOrdersAccountAccount(openOrdersAccount ag_solanago.PublicKey) *SetDelegate {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(openOrdersAccount).WRITE()
	return inst
}

// GetOpenOrdersAccountAccount gets the "openOrdersAccount" account.
func (inst *SetDelegate) GetOpenOrdersAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetDelegateAccountAccount sets the "delegateAccount" account.
func (inst *SetDelegate) SetDelegateAccountAccount(delegateAccount ag_solanago.PublicKey) *SetDelegate {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(delegateAccount)
	return inst
}

// GetDelegateAccountAccount gets the "delegateAccount" account.
func (inst *SetDelegate) GetDelegateAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

func (inst SetDelegate) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_SetDelegate,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst SetDelegate) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *SetDelegate) Validate() error {
	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Owner is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.OpenOrdersAccount is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.DelegateAccount is not set")
		}
	}
	return nil
}

func (inst *SetDelegate) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("SetDelegate")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=0]").ParentFunc(func(paramsBranch ag_treeout.Branches) {})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=3]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("     owner", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("openOrders", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("  delegate", inst.AccountMetaSlice.Get(2)))
					})
				})
		})
}

func (obj SetDelegate) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	return nil
}
func (obj *SetDelegate) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	return nil
}

// NewSetDelegateInstruction declares a new SetDelegate instruction with the provided parameters and accounts.
func NewSetDelegateInstruction(
	// Accounts:
	owner ag_solanago.PublicKey,
	openOrdersAccount ag_solanago.PublicKey,
	delegateAccount ag_solanago.PublicKey) *SetDelegate {
	return NewSetDelegateInstructionBuilder().
		SetOwnerAccount(owner).
		SetOpenOrdersAccountAccount(openOrdersAccount).
		SetDelegateAccountAccount(delegateAccount)
}
