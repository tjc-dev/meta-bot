// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package openbook_twap

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// PlaceOrder is the `placeOrder` instruction.
type PlaceOrder struct {
	PlaceOrderArgs *PlaceOrderArgs

	// [0] = [SIGNER] signer
	//
	// [1] = [WRITE] openOrdersAccount
	//
	// [2] = [WRITE] twapMarket
	//
	// [3] = [WRITE] userTokenAccount
	//
	// [4] = [WRITE] market
	//
	// [5] = [WRITE] bids
	//
	// [6] = [WRITE] asks
	//
	// [7] = [WRITE] eventHeap
	//
	// [8] = [WRITE] marketVault
	//
	// [9] = [] tokenProgram
	//
	// [10] = [] openbookProgram
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewPlaceOrderInstructionBuilder creates a new `PlaceOrder` instruction builder.
func NewPlaceOrderInstructionBuilder() *PlaceOrder {
	nd := &PlaceOrder{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 11),
	}
	return nd
}

// SetPlaceOrderArgs sets the "placeOrderArgs" parameter.
func (inst *PlaceOrder) SetPlaceOrderArgs(placeOrderArgs PlaceOrderArgs) *PlaceOrder {
	inst.PlaceOrderArgs = &placeOrderArgs
	return inst
}

// SetSignerAccount sets the "signer" account.
func (inst *PlaceOrder) SetSignerAccount(signer ag_solanago.PublicKey) *PlaceOrder {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(signer).SIGNER()
	return inst
}

// GetSignerAccount gets the "signer" account.
func (inst *PlaceOrder) GetSignerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetOpenOrdersAccountAccount sets the "openOrdersAccount" account.
func (inst *PlaceOrder) SetOpenOrdersAccountAccount(openOrdersAccount ag_solanago.PublicKey) *PlaceOrder {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(openOrdersAccount).WRITE()
	return inst
}

// GetOpenOrdersAccountAccount gets the "openOrdersAccount" account.
func (inst *PlaceOrder) GetOpenOrdersAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetTwapMarketAccount sets the "twapMarket" account.
func (inst *PlaceOrder) SetTwapMarketAccount(twapMarket ag_solanago.PublicKey) *PlaceOrder {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(twapMarket).WRITE()
	return inst
}

// GetTwapMarketAccount gets the "twapMarket" account.
func (inst *PlaceOrder) GetTwapMarketAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetUserTokenAccountAccount sets the "userTokenAccount" account.
func (inst *PlaceOrder) SetUserTokenAccountAccount(userTokenAccount ag_solanago.PublicKey) *PlaceOrder {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(userTokenAccount).WRITE()
	return inst
}

// GetUserTokenAccountAccount gets the "userTokenAccount" account.
func (inst *PlaceOrder) GetUserTokenAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetMarketAccount sets the "market" account.
func (inst *PlaceOrder) SetMarketAccount(market ag_solanago.PublicKey) *PlaceOrder {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(market).WRITE()
	return inst
}

// GetMarketAccount gets the "market" account.
func (inst *PlaceOrder) GetMarketAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetBidsAccount sets the "bids" account.
func (inst *PlaceOrder) SetBidsAccount(bids ag_solanago.PublicKey) *PlaceOrder {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(bids).WRITE()
	return inst
}

// GetBidsAccount gets the "bids" account.
func (inst *PlaceOrder) GetBidsAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

// SetAsksAccount sets the "asks" account.
func (inst *PlaceOrder) SetAsksAccount(asks ag_solanago.PublicKey) *PlaceOrder {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(asks).WRITE()
	return inst
}

// GetAsksAccount gets the "asks" account.
func (inst *PlaceOrder) GetAsksAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(6)
}

// SetEventHeapAccount sets the "eventHeap" account.
func (inst *PlaceOrder) SetEventHeapAccount(eventHeap ag_solanago.PublicKey) *PlaceOrder {
	inst.AccountMetaSlice[7] = ag_solanago.Meta(eventHeap).WRITE()
	return inst
}

// GetEventHeapAccount gets the "eventHeap" account.
func (inst *PlaceOrder) GetEventHeapAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(7)
}

// SetMarketVaultAccount sets the "marketVault" account.
func (inst *PlaceOrder) SetMarketVaultAccount(marketVault ag_solanago.PublicKey) *PlaceOrder {
	inst.AccountMetaSlice[8] = ag_solanago.Meta(marketVault).WRITE()
	return inst
}

// GetMarketVaultAccount gets the "marketVault" account.
func (inst *PlaceOrder) GetMarketVaultAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(8)
}

