package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/ignite/modules/x/mint/types"
)

func TestValidateGenesis(t *testing.T) {
	tests := []struct {
		name    string
		genesis *types.GenesisState
		err     error
	}{
		{
			name:    "should validate valid genesis",
			genesis: types.DefaultGenesisState(),
			err:     nil,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := types.ValidateGenesis(*tc.genesis)
			if tc.err != nil {
				require.Error(t, err, tc.err)
				require.Equal(t, err, tc.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
