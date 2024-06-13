package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/ignite/modules/x/claim/types"
)

func TestParamsQuery(t *testing.T) {
	ctx, tk := createClaimKeeper(t)

	t.Run("should allow params get query", func(t *testing.T) {
		params := types.DefaultParams()
		tk.SetParams(ctx, params)

		response, err := tk.Params(ctx, &types.QueryParamsRequest{})
		require.NoError(t, err)
		require.EqualValues(t, params.DecayInformation, response.Params.DecayInformation)
		require.Equal(t, params.AirdropStart.Unix(), response.Params.AirdropStart.Unix())
	})
}
