package types_test

import (
	"testing"

	sdkmath "cosmossdk.io/math"
	"github.com/stretchr/testify/require"

	"github.com/ignite/modules/x/mint/types"
)

func TestGenesisState_Validate(t *testing.T) {
	invalid := types.DefaultGenesis()
	// set inflation min to larger than inflation max
	invalid.Params.InflationMin = invalid.Params.InflationMax.Add(invalid.Params.InflationMax)

	tests := []struct {
		desc     string
		genState *types.GenesisState
		valid    bool
	}{
		{
			desc:     "default is valid",
			genState: types.DefaultGenesis(),
			valid:    true,
		},
		{
			desc:     "should prevent invalid params",
			genState: invalid,
			valid:    false,
		},
		{
			desc: "valid genesis state",
			genState: &types.GenesisState{
				Params: types.DefaultParams(),
				Minter: types.Minter{
					Inflation:        sdkmath.LegacyNewDec(80),
					AnnualProvisions: sdkmath.LegacyNewDec(22),
				},
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		// this line is used by starport scaffolding # types/genesis/testcase
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.genState.Validate()
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}
