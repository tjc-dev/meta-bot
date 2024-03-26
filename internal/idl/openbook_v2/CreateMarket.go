// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package openbook_v2

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// Create a [`Market`](crate::state::Market) for a given token pair.
type CreateMarket struct {
	Name         *string
	OracleConfig *OracleConfigParams
	QuoteLotSize *int64
	BaseLotSize  *int64
	MakerFee     *int64
	TakerFee     *int64
	TimeExpiry   *int64

	// [0] = [WRITE, SIGNER] market
	//
	// [1] = [] marketAuthority
	//
	// [2] = [WRITE] bids
	// ··········· Accounts are initialized by client,
	// ··········· anchor discriminator is set first when ix exits,
	//
	// [3] = [WRITE] asks
	//
	// [4] = [WRITE] eventHeap
	//
	// [5] = [WRITE, SIGNER] payer
	//
	// [6] = [WRITE] marketBaseVault
	//
	// [7] = [WRITE] marketQuoteVault
	//
	// [8] = [] baseMint
	//
	// [9] = [] quoteMint
	//
	// [10] = [] systemProgram
	//
	// [11] = [] tokenProgram
	//
	// [12] = [] associatedTokenProgram
	//
	// [13] = [] oracleA
	//
	// [14] = [] oracleB
	//
	// [15] = [] collectFeeAdmin
	//
	// [16] = [] openOrdersAdmin
	//
	// [17] = [] consumeEventsAdmin
	//
	// [18] = [] closeMarketAdmin
	//
	// [19] = [] eventAuthority
	//
	// [20] = [] program
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewCreateMarketInstructionBuilder creates a new `CreateMarket` instruction builder.
func NewCreateMarketInstructionBuilder() *CreateMarket {
	nd := &CreateMarket{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 21),
	}
	return nd
}

// SetName sets the "name" parameter.
func (inst *CreateMarket) SetName(name string) *CreateMarket {
	inst.Name = &name
	return inst
}

// SetOracleConfig sets the "oracleConfig" parameter.
func (inst *CreateMarket) SetOracleConfig(oracleConfig OracleConfigParams) *CreateMarket {
	inst.OracleConfig = &oracleConfig
	return inst
}

// SetQuoteLotSize sets the "quoteLotSize" parameter.
func (inst *CreateMarket) SetQuoteLotSize(quoteLotSize int64) *CreateMarket {
	inst.QuoteLotSize = &quoteLotSize
	return inst
}

// SetBaseLotSize sets the "baseLotSize" parameter.
func (inst *CreateMarket) SetBaseLotSize(baseLotSize int64) *CreateMarket {
	inst.BaseLotSize = &baseLotSize
	return inst
}

// SetMakerFee sets the "makerFee" parameter.
func (inst *CreateMarket) SetMakerFee(makerFee int64) *CreateMarket {
	inst.MakerFee = &makerFee
	return inst
}

// SetTakerFee sets the "takerFee" parameter.
func (inst *CreateMarket) SetTakerFee(takerFee int64) *CreateMarket {
	inst.TakerFee = &takerFee
	return inst
}

// SetTimeExpiry sets the "timeExpiry" parameter.
func (inst *CreateMarket) SetTimeExpiry(timeExpiry int64) *CreateMarket {
	inst.TimeExpiry = &timeExpiry
	return inst
}

// SetMarketAccount sets the "market" account.
func (inst *CreateMarket) SetMarketAccount(market ag_solanago.PublicKey) *CreateMarket {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(market).WRITE().SIGNER()
	return inst
}

// GetMarketAccount gets the "market" account.
func (inst *CreateMarket) GetMarketAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetMarketAuthorityAccount sets the "marketAuthority" account.
func (inst *CreateMarket) SetMarketAuthorityAccount(marketAuthority ag_solanago.PublicKey) *CreateMarket {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(marketAuthority)
	return inst
}

// GetMarketAuthorityAccount gets the "marketAuthority" account.
func (inst *CreateMarket) GetMarketAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetBidsAccount sets the "bids" account.
// Accounts are initialized by client,
// anchor discriminator is set first when ix exits,
func (inst *CreateMarket) SetBidsAccount(bids ag_solanago.PublicKey) *CreateMarket {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(bids).WRITE()
	return inst
}

