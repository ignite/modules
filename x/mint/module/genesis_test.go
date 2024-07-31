package mint_test

import (
	"testing"

	sdkmath "cosmossdk.io/math"
	keepertest "github.com/ignite/modules/testutil/keeper"
	"github.com/ignite/modules/testutil/nullify"
	mint "github.com/ignite/modules/x/mint/module"
	"github.com/ignite/modules/x/mint/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		Minter: types.Minter{
			Inflation:        sdkmath.LegacyNewDec(47),
			AnnualProvisions: sdkmath.LegacyNewDec(58),
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx, _ := keepertest.MintKeeper(t)
	mint.InitGenesis(ctx, k, genesisState)
	got := mint.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.Equal(t, genesisState.Minter, got.Minter)
	// this line is used by starport scaffolding # genesis/test/assert
}
