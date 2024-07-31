package keeper

import (
	"cosmossdk.io/log"
	"cosmossdk.io/store"
	"cosmossdk.io/store/metrics"
	storetypes "cosmossdk.io/store/types"
	upgradekeeper "cosmossdk.io/x/upgrade/keeper"
	upgradetypes "cosmossdk.io/x/upgrade/types"
	dbm "github.com/cosmos/cosmos-db"
	"github.com/cosmos/cosmos-sdk/codec"
	addresscodec "github.com/cosmos/cosmos-sdk/codec/address"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	distrkeeper "github.com/cosmos/cosmos-sdk/x/distribution/keeper"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	paramskeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	capabilitykeeper "github.com/cosmos/ibc-go/modules/capability/keeper"
	capabilitytypes "github.com/cosmos/ibc-go/modules/capability/types"

	"github.com/ignite/modules/testutil/sample"
	claimkeeper "github.com/ignite/modules/x/claim/keeper"

	claimtypes "github.com/ignite/modules/x/claim/types"
	minttypes "github.com/ignite/modules/x/mint/types"
)

var moduleAccountPerms = map[string][]string{
	authtypes.FeeCollectorName:     nil,
	distrtypes.ModuleName:          nil,
	stakingtypes.BondedPoolName:    {authtypes.Burner, authtypes.Staking},
	stakingtypes.NotBondedPoolName: {authtypes.Burner, authtypes.Staking},
	minttypes.ModuleName:           {authtypes.Minter},
	claimtypes.ModuleName:          {authtypes.Minter, authtypes.Burner},
}

// initializer allows to initialize each module keeper
type initializer struct {
	Codec      codec.Codec
	DB         *dbm.MemDB
	StateStore store.CommitMultiStore
}

func newInitializer() initializer {
	cdc := sample.Codec()
	db := dbm.NewMemDB()
	stateStore := store.NewCommitMultiStore(db, log.NewNopLogger(), metrics.NewNoOpMetrics())

	return initializer{
		Codec:      cdc,
		DB:         db,
		StateStore: stateStore,
	}
}

// ModuleAccountAddrs returns all the app's module account addresses.
func ModuleAccountAddrs(maccPerms map[string][]string) map[string]bool {
	modAccAddrs := make(map[string]bool)
	for acc := range maccPerms {
		modAccAddrs[authtypes.NewModuleAddress(acc).String()] = true
	}

	return modAccAddrs
}

func (i initializer) Param() paramskeeper.Keeper {
	storeKey := storetypes.NewKVStoreKey(paramstypes.StoreKey)
	tkeys := storetypes.NewTransientStoreKey(paramstypes.TStoreKey)

	i.StateStore.MountStoreWithDB(storeKey, storetypes.StoreTypeIAVL, i.DB)
	i.StateStore.MountStoreWithDB(tkeys, storetypes.StoreTypeTransient, i.DB)

	return paramskeeper.NewKeeper(
		i.Codec,
		codec.NewLegacyAmino(),
		storeKey,
		tkeys,
	)
}

func (i initializer) Auth() authkeeper.AccountKeeper {
	storeKey := storetypes.NewKVStoreKey(authtypes.StoreKey)
	i.StateStore.MountStoreWithDB(storeKey, storetypes.StoreTypeIAVL, i.DB)

	return authkeeper.NewAccountKeeper(
		i.Codec,
		runtime.NewKVStoreService(storeKey),
		authtypes.ProtoBaseAccount,
		moduleAccountPerms,
		addresscodec.NewBech32Codec(sdk.Bech32PrefixAccAddr),
		sdk.Bech32PrefixAccAddr,
		authtypes.NewModuleAddress(govtypes.ModuleName).String(),
	)
}

func (i initializer) Bank(authKeeper authkeeper.AccountKeeper) bankkeeper.Keeper {
	storeKey := storetypes.NewKVStoreKey(banktypes.StoreKey)
	i.StateStore.MountStoreWithDB(storeKey, storetypes.StoreTypeIAVL, i.DB)
	modAccAddrs := ModuleAccountAddrs(moduleAccountPerms)

	return bankkeeper.NewBaseKeeper(
		i.Codec,
		runtime.NewKVStoreService(storeKey),
		authKeeper,
		modAccAddrs,
		authtypes.NewModuleAddress(govtypes.ModuleName).String(),
		log.NewNopLogger(),
	)
}