// GetBidsAccount gets the "bids" account.
// Accounts are initialized by client,
// anchor discriminator is set first when ix exits,
func (inst *CreateMarket) GetBidsAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetAsksAccount sets the "asks" account.
func (inst *CreateMarket) SetAsksAccount(asks ag_solanago.PublicKey) *CreateMarket {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(asks).WRITE()
	return inst
}

// GetAsksAccount gets the "asks" account.
func (inst *CreateMarket) GetAsksAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetEventHeapAccount sets the "eventHeap" account.
func (inst *CreateMarket) SetEventHeapAccount(eventHeap ag_solanago.PublicKey) *CreateMarket {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(eventHeap).WRITE()
	return inst
}

// GetEventHeapAccount gets the "eventHeap" account.
func (inst *CreateMarket) GetEventHeapAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetPayerAccount sets the "payer" account.
func (inst *CreateMarket) SetPayerAccount(payer ag_solanago.PublicKey) *CreateMarket {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(payer).WRITE().SIGNER()
	return inst
}

// GetPayerAccount gets the "payer" account.
func (inst *CreateMarket) GetPayerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

// SetMarketBaseVaultAccount sets the "marketBaseVault" account.
func (inst *CreateMarket) SetMarketBaseVaultAccount(marketBaseVault ag_solanago.PublicKey) *CreateMarket {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(marketBaseVault).WRITE()
	return inst
}

// GetMarketBaseVaultAccount gets the "marketBaseVault" account.
func (inst *CreateMarket) GetMarketBaseVaultAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(6)
}

// SetMarketQuoteVaultAccount sets the "marketQuoteVault" account.
func (inst *CreateMarket) SetMarketQuoteVaultAccount(marketQuoteVault ag_solanago.PublicKey) *CreateMarket {
	inst.AccountMetaSlice[7] = ag_solanago.Meta(marketQuoteVault).WRITE()
	return inst
}

// GetMarketQuoteVaultAccount gets the "marketQuoteVault" account.
func (inst *CreateMarket) GetMarketQuoteVaultAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(7)
}

// SetBaseMintAccount sets the "baseMint" account.
func (inst *CreateMarket) SetBaseMintAccount(baseMint ag_solanago.PublicKey) *CreateMarket {
	inst.AccountMetaSlice[8] = ag_solanago.Meta(baseMint)
	return inst
}

// GetBaseMintAccount gets the "baseMint" account.
func (inst *CreateMarket) GetBaseMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(8)
}

// SetQuoteMintAccount sets the "quoteMint" account.
func (inst *CreateMarket) SetQuoteMintAccount(quoteMint ag_solanago.PublicKey) *CreateMarket {
	inst.AccountMetaSlice[9] = ag_solanago.Meta(quoteMint)
	return inst
}

// GetQuoteMintAccount gets the "quoteMint" account.
func (inst *CreateMarket) GetQuoteMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(9)
}

// SetSystemProgramAccount sets the "systemProgram" account.
func (inst *CreateMarket) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *CreateMarket {
	inst.AccountMetaSlice[10] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "systemProgram" account.
func (inst *CreateMarket) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(10)
}

// SetTokenProgramAccount sets the "tokenProgram" account.
func (inst *CreateMarket) SetTokenProgramAccount(tokenProgram ag_solanago.PublicKey) *CreateMarket {
	inst.AccountMetaSlice[11] = ag_solanago.Meta(tokenProgram)
	return inst
}

// GetTokenProgramAccount gets the "tokenProgram" account.
func (inst *CreateMarket) GetTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(11)
}

// SetAssociatedTokenProgramAccount sets the "associatedTokenProgram" account.
func (inst *CreateMarket) SetAssociatedTokenProgramAccount(associatedTokenProgram ag_solanago.PublicKey) *CreateMarket {
	inst.AccountMetaSlice[12] = ag_solanago.Meta(associatedTokenProgram)
	return inst
}

// GetAssociatedTokenProgramAccount gets the "associatedTokenProgram" account.
func (inst *CreateMarket) GetAssociatedTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(12)
}

// SetOracleAAccount sets the "oracleA" account.
func (inst *CreateMarket) SetOracleAAccount(oracleA ag_solanago.PublicKey) *CreateMarket {
	inst.AccountMetaSlice[13] = ag_solanago.Meta(oracleA)
	return inst
}

// GetOracleAAccount gets the "oracleA" account.
func (inst *CreateMarket) GetOracleAAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(13)
}

