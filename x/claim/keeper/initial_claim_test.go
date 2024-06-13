package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/ignite/modules/testutil/nullify"
	"github.com/ignite/modules/x/claim/keeper"
	"github.com/ignite/modules/x/claim/types"
)

func createTestInitialClaim(keeper *keeper.Keeper, ctx sdk.Context) types.InitialClaim {
	item := types.InitialClaim{}
	keeper.SetInitialClaim(ctx, item)
	return item
}

func TestInitialClaimGet(t *testing.T) {
	ctx, tk := createClaimKeeper(t)

	t.Run("should allow get", func(t *testing.T) {
		item := createTestInitialClaim(tk, ctx)
		rst, found := tk.GetInitialClaim(ctx)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	})
}

func TestInitialClaimRemove(t *testing.T) {
	ctx, tk := createClaimKeeper(t)

	t.Run("should allow remove", func(t *testing.T) {
		createTestInitialClaim(tk, ctx)
		tk.RemoveInitialClaim(ctx)
		_, found := tk.GetInitialClaim(ctx)
		require.False(t, found)
	})
}
