package keeper_test

import (
	"testing"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	testkeeper "github.com/ignite/modules/testutil/keeper"
	"github.com/ignite/modules/testutil/sample"
	"github.com/ignite/modules/x/claim/keeper"
	"github.com/ignite/modules/x/claim/types"
)

func TestAirdropSupplyInvariant(t *testing.T) {
	t.Run("should not break with valid state", func(t *testing.T) {
		ctx, tk, _ := testkeeper.NewTestSetup(t)

		tk.ClaimKeeper.InitializeAirdropSupply(ctx, sdk.NewCoin("test", sdkmath.NewInt(10)))
		tk.ClaimKeeper.SetClaimRecord(ctx, types.ClaimRecord{
			Address:           sample.Address(r),
			Claimable:         sdkmath.NewInt(10),
			CompletedMissions: nil,
		})

		msg, broken := keeper.AirdropSupplyInvariant(*tk.ClaimKeeper)(ctx)
		require.False(t, broken, msg)
	})

	t.Run("should not break with valid state and completed missions", func(t *testing.T) {
		ctx, tk, _ := testkeeper.NewTestSetup(t)

		tk.ClaimKeeper.InitializeAirdropSupply(ctx, sdk.NewCoin("test", sdkmath.NewInt(10)))
		tk.ClaimKeeper.SetClaimRecord(ctx, types.ClaimRecord{
			Address:           sample.Address(r),
			Claimable:         sdkmath.NewInt(10),
			CompletedMissions: []uint64{0, 1},
		})
		tk.ClaimKeeper.SetMission(ctx, types.Mission{
			MissionID:   0,
			Description: "",
			Weight:      sdk.ZeroDec(),
		})
		tk.ClaimKeeper.SetMission(ctx, types.Mission{
			MissionID:   1,
			Description: "",
			Weight:      sdk.ZeroDec(),
		})

		msg, broken := keeper.AirdropSupplyInvariant(*tk.ClaimKeeper)(ctx)
		require.False(t, broken, msg)
	})

	t.Run("should break with duplicated address in claim record", func(t *testing.T) {
		ctx, tk, _ := testkeeper.NewTestSetup(t)

		tk.ClaimKeeper.InitializeAirdropSupply(ctx, sdk.NewCoin("test", sdkmath.NewInt(10)))
		addr := sample.Address(r)
		tk.ClaimKeeper.SetClaimRecord(ctx, types.ClaimRecord{
			Address:           addr,
			Claimable:         sdkmath.NewInt(5),
			CompletedMissions: nil,
		})
		tk.ClaimKeeper.SetClaimRecord(ctx, types.ClaimRecord{
			Address:           addr,
			Claimable:         sdkmath.NewInt(5),
			CompletedMissions: nil,
		})

		msg, broken := keeper.AirdropSupplyInvariant(*tk.ClaimKeeper)(ctx)
		require.True(t, broken, msg)
	})

	t.Run("should break with address completing non existing mission", func(t *testing.T) {
		ctx, tk, _ := testkeeper.NewTestSetup(t)

		tk.ClaimKeeper.InitializeAirdropSupply(ctx, sdk.NewCoin("test", sdkmath.NewInt(10)))
		tk.ClaimKeeper.SetClaimRecord(ctx, types.ClaimRecord{
			Address:           sample.Address(r),
			Claimable:         sdkmath.NewInt(10),
			CompletedMissions: []uint64{0, 1, 2},
		})

		msg, broken := keeper.AirdropSupplyInvariant(*tk.ClaimKeeper)(ctx)
		require.True(t, broken, msg)
	})

	t.Run("should break with airdrop supply not equal to claimable amounts", func(t *testing.T) {
		ctx, tk, _ := testkeeper.NewTestSetup(t)

		tk.ClaimKeeper.InitializeAirdropSupply(ctx, sdk.NewCoin("test", sdkmath.NewInt(10)))
		tk.ClaimKeeper.SetClaimRecord(ctx, types.ClaimRecord{
			Address:           sample.Address(r),
			Claimable:         sdkmath.NewInt(9),
			CompletedMissions: nil,
		})

		msg, broken := keeper.AirdropSupplyInvariant(*tk.ClaimKeeper)(ctx)
		require.True(t, broken, msg)
	})
}