// SetOracleBAccount sets the "oracleB" account.
func (inst *CreateMarket) SetOracleBAccount(oracleB ag_solanago.PublicKey) *CreateMarket {
	inst.AccountMetaSlice[14] = ag_solanago.Meta(oracleB)
	return inst
}

// GetOracleBAccount gets the "oracleB" account.
func (inst *CreateMarket) GetOracleBAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(14)
}

// SetCollectFeeAdminAccount sets the "collectFeeAdmin" account.
func (inst *CreateMarket) SetCollectFeeAdminAccount(collectFeeAdmin ag_solanago.PublicKey) *CreateMarket {
	inst.AccountMetaSlice[15] = ag_solanago.Meta(collectFeeAdmin)
	return inst
}

// GetCollectFeeAdminAccount gets the "collectFeeAdmin" account.
func (inst *CreateMarket) GetCollectFeeAdminAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(15)
}

// SetOpenOrdersAdminAccount sets the "openOrdersAdmin" account.
func (inst *CreateMarket) SetOpenOrdersAdminAccount(openOrdersAdmin ag_solanago.PublicKey) *CreateMarket {
	inst.AccountMetaSlice[16] = ag_solanago.Meta(openOrdersAdmin)
	return inst
}

// GetOpenOrdersAdminAccount gets the "openOrdersAdmin" account.
func (inst *CreateMarket) GetOpenOrdersAdminAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(16)
}

// SetConsumeEventsAdminAccount sets the "consumeEventsAdmin" account.
func (inst *CreateMarket) SetConsumeEventsAdminAccount(consumeEventsAdmin ag_solanago.PublicKey) *CreateMarket {
	inst.AccountMetaSlice[17] = ag_solanago.Meta(consumeEventsAdmin)
	return inst
}

// GetConsumeEventsAdminAccount gets the "consumeEventsAdmin" account.
func (inst *CreateMarket) GetConsumeEventsAdminAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(17)
}

// SetCloseMarketAdminAccount sets the "closeMarketAdmin" account.
func (inst *CreateMarket) SetCloseMarketAdminAccount(closeMarketAdmin ag_solanago.PublicKey) *CreateMarket {
	inst.AccountMetaSlice[18] = ag_solanago.Meta(closeMarketAdmin)
	return inst
}

// GetCloseMarketAdminAccount gets the "closeMarketAdmin" account.
func (inst *CreateMarket) GetCloseMarketAdminAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(18)
}

// SetEventAuthorityAccount sets the "eventAuthority" account.
func (inst *CreateMarket) SetEventAuthorityAccount(eventAuthority ag_solanago.PublicKey) *CreateMarket {
	inst.AccountMetaSlice[19] = ag_solanago.Meta(eventAuthority)
	return inst
}

// GetEventAuthorityAccount gets the "eventAuthority" account.
func (inst *CreateMarket) GetEventAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(19)
}

// SetProgramAccount sets the "program" account.
func (inst *CreateMarket) SetProgramAccount(program ag_solanago.PublicKey) *CreateMarket {
	inst.AccountMetaSlice[20] = ag_solanago.Meta(program)
	return inst
}

// GetProgramAccount gets the "program" account.
func (inst *CreateMarket) GetProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(20)
}

