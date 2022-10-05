package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/ignite/modules/x/mint/types"
)

func TestValidateGenesis(t *testing.T) {
	invalid := types.DefaultGenesis()
	// set inflation min to larger than inflation max
	invalid.Params.InflationMin = invalid.Params.InflationMax.Add(invalid.Params.InflationMax)

	tests := []struct {
		name    string
		genesis *types.GenesisState
		isValid bool
	}{
		{
			name:    "should validate valid genesis",
			genesis: types.DefaultGenesis(),
			isValid: true,
		},
		{
			name:    "should prevent invalid params",
			genesis: invalid,
			isValid: false,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.genesis.Validate()
			if !tc.isValid {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
		})
	}
}
