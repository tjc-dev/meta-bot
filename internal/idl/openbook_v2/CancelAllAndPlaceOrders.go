// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package openbook_v2

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// Cancel orders and place multiple orders.
type CancelAllAndPlaceOrders struct {
	OrdersType *PlaceOrderType
	Bids       *[]PlaceMultipleOrdersArgs
	Asks       *[]PlaceMultipleOrdersArgs
	Limit      *uint8

	// [0] = [SIGNER] signer
	//
	// [1] = [WRITE] openOrdersAccount
	//
	// [2] = [SIGNER] openOrdersAdmin
	//
	// [3] = [WRITE] userQuoteAccount
	//
	// [4] = [WRITE] userBaseAccount
	//
	// [5] = [WRITE] market
	//
	// [6] = [WRITE] bids
	//
	// [7] = [WRITE] asks
	//
	// [8] = [WRITE] eventHeap
	//
	// [9] = [WRITE] marketQuoteVault
	//
	// [10] = [WRITE] marketBaseVault
	//
	// [11] = [] oracleA
	//
	// [12] = [] oracleB
	//
	// [13] = [] tokenProgram
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewCancelAllAndPlaceOrdersInstructionBuilder creates a new `CancelAllAndPlaceOrders` instruction builder.
func NewCancelAllAndPlaceOrdersInstructionBuilder() *CancelAllAndPlaceOrders {
	nd := &CancelAllAndPlaceOrders{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 14),
	}
	return nd
}

// SetOrdersType sets the "ordersType" parameter.
func (inst *CancelAllAndPlaceOrders) SetOrdersType(ordersType PlaceOrderType) *CancelAllAndPlaceOrders {
	inst.OrdersType = &ordersType
	return inst
}

// SetBids sets the "bids" parameter.
func (inst *CancelAllAndPlaceOrders) SetBids(bids []PlaceMultipleOrdersArgs) *CancelAllAndPlaceOrders {
	inst.Bids = &bids
	return inst
}

// SetAsks sets the "asks" parameter.
func (inst *CancelAllAndPlaceOrders) SetAsks(asks []PlaceMultipleOrdersArgs) *CancelAllAndPlaceOrders {
	inst.Asks = &asks
	return inst
}

// SetLimit sets the "limit" parameter.
func (inst *CancelAllAndPlaceOrders) SetLimit(limit uint8) *CancelAllAndPlaceOrders {
	inst.Limit = &limit
	return inst
}

// SetSignerAccount sets the "signer" account.
func (inst *CancelAllAndPlaceOrders) SetSignerAccount(signer ag_solanago.PublicKey) *CancelAllAndPlaceOrders {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(signer).SIGNER()
	return inst
}

// GetSignerAccount gets the "signer" account.
func (inst *CancelAllAndPlaceOrders) GetSignerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetOpenOrdersAccountAccount sets the "openOrdersAccount" account.
func (inst *CancelAllAndPlaceOrders) SetOpenOrdersAccountAccount(openOrdersAccount ag_solanago.PublicKey) *CancelAllAndPlaceOrders {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(openOrdersAccount).WRITE()
	return inst
}

// GetOpenOrdersAccountAccount gets the "openOrdersAccount" account.
func (inst *CancelAllAndPlaceOrders) GetOpenOrdersAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetOpenOrdersAdminAccount sets the "openOrdersAdmin" account.
func (inst *CancelAllAndPlaceOrders) SetOpenOrdersAdminAccount(openOrdersAdmin ag_solanago.PublicKey) *CancelAllAndPlaceOrders {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(openOrdersAdmin).SIGNER()
	return inst
}

// GetOpenOrdersAdminAccount gets the "openOrdersAdmin" account.
func (inst *CancelAllAndPlaceOrders) GetOpenOrdersAdminAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetUserQuoteAccountAccount sets the "userQuoteAccount" account.
func (inst *CancelAllAndPlaceOrders) SetUserQuoteAccountAccount(userQuoteAccount ag_solanago.PublicKey) *CancelAllAndPlaceOrders {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(userQuoteAccount).WRITE()
	return inst
}

// GetUserQuoteAccountAccount gets the "userQuoteAccount" account.
func (inst *CancelAllAndPlaceOrders) GetUserQuoteAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetUserBaseAccountAccount sets the "userBaseAccount" account.
func (inst *CancelAllAndPlaceOrders) SetUserBaseAccountAccount(userBaseAccount ag_solanago.PublicKey) *CancelAllAndPlaceOrders {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(userBaseAccount).WRITE()
	return inst
}

// GetUserBaseAccountAccount gets the "userBaseAccount" account.
func (inst *CancelAllAndPlaceOrders) GetUserBaseAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetMarketAccount sets the "market" account.
func (inst *CancelAllAndPlaceOrders) SetMarketAccount(market ag_solanago.PublicKey) *CancelAllAndPlaceOrders {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(market).WRITE()
	return inst
}

// GetMarketAccount gets the "market" account.
func (inst *CancelAllAndPlaceOrders) GetMarketAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