func (inst CreateMarket) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_CreateMarket,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst CreateMarket) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *CreateMarket) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.Name == nil {
			return errors.New("Name parameter is not set")
		}
		if inst.OracleConfig == nil {
			return errors.New("OracleConfig parameter is not set")
		}
		if inst.QuoteLotSize == nil {
			return errors.New("QuoteLotSize parameter is not set")
		}
		if inst.BaseLotSize == nil {
			return errors.New("BaseLotSize parameter is not set")
		}
		if inst.MakerFee == nil {
			return errors.New("MakerFee parameter is not set")
		}
		if inst.TakerFee == nil {
			return errors.New("TakerFee parameter is not set")
		}
		if inst.TimeExpiry == nil {
			return errors.New("TimeExpiry parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Market is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.MarketAuthority is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.Bids is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.Asks is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.EventHeap is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.Payer is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.MarketBaseVault is not set")
		}
		if inst.AccountMetaSlice[7] == nil {
			return errors.New("accounts.MarketQuoteVault is not set")
		}
		if inst.AccountMetaSlice[8] == nil {
			return errors.New("accounts.BaseMint is not set")
		}
		if inst.AccountMetaSlice[9] == nil {
			return errors.New("accounts.QuoteMint is not set")
		}
		if inst.AccountMetaSlice[10] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
		if inst.AccountMetaSlice[11] == nil {
			return errors.New("accounts.TokenProgram is not set")
		}
		if inst.AccountMetaSlice[12] == nil {
			return errors.New("accounts.AssociatedTokenProgram is not set")
		}
		if inst.AccountMetaSlice[13] == nil {
			return errors.New("accounts.OracleA is not set")
		}
		if inst.AccountMetaSlice[14] == nil {
			return errors.New("accounts.OracleB is not set")
		}
		if inst.AccountMetaSlice[15] == nil {
			return errors.New("accounts.CollectFeeAdmin is not set")
		}
		if inst.AccountMetaSlice[16] == nil {
			return errors.New("accounts.OpenOrdersAdmin is not set")
		}
		if inst.AccountMetaSlice[17] == nil {
			return errors.New("accounts.ConsumeEventsAdmin is not set")
		}
		if inst.AccountMetaSlice[18] == nil {
			return errors.New("accounts.CloseMarketAdmin is not set")
		}
		if inst.AccountMetaSlice[19] == nil {
			return errors.New("accounts.EventAuthority is not set")
		}
		if inst.AccountMetaSlice[20] == nil {
			return errors.New("accounts.Program is not set")
		}
	}
	return nil
}

func (inst *CreateMarket) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("CreateMarket")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=7]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("        Name", *inst.Name))
						paramsBranch.Child(ag_format.Param("OracleConfig", *inst.OracleConfig))
						paramsBranch.Child(ag_format.Param("QuoteLotSize", *inst.QuoteLotSize))
						paramsBranch.Child(ag_format.Param(" BaseLotSize", *inst.BaseLotSize))
						paramsBranch.Child(ag_format.Param("    MakerFee", *inst.MakerFee))
						paramsBranch.Child(ag_format.Param("    TakerFee", *inst.TakerFee))
						paramsBranch.Child(ag_format.Param("  TimeExpiry", *inst.TimeExpiry))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=21]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("                market", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("       marketAuthority", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("                  bids", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("                  asks", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("             eventHeap", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("                 payer", inst.AccountMetaSlice.Get(5)))
						accountsBranch.Child(ag_format.Meta("       marketBaseVault", inst.AccountMetaSlice.Get(6)))
						accountsBranch.Child(ag_format.Meta("      marketQuoteVault", inst.AccountMetaSlice.Get(7)))
						accountsBranch.Child(ag_format.Meta("              baseMint", inst.AccountMetaSlice.Get(8)))
						accountsBranch.Child(ag_format.Meta("             quoteMint", inst.AccountMetaSlice.Get(9)))
						accountsBranch.Child(ag_format.Meta("         systemProgram", inst.AccountMetaSlice.Get(10)))
						accountsBranch.Child(ag_format.Meta("          tokenProgram", inst.AccountMetaSlice.Get(11)))
						accountsBranch.Child(ag_format.Meta("associatedTokenProgram", inst.AccountMetaSlice.Get(12)))
						accountsBranch.Child(ag_format.Meta("               oracleA", inst.AccountMetaSlice.Get(13)))
						accountsBranch.Child(ag_format.Meta("               oracleB", inst.AccountMetaSlice.Get(14)))
						accountsBranch.Child(ag_format.Meta("       collectFeeAdmin", inst.AccountMetaSlice.Get(15)))
						accountsBranch.Child(ag_format.Meta("       openOrdersAdmin", inst.AccountMetaSlice.Get(16)))
						accountsBranch.Child(ag_format.Meta("    consumeEventsAdmin", inst.AccountMetaSlice.Get(17)))
						accountsBranch.Child(ag_format.Meta("      closeMarketAdmin", inst.AccountMetaSlice.Get(18)))
						accountsBranch.Child(ag_format.Meta("        eventAuthority", inst.AccountMetaSlice.Get(19)))
						accountsBranch.Child(ag_format.Meta("               program", inst.AccountMetaSlice.Get(20)))
					})
				})
		})
}

