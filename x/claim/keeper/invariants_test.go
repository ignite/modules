package keeper_test

import (
	"testing"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/ignite/modules/testutil/sample"
	"github.com/ignite/modules/x/claim/keeper"
	"github.com/ignite/modules/x/claim/types"
)

func TestClaimRecordInvariant(t *testing.T) {
	t.Run("should not break with a completed and claimed mission", func(t *testing.T) {
		testSuite := createClaimKeeper(t)
		ctx := testSuite.ctx
		tk := testSuite.tk

		tk.SetMission(ctx, types.Mission{
			MissionID:   10,
			Description: "test mission",
			Weight:      sdkmath.LegacyNewDec(100),
		})
		tk.SetClaimRecord(ctx, types.ClaimRecord{
			Address:           sample.AccAddress(),
			Claimable:         sdkmath.NewInt(10),
			CompletedMissions: []uint64{10},
			ClaimedMissions:   []uint64{10},
		})

		msg, broken := keeper.ClaimRecordInvariant(*tk)(ctx)
		require.False(t, broken, msg)
	})
	t.Run("should not break with a completed but not claimed mission", func(t *testing.T) {
		testSuite := createClaimKeeper(t)
		ctx := testSuite.ctx
		tk := testSuite.tk

		tk.SetMission(ctx, types.Mission{
			MissionID:   10,
			Description: "test mission",
			Weight:      sdkmath.LegacyNewDec(100),
		})
		tk.SetClaimRecord(ctx, types.ClaimRecord{
			Address:           sample.AccAddress(),
			Claimable:         sdkmath.NewInt(10),
			CompletedMissions: []uint64{10},
		})

		msg, broken := keeper.ClaimRecordInvariant(*tk)(ctx)
		require.False(t, broken, msg)
	})
	t.Run("should break with claimed but not completed mission", func(t *testing.T) {
		testSuite := createClaimKeeper(t)
		ctx := testSuite.ctx
		tk := testSuite.tk

		tk.SetClaimRecord(ctx, types.ClaimRecord{
			Address:           sample.AccAddress(),
			Claimable:         sdkmath.NewInt(10),
			CompletedMissions: []uint64{},
			ClaimedMissions:   []uint64{10},
		})
		tk.SetMission(ctx, types.Mission{
			MissionID:   10,
			Description: "test mission",
			Weight:      sdkmath.LegacyNewDec(100),
		})

		msg, broken := keeper.ClaimRecordInvariant(*tk)(ctx)
		require.True(t, broken, msg)
	})
}

func TestClaimRecordMissionInvariant(t *testing.T) {
	t.Run("should not break with valid state", func(t *testing.T) {
		testSuite := createClaimKeeper(t)
		ctx := testSuite.ctx
		tk := testSuite.tk

		tk.SetClaimRecord(ctx, types.ClaimRecord{
			Address:           sample.AccAddress(),
			Claimable:         sdkmath.NewInt(10),
			CompletedMissions: []uint64{0, 1},
		})
		tk.SetMission(ctx, types.Mission{
			MissionID:   0,
			Description: "mission 0",
			Weight:      sdkmath.LegacyZeroDec(),
		})
		tk.SetMission(ctx, types.Mission{
			MissionID:   1,
			Description: "mission 1",
			Weight:      sdkmath.LegacyZeroDec(),
		})

		msg, broken := keeper.ClaimRecordMissionInvariant(*tk)(ctx)
		require.False(t, broken, msg)
	})
	t.Run("should break with invalid state", func(t *testing.T) {
		testSuite := createClaimKeeper(t)
		ctx := testSuite.ctx
		tk := testSuite.tk

		tk.SetClaimRecord(ctx, types.ClaimRecord{
			Address:           sample.AccAddress(),
			Claimable:         sdkmath.NewInt(10),
			CompletedMissions: []uint64{0, 1},
		})
		tk.SetMission(ctx, types.Mission{
			MissionID:   1,
			Description: "mission 1",
			Weight:      sdkmath.LegacyZeroDec(),
		})

		msg, broken := keeper.ClaimRecordMissionInvariant(*tk)(ctx)
		require.True(t, broken, msg)
	})
}