// SetBidsAccount sets the "bids" account.
func (inst *CancelAllAndPlaceOrders) SetBidsAccount(bids ag_solanago.PublicKey) *CancelAllAndPlaceOrders {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(bids).WRITE()
	return inst
}

// GetBidsAccount gets the "bids" account.
func (inst *CancelAllAndPlaceOrders) GetBidsAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(6)
}

// SetAsksAccount sets the "asks" account.
func (inst *CancelAllAndPlaceOrders) SetAsksAccount(asks ag_solanago.PublicKey) *CancelAllAndPlaceOrders {
	inst.AccountMetaSlice[7] = ag_solanago.Meta(asks).WRITE()
	return inst
}

// GetAsksAccount gets the "asks" account.
func (inst *CancelAllAndPlaceOrders) GetAsksAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(7)
}

// SetEventHeapAccount sets the "eventHeap" account.
func (inst *CancelAllAndPlaceOrders) SetEventHeapAccount(eventHeap ag_solanago.PublicKey) *CancelAllAndPlaceOrders {
	inst.AccountMetaSlice[8] = ag_solanago.Meta(eventHeap).WRITE()
	return inst
}

// GetEventHeapAccount gets the "eventHeap" account.
func (inst *CancelAllAndPlaceOrders) GetEventHeapAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(8)
}

// SetMarketQuoteVaultAccount sets the "marketQuoteVault" account.
func (inst *CancelAllAndPlaceOrders) SetMarketQuoteVaultAccount(marketQuoteVault ag_solanago.PublicKey) *CancelAllAndPlaceOrders {
	inst.AccountMetaSlice[9] = ag_solanago.Meta(marketQuoteVault).WRITE()
	return inst
}

// GetMarketQuoteVaultAccount gets the "marketQuoteVault" account.
func (inst *CancelAllAndPlaceOrders) GetMarketQuoteVaultAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(9)
}

// SetMarketBaseVaultAccount sets the "marketBaseVault" account.
func (inst *CancelAllAndPlaceOrders) SetMarketBaseVaultAccount(marketBaseVault ag_solanago.PublicKey) *CancelAllAndPlaceOrders {
	inst.AccountMetaSlice[10] = ag_solanago.Meta(marketBaseVault).WRITE()
	return inst
}

// GetMarketBaseVaultAccount gets the "marketBaseVault" account.
func (inst *CancelAllAndPlaceOrders) GetMarketBaseVaultAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(10)
}

// SetOracleAAccount sets the "oracleA" account.
func (inst *CancelAllAndPlaceOrders) SetOracleAAccount(oracleA ag_solanago.PublicKey) *CancelAllAndPlaceOrders {
	inst.AccountMetaSlice[11] = ag_solanago.Meta(oracleA)
	return inst
}

// GetOracleAAccount gets the "oracleA" account.
func (inst *CancelAllAndPlaceOrders) GetOracleAAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(11)
}

// SetOracleBAccount sets the "oracleB" account.
func (inst *CancelAllAndPlaceOrders) SetOracleBAccount(oracleB ag_solanago.PublicKey) *CancelAllAndPlaceOrders {
	inst.AccountMetaSlice[12] = ag_solanago.Meta(oracleB)
	return inst
}

// GetOracleBAccount gets the "oracleB" account.
func (inst *CancelAllAndPlaceOrders) GetOracleBAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(12)
}

// SetTokenProgramAccount sets the "tokenProgram" account.
func (inst *CancelAllAndPlaceOrders) SetTokenProgramAccount(tokenProgram ag_solanago.PublicKey) *CancelAllAndPlaceOrders {
	inst.AccountMetaSlice[13] = ag_solanago.Meta(tokenProgram)
	return inst
}

// GetTokenProgramAccount gets the "tokenProgram" account.
func (inst *CancelAllAndPlaceOrders) GetTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(13)
}

func (inst CancelAllAndPlaceOrders) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_CancelAllAndPlaceOrders,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst CancelAllAndPlaceOrders) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *CancelAllAndPlaceOrders) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.OrdersType == nil {
			return errors.New("OrdersType parameter is not set")
		}
		if inst.Bids == nil {
			return errors.New("Bids parameter is not set")
		}
		if inst.Asks == nil {
			return errors.New("Asks parameter is not set")
		}
		if inst.Limit == nil {
			return errors.New("Limit parameter is not set")
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
			return errors.New("accounts.OpenOrdersAdmin is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.UserQuoteAccount is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.UserBaseAccount is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.Market is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.Bids is not set")
		}
		if inst.AccountMetaSlice[7] == nil {
			return errors.New("accounts.Asks is not set")
		}
		if inst.AccountMetaSlice[8] == nil {
			return errors.New("accounts.EventHeap is not set")
		}
		if inst.AccountMetaSlice[9] == nil {
			return errors.New("accounts.MarketQuoteVault is not set")
		}
		if inst.AccountMetaSlice[10] == nil {
			return errors.New("accounts.MarketBaseVault is not set")
		}
		if inst.AccountMetaSlice[11] == nil {
			return errors.New("accounts.OracleA is not set")
		}
		if inst.AccountMetaSlice[12] == nil {
			return errors.New("accounts.OracleB is not set")
		}
		if inst.AccountMetaSlice[13] == nil {
			return errors.New("accounts.TokenProgram is not set")
		}
	}
	return nil
}

