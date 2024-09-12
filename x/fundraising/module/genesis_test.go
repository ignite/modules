package fundraising_test

import (
	"encoding/json"
	"math/rand"
	"testing"

	"cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/stretchr/testify/require"

	keepertest "github.com/ignite/modules/testutil/keeper"
	"github.com/ignite/modules/testutil/nullify"
	"github.com/ignite/modules/testutil/sample"
	fundraising "github.com/ignite/modules/x/fundraising/module"
	"github.com/ignite/modules/x/fundraising/types"
)

var r *rand.Rand

// initialize random generator
func init() {
	s := rand.NewSource(1)
	r = rand.New(s)
}

func TestGenesis(t *testing.T) {
	auctionAny1, _ := types.PackAuction(types.NewFixedPriceAuction(
		&types.BaseAuction{AuctionID: 1},
		sdk.NewInt64Coin("denom1", 1_000),
	))
	auctionAny2, _ := types.PackAuction(types.NewFixedPriceAuction(
		&types.BaseAuction{AuctionID: 2},
		sdk.NewInt64Coin("denom2", 2_000),
	))

	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		AllowedBidderList: []types.AllowedBidder{
			{
				AuctionID: 0,
				Bidder:    sample.Address(r),
			},
			{
				AuctionID: 1,
				Bidder:    sample.Address(r),
			},
		},
		VestingQueueList: []types.VestingQueue{
			{
				AuctionID: 0,
			},
			{
				AuctionID: 1,
			},
		},
		BidList: []types.Bid{
			{
				BidID: 0,
			},
			{
				BidID: 1,
			},
		},
		AuctionList: []*codectypes.Any{auctionAny1, auctionAny2},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx, _ := keepertest.FundraisingKeeper(t)
	err := fundraising.InitGenesis(ctx, k, genesisState)
	require.NoError(t, err)
	got, err := fundraising.ExportGenesis(ctx, k)
	require.NoError(t, err)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.AllowedBidderList, got.AllowedBidderList)
	require.ElementsMatch(t, genesisState.VestingQueueList, got.VestingQueueList)
	require.ElementsMatch(t, genesisState.BidList, got.BidList)
	require.ElementsMatch(t, genesisState.AuctionList, got.AuctionList)
	// this line is used by starport scaffolding # genesis/test/assert
}

// TestRandomizedGenState tests the normal scenario of applying RandomizedGenState.
// Abnormal scenarios are not tested here.
func TestRandomizedGenState(t *testing.T) {
	interfaceRegistry := codectypes.NewInterfaceRegistry()
	cdc := codec.NewProtoCodec(interfaceRegistry)
	s := rand.NewSource(1)
	r := rand.New(s)

	simState := module.SimulationState{
		AppParams:    make(simtypes.AppParams),
		Cdc:          cdc,
		Rand:         r,
		NumBonded:    3,
		Accounts:     simtypes.RandomAccounts(r, 3),
		InitialStake: math.NewInt(1000),
		GenState:     make(map[string]json.RawMessage),
	}

	fundraising.RandomizedGenState(&simState)

	var genState types.GenesisState
	simState.Cdc.MustUnmarshalJSON(simState.GenState[types.ModuleName], &genState)

	dec1 := sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, math.NewInt(36122540)))
	dec3 := uint32(5)

	require.Equal(t, dec1, genState.Params.AuctionCreationFee)
	require.Equal(t, dec3, genState.Params.ExtendedPeriod)
}

// TestRandomizedGenState tests abnormal scenarios of applying RandomizedGenState.
func TestRandomizedGenState1(t *testing.T) {
	interfaceRegistry := codectypes.NewInterfaceRegistry()
	cdc := codec.NewProtoCodec(interfaceRegistry)

	s := rand.NewSource(1)
	r := rand.New(s)

	// all these tests will panic
	tests := []struct {
		simState module.SimulationState
		panicMsg string
	}{
		{ // panic => reason: incomplete initialization of the simState
			module.SimulationState{}, "invalid memory address or nil pointer dereference"},
		{ // panic => reason: incomplete initialization of the simState
			module.SimulationState{
				AppParams: make(simtypes.AppParams),
				Cdc:       cdc,
				Rand:      r,
			}, "assignment to entry in nil map"},
	}

	for _, tt := range tests {
		require.Panicsf(t, func() { fundraising.RandomizedGenState(&tt.simState) }, tt.panicMsg)
	}
}
