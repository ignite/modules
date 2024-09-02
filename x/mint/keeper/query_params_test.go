package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "github.com/ignite/modules/testutil/keeper"
	"github.com/ignite/modules/x/mint/keeper"
	"github.com/ignite/modules/x/mint/types"
)

func TestParamsQuery(t *testing.T) {
	k, ctx, _ := keepertest.MintKeeper(t)

	qs := keeper.NewQueryServerImpl(k)
	params := types.DefaultParams()
	require.NoError(t, k.Params.Set(ctx, params))

	response, err := qs.Params(ctx, &types.QueryParamsRequest{})
	require.NoError(t, err)
	require.Equal(t, &types.QueryParamsResponse{Params: params}, response)
}