// SetTokenProgramAccount sets the "tokenProgram" account.
func (inst *PlaceOrder) SetTokenProgramAccount(tokenProgram ag_solanago.PublicKey) *PlaceOrder {
	inst.AccountMetaSlice[9] = ag_solanago.Meta(tokenProgram)
	return inst
}

// GetTokenProgramAccount gets the "tokenProgram" account.
func (inst *PlaceOrder) GetTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(9)
}

// SetOpenbookProgramAccount sets the "openbookProgram" account.
func (inst *PlaceOrder) SetOpenbookProgramAccount(openbookProgram ag_solanago.PublicKey) *PlaceOrder {
	inst.AccountMetaSlice[10] = ag_solanago.Meta(openbookProgram)
	return inst
}

// GetOpenbookProgramAccount gets the "openbookProgram" account.
func (inst *PlaceOrder) GetOpenbookProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(10)
}

func (inst PlaceOrder) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_PlaceOrder,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst PlaceOrder) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *PlaceOrder) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.PlaceOrderArgs == nil {
			return errors.New("PlaceOrderArgs parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Signer is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.OpenOrdersAccount is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.TwapMarket is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.UserTokenAccount is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.Market is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.Bids is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.Asks is not set")
		}
		if inst.AccountMetaSlice[7] == nil {
			return errors.New("accounts.EventHeap is not set")
		}
		if inst.AccountMetaSlice[8] == nil {
			return errors.New("accounts.MarketVault is not set")
		}
		if inst.AccountMetaSlice[9] == nil {
			return errors.New("accounts.TokenProgram is not set")
		}
		if inst.AccountMetaSlice[10] == nil {
			return errors.New("accounts.OpenbookProgram is not set")
		}
	}
	return nil
}

func (inst *PlaceOrder) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("PlaceOrder")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("PlaceOrderArgs", *inst.PlaceOrderArgs))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=11]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("         signer", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("     openOrders", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("     twapMarket", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("      userToken", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("         market", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("           bids", inst.AccountMetaSlice.Get(5)))
						accountsBranch.Child(ag_format.Meta("           asks", inst.AccountMetaSlice.Get(6)))
						accountsBranch.Child(ag_format.Meta("      eventHeap", inst.AccountMetaSlice.Get(7)))
						accountsBranch.Child(ag_format.Meta("    marketVault", inst.AccountMetaSlice.Get(8)))
						accountsBranch.Child(ag_format.Meta("   tokenProgram", inst.AccountMetaSlice.Get(9)))
						accountsBranch.Child(ag_format.Meta("openbookProgram", inst.AccountMetaSlice.Get(10)))
					})
				})
		})
}

func (obj PlaceOrder) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `PlaceOrderArgs` param:
	err = encoder.Encode(obj.PlaceOrderArgs)
	if err != nil {
		return err
	}
	return nil
}
func (obj *PlaceOrder) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `PlaceOrderArgs`:
	err = decoder.Decode(&obj.PlaceOrderArgs)
	if err != nil {
		return err
	}
	return nil
}

// NewPlaceOrderInstruction declares a new PlaceOrder instruction with the provided parameters and accounts.
func NewPlaceOrderInstruction(
	// Parameters:
	placeOrderArgs PlaceOrderArgs,
	// Accounts:
	signer ag_solanago.PublicKey,
	openOrdersAccount ag_solanago.PublicKey,
	twapMarket ag_solanago.PublicKey,
	userTokenAccount ag_solanago.PublicKey,
	market ag_solanago.PublicKey,
	bids ag_solanago.PublicKey,
	asks ag_solanago.PublicKey,
	eventHeap ag_solanago.PublicKey,
	marketVault ag_solanago.PublicKey,
	tokenProgram ag_solanago.PublicKey,
	openbookProgram ag_solanago.PublicKey) *PlaceOrder {
	return NewPlaceOrderInstructionBuilder().
		SetPlaceOrderArgs(placeOrderArgs).
		SetSignerAccount(signer).
		SetOpenOrdersAccountAccount(openOrdersAccount).
		SetTwapMarketAccount(twapMarket).
		SetUserTokenAccountAccount(userTokenAccount).
		SetMarketAccount(market).
		SetBidsAccount(bids).
		SetAsksAccount(asks).
		SetEventHeapAccount(eventHeap).
		SetMarketVaultAccount(marketVault).
		SetTokenProgramAccount(tokenProgram).
		SetOpenbookProgramAccount(openbookProgram)
}