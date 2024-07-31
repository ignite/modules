package keeper

import (
	"fmt"

	"cosmossdk.io/collections"
	"cosmossdk.io/core/address"
	"cosmossdk.io/core/store"
	"cosmossdk.io/log"
	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/ignite/modules/x/claim/types"
)

type (
	Keeper struct {
		cdc          codec.BinaryCodec
		addressCodec address.Codec
		storeService store.KVStoreService
		logger       log.Logger

		// the address capable of executing a MsgUpdateParams message.
		// Typically, this should be the x/gov module account.
		authority string

		Schema        collections.Schema
		Params        collections.Item[types.Params]
		ClaimRecord   collections.Map[string, types.ClaimRecord]
		MissionSeq    collections.Sequence
		Mission       collections.Map[uint64, types.Mission]
		InitialClaim  collections.Item[types.InitialClaim]
		AirdropSupply collections.Item[types.AirdropSupply]
		// this line is used by starport scaffolding # collection/type

		accountKeeper types.AccountKeeper
		bankKeeper    types.BankKeeper
		distrKeeper   types.DistrKeeper
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	addressCodec address.Codec,
	storeService store.KVStoreService,
	logger log.Logger,
	authority string,

	accountKeeper types.AccountKeeper,
	bankKeeper types.BankKeeper,
	distrKeeper types.DistrKeeper,
) Keeper {
	if _, err := addressCodec.StringToBytes(authority); err != nil {
		panic(fmt.Sprintf("invalid authority address %s: %s", authority, err))
	}

	sb := collections.NewSchemaBuilder(storeService)

	k := Keeper{
		cdc:           cdc,
		addressCodec:  addressCodec,
		storeService:  storeService,
		authority:     authority,
		logger:        logger,
		accountKeeper: accountKeeper,
		bankKeeper:    bankKeeper,
		distrKeeper:   distrKeeper,
		Params:        collections.NewItem(sb, types.ParamsKey, "params", codec.CollValue[types.Params](cdc)),
		ClaimRecord:   collections.NewMap(sb, types.ClaimRecordKey, "claimRecord", collections.StringKey, codec.CollValue[types.ClaimRecord](cdc)),
		MissionSeq:    collections.NewSequence(sb, types.MissionCountKey, "mission"),
		Mission:       collections.NewMap(sb, types.MissionKey, "mission_seq", collections.Uint64Key, codec.CollValue[types.Mission](cdc)),
		InitialClaim:  collections.NewItem(sb, types.InitialClaimKey, "initialClaim", codec.CollValue[types.InitialClaim](cdc)),
		AirdropSupply: collections.NewItem(sb, types.AirdropSupplyKey, "airdropSupply", codec.CollValue[types.AirdropSupply](cdc)),
		// this line is used by starport scaffolding # collection/instantiate
	}

	schema, err := sb.Build()
	if err != nil {
		panic(err)
	}
	k.Schema = schema

	return k
}

// GetAuthority returns the module's authority.
func (k Keeper) GetAuthority() string {
	return k.authority
}

// Logger returns a module-specific logger.
func (k Keeper) Logger() log.Logger {
	return k.logger.With("module", fmt.Sprintf("x/%s", types.ModuleName))
}
