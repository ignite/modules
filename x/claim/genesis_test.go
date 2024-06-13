package claim_test

import (
	"testing"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/ignite/modules/testutil/nullify"
	"github.com/ignite/modules/testutil/sample"
	"github.com/ignite/modules/x/claim"
	"github.com/ignite/modules/x/claim/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		ClaimRecords: []types.ClaimRecord{
			{
				Address: sample.AccAddress(),
			},
			{
				Address: sample.AccAddress(),
			},
		},
		Missions: []types.Mission{
			{
				MissionID: 0,
			},
			{
				MissionID: 1,
			},
		},
		AirdropSupply: sdk.NewCoin("foo", sdkmath.NewInt(1000)),
		InitialClaim: types.InitialClaim{
			Enabled:   true,
			MissionID: 35,
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	ctx, tk := createClaimKeeper(t)

	t.Run("should allow import and export of genesis", func(t *testing.T) {
		claim.InitGenesis(ctx, *tk, genesisState)
		got := claim.ExportGenesis(ctx, *tk)
		require.NotNil(t, got)

		nullify.Fill(&genesisState)
		nullify.Fill(got)

		require.ElementsMatch(t, genesisState.ClaimRecords, got.ClaimRecords)
		require.ElementsMatch(t, genesisState.Missions, got.Missions)
		require.Equal(t, genesisState.AirdropSupply, got.AirdropSupply)
		require.Equal(t, genesisState.InitialClaim, got.InitialClaim)
		// this line is used by starport scaffolding # genesis/test/assert
	})
}