func TestAirdropSupplyInvariant(t *testing.T) {
	t.Run("should not break with valid state", func(t *testing.T) {
		testSuite := createClaimKeeper(t)
		ctx := testSuite.ctx
		tk := testSuite.tk

		tk.InitializeAirdropSupply(ctx, sdk.NewCoin("test", sdkmath.NewInt(10)))
		tk.SetClaimRecord(ctx, types.ClaimRecord{
			Address:           sample.AccAddress(),
			Claimable:         sdkmath.NewInt(10),
			CompletedMissions: nil,
		})

		msg, broken := keeper.AirdropSupplyInvariant(*tk)(ctx)
		require.False(t, broken, msg)
	})

	t.Run("should not break with valid state and completed missions", func(t *testing.T) {
		testSuite := createClaimKeeper(t)
		ctx := testSuite.ctx
		tk := testSuite.tk

		tk.InitializeAirdropSupply(ctx, sdk.NewCoin("test", sdkmath.NewInt(10)))
		tk.SetClaimRecord(ctx, types.ClaimRecord{
			Address:           sample.AccAddress(),
			Claimable:         sdkmath.NewInt(10),
			CompletedMissions: []uint64{0, 1},
		})
		tk.SetMission(ctx, types.Mission{
			MissionID:   0,
			Description: "",
			Weight:      sdkmath.LegacyZeroDec(),
		})
		tk.SetMission(ctx, types.Mission{
			MissionID:   1,
			Description: "",
			Weight:      sdkmath.LegacyZeroDec(),
		})

		msg, broken := keeper.AirdropSupplyInvariant(*tk)(ctx)
		require.False(t, broken, msg)
	})

	t.Run("should break with duplicated address in claim record", func(t *testing.T) {
		testSuite := createClaimKeeper(t)
		ctx := testSuite.ctx
		tk := testSuite.tk

		tk.InitializeAirdropSupply(ctx, sdk.NewCoin("test", sdkmath.NewInt(10)))
		addr := sample.AccAddress()
		tk.SetClaimRecord(ctx, types.ClaimRecord{
			Address:           addr,
			Claimable:         sdkmath.NewInt(5),
			CompletedMissions: nil,
		})
		tk.SetClaimRecord(ctx, types.ClaimRecord{
			Address:           addr,
			Claimable:         sdkmath.NewInt(5),
			CompletedMissions: nil,
		})

		msg, broken := keeper.AirdropSupplyInvariant(*tk)(ctx)
		require.True(t, broken, msg)
	})

	t.Run("should break with address completing non existing mission", func(t *testing.T) {
		testSuite := createClaimKeeper(t)
		ctx := testSuite.ctx
		tk := testSuite.tk

		tk.InitializeAirdropSupply(ctx, sdk.NewCoin("test", sdkmath.NewInt(10)))
		tk.SetClaimRecord(ctx, types.ClaimRecord{
			Address:           sample.AccAddress(),
			Claimable:         sdkmath.NewInt(10),
			CompletedMissions: []uint64{0, 1, 2},
		})

		msg, broken := keeper.AirdropSupplyInvariant(*tk)(ctx)
		require.True(t, broken, msg)
	})

	t.Run("should break with airdrop supply not equal to claimable amounts", func(t *testing.T) {
		testSuite := createClaimKeeper(t)
		ctx := testSuite.ctx
		tk := testSuite.tk

		tk.InitializeAirdropSupply(ctx, sdk.NewCoin("test", sdkmath.NewInt(10)))
		tk.SetClaimRecord(ctx, types.ClaimRecord{
			Address:           sample.AccAddress(),
			Claimable:         sdkmath.NewInt(9),
			CompletedMissions: nil,
		})

		msg, broken := keeper.AirdropSupplyInvariant(*tk)(ctx)
		require.True(t, broken, msg)
	})
}
