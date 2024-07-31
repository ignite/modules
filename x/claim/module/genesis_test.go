package claim_test

import (
	"testing"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	keepertest "github.com/ignite/modules/testutil/keeper"
	"github.com/ignite/modules/testutil/nullify"
	claim "github.com/ignite/modules/x/claim/module"
	"github.com/ignite/modules/x/claim/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		ClaimRecordList: []types.ClaimRecord{
			{
				Address: "0",
			},
			{
				Address: "1",
			},
		},
		MissionList: []types.Mission{
			{
				MissionID: 0,
			},
			{
				MissionID: 1,
			},
		},
		MissionCount: 2,
		InitialClaim: types.InitialClaim{
			Enabled:   true,
			MissionID: 64,
		},
		AirdropSupply: types.AirdropSupply{
			Supply: sdk.NewCoin(sdk.DefaultBondDenom, sdkmath.NewInt(20)),
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx, _ := keepertest.ClaimKeeper(t)
	claim.InitGenesis(ctx, k, genesisState)
	got := claim.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.ClaimRecordList, got.ClaimRecordList)
	require.ElementsMatch(t, genesisState.MissionList, got.MissionList)
	require.Equal(t, genesisState.MissionCount, got.MissionCount)
	require.Equal(t, genesisState.InitialClaim, got.InitialClaim)
	require.Equal(t, genesisState.AirdropSupply, got.AirdropSupply)
	// this line is used by starport scaffolding # genesis/test/assert
}