func (i initializer) Capability() *capabilitykeeper.Keeper {
	storeKey := storetypes.NewKVStoreKey(capabilitytypes.StoreKey)
	memStoreKey := storetypes.NewMemoryStoreKey(capabilitytypes.MemStoreKey)

	i.StateStore.MountStoreWithDB(storeKey, storetypes.StoreTypeIAVL, i.DB)
	i.StateStore.MountStoreWithDB(memStoreKey, storetypes.StoreTypeMemory, i.DB)

	return capabilitykeeper.NewKeeper(i.Codec, storeKey, memStoreKey)
}

// create mock ProtocolVersionSetter for UpgradeKeeper

type ProtocolVersionSetter struct{}

func (vs ProtocolVersionSetter) SetProtocolVersion(uint64) {}

func (i initializer) Upgrade() *upgradekeeper.Keeper {
	storeKey := storetypes.NewKVStoreKey(upgradetypes.StoreKey)
	i.StateStore.MountStoreWithDB(storeKey, storetypes.StoreTypeIAVL, i.DB)

	skipUpgradeHeights := make(map[int64]bool)
	vs := ProtocolVersionSetter{}

	return upgradekeeper.NewKeeper(
		skipUpgradeHeights,
		runtime.NewKVStoreService(storeKey),
		i.Codec,
		"",
		vs,
		authtypes.NewModuleAddress(govtypes.ModuleName).String(),
	)
}

func (i initializer) Staking(
	authKeeper authkeeper.AccountKeeper,
	bankKeeper bankkeeper.Keeper,
) *stakingkeeper.Keeper {
	storeKey := storetypes.NewKVStoreKey(stakingtypes.StoreKey)
	i.StateStore.MountStoreWithDB(storeKey, storetypes.StoreTypeIAVL, i.DB)

	return stakingkeeper.NewKeeper(
		i.Codec,
		runtime.NewKVStoreService(storeKey),
		authKeeper,
		bankKeeper,
		authtypes.NewModuleAddress(govtypes.ModuleName).String(),
		addresscodec.NewBech32Codec(sdk.Bech32PrefixValAddr),
		addresscodec.NewBech32Codec(sdk.Bech32PrefixConsAddr),
	)
}

func (i initializer) Distribution(
	authKeeper authkeeper.AccountKeeper,
	bankKeeper bankkeeper.Keeper,
	stakingKeeper *stakingkeeper.Keeper,
) distrkeeper.Keeper {
	storeKey := storetypes.NewKVStoreKey(distrtypes.StoreKey)
	i.StateStore.MountStoreWithDB(storeKey, storetypes.StoreTypeIAVL, i.DB)

	return distrkeeper.NewKeeper(
		i.Codec,
		runtime.NewKVStoreService(storeKey),
		authKeeper,
		bankKeeper,
		stakingKeeper,
		authtypes.FeeCollectorName,
		authtypes.NewModuleAddress(govtypes.ModuleName).String(),
	)
}

func (i initializer) Claim(
	accountKeeper authkeeper.AccountKeeper,
	distrKeeper distrkeeper.Keeper,
	bankKeeper bankkeeper.Keeper,
) claimkeeper.Keeper {
	storeKey := storetypes.NewKVStoreKey(claimtypes.StoreKey)
	i.StateStore.MountStoreWithDB(storeKey, storetypes.StoreTypeIAVL, i.DB)

	addressCodec := addresscodec.NewBech32Codec(sdk.GetConfig().GetBech32AccountAddrPrefix())
	authority := authtypes.NewModuleAddress(govtypes.ModuleName)

	return claimkeeper.NewKeeper(
		i.Codec,
		addressCodec,
		runtime.NewKVStoreService(storeKey),
		log.NewNopLogger(),
		authority.String(),
		accountKeeper,
		bankKeeper,
		distrKeeper,
	)
}
