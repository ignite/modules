package keeper_test

import (
	"testing"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/ignite/modules/testutil/nullify"
	"github.com/ignite/modules/testutil/sample"
	"github.com/ignite/modules/x/claim/keeper"
	"github.com/ignite/modules/x/claim/types"
)

func createNClaimRecord(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.ClaimRecord {
	items := make([]types.ClaimRecord, n)
	for i := range items {
		items[i].Address = sample.AccAddress()
		items[i].Claimable = sdkmath.NewInt(r.Int63())

		keeper.SetClaimRecord(ctx, items[i])
	}
	return items
}

func TestClaimRecordGet(t *testing.T) {
	ctx, tk := createClaimKeeper(t)

	t.Run("should allow get", func(t *testing.T) {
		items := createNClaimRecord(tk, ctx, 10)
		for _, item := range items {
			rst, found := tk.GetClaimRecord(ctx,
				item.Address,
			)
			require.True(t, found)
			require.Equal(t,
				nullify.Fill(&item),
				nullify.Fill(&rst),
			)
		}
	})
}

func TestClaimRecordRemove(t *testing.T) {
	ctx, tk := createClaimKeeper(t)

	t.Run("should allow remove", func(t *testing.T) {
		items := createNClaimRecord(tk, ctx, 10)
		for _, item := range items {
			tk.RemoveClaimRecord(ctx,
				item.Address,
			)
			_, found := tk.GetClaimRecord(ctx,
				item.Address,
			)
			require.False(t, found)
		}
	})
}

func TestClaimRecordGetAll(t *testing.T) {
	ctx, tk := createClaimKeeper(t)

	t.Run("should allow get all", func(t *testing.T) {
		items := createNClaimRecord(tk, ctx, 10)
		require.ElementsMatch(t,
			nullify.Fill(items),
			nullify.Fill(tk.GetAllClaimRecord(ctx)),
		)
	})
}
