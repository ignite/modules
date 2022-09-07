// Package keeper provides methods to initialize SDK keepers with local storage for test purposes
package keeper

import (
	"testing"
	"time"

	distrkeeper "github.com/cosmos/cosmos-sdk/x/distribution/keeper"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	ibcclienttypes "github.com/cosmos/ibc-go/v5/modules/core/02-client/types"
	ibcconnectiontypes "github.com/cosmos/ibc-go/v5/modules/core/03-connection/types"
	ibckeeper "github.com/cosmos/ibc-go/v5/modules/core/keeper"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	claimkeeper "github.com/ignite/modules/x/claim/keeper"
	claimtypes "github.com/ignite/modules/x/claim/types"
)

var (
	// ExampleTimestamp is a timestamp used as the current time for the context of the keepers returned from the package
	ExampleTimestamp = time.Date(2020, time.January, 1, 12, 0, 0, 0, time.UTC)

	// ExampleHeight is a block height used as the current block height for the context of test keeper
	ExampleHeight = int64(1111)
)

// TestKeepers holds all keepers used during keeper tests for all modules
type TestKeepers struct {
	T             testing.TB
	AccountKeeper authkeeper.AccountKeeper
	BankKeeper    bankkeeper.Keeper
	DistrKeeper   distrkeeper.Keeper
	IBCKeeper     *ibckeeper.Keeper
	StakingKeeper stakingkeeper.Keeper
	ClaimKeeper   *claimkeeper.Keeper
}

// TestMsgServers holds all message servers used during keeper tests for all modules
type TestMsgServers struct {
	T        testing.TB
	ClaimSrv claimtypes.MsgServer
}

// NewTestSetup returns initialized instances of all the keepers and message servers of the modules
func NewTestSetup(t testing.TB) (sdk.Context, TestKeepers, TestMsgServers) {
	initializer := newInitializer()

	paramKeeper := initializer.Param()
	capabilityKeeper := initializer.Capability()
	authKeeper := initializer.Auth(paramKeeper)
	bankKeeper := initializer.Bank(paramKeeper, authKeeper)
	stakingKeeper := initializer.Staking(authKeeper, bankKeeper, paramKeeper)
	distrKeeper := initializer.Distribution(authKeeper, bankKeeper, stakingKeeper, paramKeeper)
	upgradeKeeper := initializer.Upgrade()
	ibcKeeper := initializer.IBC(paramKeeper, stakingKeeper, *capabilityKeeper, upgradeKeeper)
	claimKeeper := initializer.Claim(paramKeeper, authKeeper, distrKeeper, bankKeeper)
	require.NoError(t, initializer.StateStore.LoadLatestVersion())

	// Create a context using a custom timestamp
	ctx := sdk.NewContext(initializer.StateStore, tmproto.Header{
		Time:   ExampleTimestamp,
		Height: ExampleHeight,
	}, false, log.NewNopLogger())

	// Initialize community pool
	distrKeeper.SetFeePool(ctx, distrtypes.InitialFeePool())

	// Initialize params
	distrKeeper.SetParams(ctx, distrtypes.DefaultParams())
	stakingKeeper.SetParams(ctx, stakingtypes.DefaultParams())
	claimKeeper.SetParams(ctx, claimtypes.DefaultParams())
	setIBCDefaultParams(ctx, ibcKeeper)

	claimSrv := claimkeeper.NewMsgServerImpl(*claimKeeper)

	return ctx, TestKeepers{
			T:             t,
			AccountKeeper: authKeeper,
			BankKeeper:    bankKeeper,
			DistrKeeper:   distrKeeper,
			IBCKeeper:     ibcKeeper,
			StakingKeeper: stakingKeeper,
			ClaimKeeper:   claimKeeper,
		}, TestMsgServers{
			T:        t,
			ClaimSrv: claimSrv,
		}
}

// setIBCDefaultParams set default params for IBC client and connection keepers
func setIBCDefaultParams(ctx sdk.Context, ibcKeeper *ibckeeper.Keeper) {
	ibcKeeper.ClientKeeper.SetParams(ctx, ibcclienttypes.DefaultParams())
	ibcKeeper.ConnectionKeeper.SetParams(ctx, ibcconnectiontypes.DefaultParams())
	ibcKeeper.ClientKeeper.SetNextClientSequence(ctx, 0)
}