func (inst *CancelAllAndPlaceOrders) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("CancelAllAndPlaceOrders")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=4]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("OrdersType", *inst.OrdersType))
						paramsBranch.Child(ag_format.Param("      Bids", *inst.Bids))
						paramsBranch.Child(ag_format.Param("      Asks", *inst.Asks))
						paramsBranch.Child(ag_format.Param("     Limit", *inst.Limit))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=14]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("          signer", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("      openOrders", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta(" openOrdersAdmin", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("       userQuote", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("        userBase", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("          market", inst.AccountMetaSlice.Get(5)))
						accountsBranch.Child(ag_format.Meta("            bids", inst.AccountMetaSlice.Get(6)))
						accountsBranch.Child(ag_format.Meta("            asks", inst.AccountMetaSlice.Get(7)))
						accountsBranch.Child(ag_format.Meta("       eventHeap", inst.AccountMetaSlice.Get(8)))
						accountsBranch.Child(ag_format.Meta("marketQuoteVault", inst.AccountMetaSlice.Get(9)))
						accountsBranch.Child(ag_format.Meta(" marketBaseVault", inst.AccountMetaSlice.Get(10)))
						accountsBranch.Child(ag_format.Meta("         oracleA", inst.AccountMetaSlice.Get(11)))
						accountsBranch.Child(ag_format.Meta("         oracleB", inst.AccountMetaSlice.Get(12)))
						accountsBranch.Child(ag_format.Meta("    tokenProgram", inst.AccountMetaSlice.Get(13)))
					})
				})
		})
}

func (obj CancelAllAndPlaceOrders) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `OrdersType` param:
	err = encoder.Encode(obj.OrdersType)
	if err != nil {
		return err
	}
	// Serialize `Bids` param:
	err = encoder.Encode(obj.Bids)
	if err != nil {
		return err
	}
	// Serialize `Asks` param:
	err = encoder.Encode(obj.Asks)
	if err != nil {
		return err
	}
	// Serialize `Limit` param:
	err = encoder.Encode(obj.Limit)
	if err != nil {
		return err
	}
	return nil
}
func (obj *CancelAllAndPlaceOrders) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `OrdersType`:
	err = decoder.Decode(&obj.OrdersType)
	if err != nil {
		return err
	}
	// Deserialize `Bids`:
	err = decoder.Decode(&obj.Bids)
	if err != nil {
		return err
	}
	// Deserialize `Asks`:
	err = decoder.Decode(&obj.Asks)
	if err != nil {
		return err
	}
	// Deserialize `Limit`:
	err = decoder.Decode(&obj.Limit)
	if err != nil {
		return err
	}
	return nil
}

// NewCancelAllAndPlaceOrdersInstruction declares a new CancelAllAndPlaceOrders instruction with the provided parameters and accounts.
func NewCancelAllAndPlaceOrdersInstruction(
	// Parameters:
	ordersType PlaceOrderType,
	bids []PlaceMultipleOrdersArgs,
	asks []PlaceMultipleOrdersArgs,
	limit uint8,
	// Accounts:
	signer ag_solanago.PublicKey,
	openOrdersAccount ag_solanago.PublicKey,
	openOrdersAdmin ag_solanago.PublicKey,
	userQuoteAccount ag_solanago.PublicKey,
	userBaseAccount ag_solanago.PublicKey,
	market ag_solanago.PublicKey,
	bidsAccount ag_solanago.PublicKey,
	asksAccount ag_solanago.PublicKey,
	eventHeap ag_solanago.PublicKey,
	marketQuoteVault ag_solanago.PublicKey,
	marketBaseVault ag_solanago.PublicKey,
	oracleA ag_solanago.PublicKey,
	oracleB ag_solanago.PublicKey,
	tokenProgram ag_solanago.PublicKey) *CancelAllAndPlaceOrders {
	return NewCancelAllAndPlaceOrdersInstructionBuilder().
		SetOrdersType(ordersType).
		SetBids(bids).
		SetAsks(asks).
		SetLimit(limit).
		SetSignerAccount(signer).
		SetOpenOrdersAccountAccount(openOrdersAccount).
		SetOpenOrdersAdminAccount(openOrdersAdmin).
		SetUserQuoteAccountAccount(userQuoteAccount).
		SetUserBaseAccountAccount(userBaseAccount).
		SetMarketAccount(market).
		SetBidsAccount(bidsAccount).
		SetAsksAccount(asksAccount).
		SetEventHeapAccount(eventHeap).
		SetMarketQuoteVaultAccount(marketQuoteVault).
		SetMarketBaseVaultAccount(marketBaseVault).
		SetOracleAAccount(oracleA).
		SetOracleBAccount(oracleB).
		SetTokenProgramAccount(tokenProgram)
}
