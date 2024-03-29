// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package openbook_v2

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// Close an [`OpenOrdersIndexer`](crate::state::OpenOrdersIndexer) account.
type CloseOpenOrdersIndexer struct {

	// [0] = [SIGNER] owner
	//
	// [1] = [WRITE] openOrdersIndexer
	//
	// [2] = [WRITE] solDestination
	//
	// [3] = [] tokenProgram
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewCloseOpenOrdersIndexerInstructionBuilder creates a new `CloseOpenOrdersIndexer` instruction builder.
func NewCloseOpenOrdersIndexerInstructionBuilder() *CloseOpenOrdersIndexer {
	nd := &CloseOpenOrdersIndexer{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 4),
	}
	return nd
}

// SetOwnerAccount sets the "owner" account.
func (inst *CloseOpenOrdersIndexer) SetOwnerAccount(owner ag_solanago.PublicKey) *CloseOpenOrdersIndexer {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(owner).SIGNER()
	return inst
}

// GetOwnerAccount gets the "owner" account.
func (inst *CloseOpenOrdersIndexer) GetOwnerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetOpenOrdersIndexerAccount sets the "openOrdersIndexer" account.
func (inst *CloseOpenOrdersIndexer) SetOpenOrdersIndexerAccount(openOrdersIndexer ag_solanago.PublicKey) *CloseOpenOrdersIndexer {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(openOrdersIndexer).WRITE()
	return inst
}

// GetOpenOrdersIndexerAccount gets the "openOrdersIndexer" account.
func (inst *CloseOpenOrdersIndexer) GetOpenOrdersIndexerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetSolDestinationAccount sets the "solDestination" account.
func (inst *CloseOpenOrdersIndexer) SetSolDestinationAccount(solDestination ag_solanago.PublicKey) *CloseOpenOrdersIndexer {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(solDestination).WRITE()
	return inst
}

// GetSolDestinationAccount gets the "solDestination" account.
func (inst *CloseOpenOrdersIndexer) GetSolDestinationAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetTokenProgramAccount sets the "tokenProgram" account.
func (inst *CloseOpenOrdersIndexer) SetTokenProgramAccount(tokenProgram ag_solanago.PublicKey) *CloseOpenOrdersIndexer {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(tokenProgram)
	return inst
}

// GetTokenProgramAccount gets the "tokenProgram" account.
func (inst *CloseOpenOrdersIndexer) GetTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

func (inst CloseOpenOrdersIndexer) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_CloseOpenOrdersIndexer,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst CloseOpenOrdersIndexer) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *CloseOpenOrdersIndexer) Validate() error {
	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Owner is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.OpenOrdersIndexer is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.SolDestination is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.TokenProgram is not set")
		}
	}
	return nil
}

func (inst *CloseOpenOrdersIndexer) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("CloseOpenOrdersIndexer")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=0]").ParentFunc(func(paramsBranch ag_treeout.Branches) {})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=4]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("            owner", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("openOrdersIndexer", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("   solDestination", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("     tokenProgram", inst.AccountMetaSlice.Get(3)))
					})
				})
		})
}

func (obj CloseOpenOrdersIndexer) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	return nil
}
func (obj *CloseOpenOrdersIndexer) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	return nil
}

// NewCloseOpenOrdersIndexerInstruction declares a new CloseOpenOrdersIndexer instruction with the provided parameters and accounts.
func NewCloseOpenOrdersIndexerInstruction(
	// Accounts:
	owner ag_solanago.PublicKey,
	openOrdersIndexer ag_solanago.PublicKey,
	solDestination ag_solanago.PublicKey,
	tokenProgram ag_solanago.PublicKey) *CloseOpenOrdersIndexer {
	return NewCloseOpenOrdersIndexerInstructionBuilder().
		SetOwnerAccount(owner).
		SetOpenOrdersIndexerAccount(openOrdersIndexer).
		SetSolDestinationAccount(solDestination).
		SetTokenProgramAccount(tokenProgram)
}
