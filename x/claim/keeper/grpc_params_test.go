package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	testkeeper "github.com/ignite/modules/testutil/keeper"
	"github.com/ignite/modules/x/claim/types"
)

func TestParamsQuery(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)

	t.Run("should allow params get query", func(t *testing.T) {
		params := types.DefaultParams()
		tk.ClaimKeeper.SetParams(ctx, params)

		response, err := tk.ClaimKeeper.Params(ctx, &types.QueryParamsRequest{})
		require.NoError(t, err)
		require.EqualValues(t, params.DecayInformation, response.Params.DecayInformation)
		require.Equal(t, params.AirdropStart.Unix(), response.Params.AirdropStart.Unix())
	})
}