func (obj CreateMarket) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Name` param:
	err = encoder.Encode(obj.Name)
	if err != nil {
		return err
	}
	// Serialize `OracleConfig` param:
	err = encoder.Encode(obj.OracleConfig)
	if err != nil {
		return err
	}
	// Serialize `QuoteLotSize` param:
	err = encoder.Encode(obj.QuoteLotSize)
	if err != nil {
		return err
	}
	// Serialize `BaseLotSize` param:
	err = encoder.Encode(obj.BaseLotSize)
	if err != nil {
		return err
	}
	// Serialize `MakerFee` param:
	err = encoder.Encode(obj.MakerFee)
	if err != nil {
		return err
	}
	// Serialize `TakerFee` param:
	err = encoder.Encode(obj.TakerFee)
	if err != nil {
		return err
	}
	// Serialize `TimeExpiry` param:
	err = encoder.Encode(obj.TimeExpiry)
	if err != nil {
		return err
	}
	return nil
}
func (obj *CreateMarket) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Name`:
	err = decoder.Decode(&obj.Name)
	if err != nil {
		return err
	}
	// Deserialize `OracleConfig`:
	err = decoder.Decode(&obj.OracleConfig)
	if err != nil {
		return err
	}
	// Deserialize `QuoteLotSize`:
	err = decoder.Decode(&obj.QuoteLotSize)
	if err != nil {
		return err
	}
	// Deserialize `BaseLotSize`:
	err = decoder.Decode(&obj.BaseLotSize)
	if err != nil {
		return err
	}
	// Deserialize `MakerFee`:
	err = decoder.Decode(&obj.MakerFee)
	if err != nil {
		return err
	}
	// Deserialize `TakerFee`:
	err = decoder.Decode(&obj.TakerFee)
	if err != nil {
		return err
	}
	// Deserialize `TimeExpiry`:
	err = decoder.Decode(&obj.TimeExpiry)
	if err != nil {
		return err
	}
	return nil
}

// NewCreateMarketInstruction declares a new CreateMarket instruction with the provided parameters and accounts.
func NewCreateMarketInstruction(
	// Parameters:
	name string,
	oracleConfig OracleConfigParams,
	quoteLotSize int64,
	baseLotSize int64,
	makerFee int64,
	takerFee int64,
	timeExpiry int64,
	// Accounts:
	market ag_solanago.PublicKey,
	marketAuthority ag_solanago.PublicKey,
	bids ag_solanago.PublicKey,
	asks ag_solanago.PublicKey,
	eventHeap ag_solanago.PublicKey,
	payer ag_solanago.PublicKey,
	marketBaseVault ag_solanago.PublicKey,
	marketQuoteVault ag_solanago.PublicKey,
	baseMint ag_solanago.PublicKey,
	quoteMint ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey,
	tokenProgram ag_solanago.PublicKey,
	associatedTokenProgram ag_solanago.PublicKey,
	oracleA ag_solanago.PublicKey,
	oracleB ag_solanago.PublicKey,
	collectFeeAdmin ag_solanago.PublicKey,
	openOrdersAdmin ag_solanago.PublicKey,
	consumeEventsAdmin ag_solanago.PublicKey,
	closeMarketAdmin ag_solanago.PublicKey,
	eventAuthority ag_solanago.PublicKey,
	program ag_solanago.PublicKey) *CreateMarket {
	return NewCreateMarketInstructionBuilder().
		SetName(name).
		SetOracleConfig(oracleConfig).
		SetQuoteLotSize(quoteLotSize).
		SetBaseLotSize(baseLotSize).
		SetMakerFee(makerFee).
		SetTakerFee(takerFee).
		SetTimeExpiry(timeExpiry).
		SetMarketAccount(market).
		SetMarketAuthorityAccount(marketAuthority).
		SetBidsAccount(bids).
		SetAsksAccount(asks).
		SetEventHeapAccount(eventHeap).
		SetPayerAccount(payer).
		SetMarketBaseVaultAccount(marketBaseVault).
		SetMarketQuoteVaultAccount(marketQuoteVault).
		SetBaseMintAccount(baseMint).
		SetQuoteMintAccount(quoteMint).
		SetSystemProgramAccount(systemProgram).
		SetTokenProgramAccount(tokenProgram).
		SetAssociatedTokenProgramAccount(associatedTokenProgram).
		SetOracleAAccount(oracleA).
		SetOracleBAccount(oracleB).
		SetCollectFeeAdminAccount(collectFeeAdmin).
		SetOpenOrdersAdminAccount(openOrdersAdmin).
		SetConsumeEventsAdminAccount(consumeEventsAdmin).
		SetCloseMarketAdminAccount(closeMarketAdmin).
		SetEventAuthorityAccount(eventAuthority).
		SetProgramAccount(program)
}
