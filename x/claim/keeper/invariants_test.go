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

func TestClaimRecordInvariant(t *testing.T) {
	t.Run("should not break with a completed and claimed mission", func(t *testing.T) {
		ctx, tk, _ := testkeeper.NewTestSetup(t)

		missionID := uint64(10)
		err := tk.ClaimKeeper.Mission.Set(ctx, missionID, types.Mission{
			MissionID:   missionID,
			Description: "test mission",
			Weight:      sdkmath.LegacyNewDec(int64(missionID)),
		})
		require.NoError(t, err)

		address := sample.Address(r)
		err = tk.ClaimKeeper.ClaimRecord.Set(ctx, address, types.ClaimRecord{
			Address:           address,
			Claimable:         sdkmath.NewInt(10),
			CompletedMissions: []uint64{missionID},
			ClaimedMissions:   []uint64{missionID},
		})
		require.NoError(t, err)

		msg, broken := keeper.ClaimRecordInvariant(*tk.ClaimKeeper)(ctx)
		require.False(t, broken, msg)
	})
	t.Run("should not break with a completed but not claimed mission", func(t *testing.T) {
		ctx, tk, _ := testkeeper.NewTestSetup(t)

		missionID := uint64(10)
		err := tk.ClaimKeeper.Mission.Set(ctx, missionID, types.Mission{
			MissionID:   missionID,
			Description: "test mission",
			Weight:      sdkmath.LegacyNewDec(100),
		})
		require.NoError(t, err)

		address := sample.Address(r)
		err = tk.ClaimKeeper.ClaimRecord.Set(ctx, address, types.ClaimRecord{
			Address:           address,
			Claimable:         sdkmath.NewInt(10),
			CompletedMissions: []uint64{missionID},
		})
		require.NoError(t, err)

		msg, broken := keeper.ClaimRecordInvariant(*tk.ClaimKeeper)(ctx)
		require.False(t, broken, msg)
	})
	t.Run("should break with claimed but not completed mission", func(t *testing.T) {
		ctx, tk, _ := testkeeper.NewTestSetup(t)

		address := sample.Address(r)
		err := tk.ClaimKeeper.ClaimRecord.Set(ctx, address, types.ClaimRecord{
			Address:           sample.Address(r),
			Claimable:         sdkmath.NewInt(10),
			CompletedMissions: []uint64{},
			ClaimedMissions:   []uint64{10},
		})
		require.NoError(t, err)

		missionID := uint64(10)
		err = tk.ClaimKeeper.Mission.Set(ctx, missionID, types.Mission{
			MissionID:   missionID,
			Description: "test mission",
			Weight:      sdkmath.LegacyNewDec(100),
		})
		require.NoError(t, err)

		msg, broken := keeper.ClaimRecordInvariant(*tk.ClaimKeeper)(ctx)
		require.True(t, broken, msg)
	})
}

func TestClaimRecordMissionInvariant(t *testing.T) {
	t.Run("should not break with valid state", func(t *testing.T) {
		ctx, tk, _ := testkeeper.NewTestSetup(t)

		address := sample.Address(r)
		err := tk.ClaimKeeper.ClaimRecord.Set(ctx, address, types.ClaimRecord{
			Address:           address,
			Claimable:         sdkmath.NewInt(10),
			CompletedMissions: []uint64{0, 1},
		})
		require.NoError(t, err)

		missionID := uint64(0)
		err = tk.ClaimKeeper.Mission.Set(ctx, missionID, types.Mission{
			MissionID:   missionID,
			Description: "mission 0",
			Weight:      sdkmath.LegacyZeroDec(),
		})
		require.NoError(t, err)

		missionID = uint64(1)
		err = tk.ClaimKeeper.Mission.Set(ctx, missionID, types.Mission{
			MissionID:   missionID,
			Description: "mission 1",
			Weight:      sdkmath.LegacyZeroDec(),
		})
		require.NoError(t, err)

		msg, broken := keeper.ClaimRecordMissionInvariant(*tk.ClaimKeeper)(ctx)
		require.False(t, broken, msg)
	})
	t.Run("should break with invalid state", func(t *testing.T) {
		ctx, tk, _ := testkeeper.NewTestSetup(t)

		address := sample.Address(r)
		err := tk.ClaimKeeper.ClaimRecord.Set(ctx, address, types.ClaimRecord{
			Address:           address,
			Claimable:         sdkmath.NewInt(10),
			CompletedMissions: []uint64{0, 1},
		})
		require.NoError(t, err)

		missionID := uint64(1)
		err = tk.ClaimKeeper.Mission.Set(ctx, missionID, types.Mission{
			MissionID:   missionID,
			Description: "mission 1",
			Weight:      sdkmath.LegacyZeroDec(),
		})
		require.NoError(t, err)

		msg, broken := keeper.ClaimRecordMissionInvariant(*tk.ClaimKeeper)(ctx)
		require.True(t, broken, msg)
	})
}

func TestAirdropSupplyInvariant(t *testing.T) {
	t.Run("should not break with valid state", func(t *testing.T) {
		ctx, tk, _ := testkeeper.NewTestSetup(t)

		err := tk.ClaimKeeper.InitializeAirdropSupply(ctx, sdk.NewCoin("test", sdkmath.NewInt(10)))
		require.NoError(t, err)

		address := sample.Address(r)
		err = tk.ClaimKeeper.ClaimRecord.Set(ctx, address, types.ClaimRecord{
			Address:           address,
			Claimable:         sdkmath.NewInt(10),
			CompletedMissions: nil,
		})
		require.NoError(t, err)

		msg, broken := keeper.AirdropSupplyInvariant(*tk.ClaimKeeper)(ctx)
		require.False(t, broken, msg)
	})

	t.Run("should not break with valid state and completed missions", func(t *testing.T) {
		ctx, tk, _ := testkeeper.NewTestSetup(t)

		err := tk.ClaimKeeper.InitializeAirdropSupply(ctx, sdk.NewCoin("test", sdkmath.NewInt(10)))
		require.NoError(t, err)

		address := sample.Address(r)
		err = tk.ClaimKeeper.ClaimRecord.Set(ctx, address, types.ClaimRecord{
			Address:           address,
			Claimable:         sdkmath.NewInt(10),
			CompletedMissions: []uint64{0, 1},
		})
		require.NoError(t, err)

		missionID := uint64(0)
		err = tk.ClaimKeeper.Mission.Set(ctx, missionID, types.Mission{
			MissionID:   missionID,
			Description: "",
			Weight:      sdkmath.LegacyZeroDec(),
		})
		require.NoError(t, err)

		missionID = uint64(1)
		err = tk.ClaimKeeper.Mission.Set(ctx, missionID, types.Mission{
			MissionID:   missionID,
			Description: "",
			Weight:      sdkmath.LegacyZeroDec(),
		})
		require.NoError(t, err)

		msg, broken := keeper.AirdropSupplyInvariant(*tk.ClaimKeeper)(ctx)
		require.False(t, broken, msg)
	})

	t.Run("should break with duplicated address in claim record", func(t *testing.T) {
		ctx, tk, _ := testkeeper.NewTestSetup(t)

		err := tk.ClaimKeeper.InitializeAirdropSupply(ctx, sdk.NewCoin("test", sdkmath.NewInt(10)))
		require.NoError(t, err)

		addr := sample.Address(r)
		err = tk.ClaimKeeper.ClaimRecord.Set(ctx, addr, types.ClaimRecord{
			Address:           addr,
			Claimable:         sdkmath.NewInt(5),
			CompletedMissions: nil,
		})
		require.NoError(t, err)

		err = tk.ClaimKeeper.ClaimRecord.Set(ctx, addr, types.ClaimRecord{
			Address:           addr,
			Claimable:         sdkmath.NewInt(5),
			CompletedMissions: nil,
		})
		require.NoError(t, err)

		msg, broken := keeper.AirdropSupplyInvariant(*tk.ClaimKeeper)(ctx)
		require.True(t, broken, msg)
	})

	t.Run("should break with address completing non existing mission", func(t *testing.T) {
		ctx, tk, _ := testkeeper.NewTestSetup(t)

		err := tk.ClaimKeeper.InitializeAirdropSupply(ctx, sdk.NewCoin("test", sdkmath.NewInt(10)))
		require.NoError(t, err)

		address := sample.Address(r)
		err = tk.ClaimKeeper.ClaimRecord.Set(ctx, address, types.ClaimRecord{
			Address:           address,
			Claimable:         sdkmath.NewInt(10),
			CompletedMissions: []uint64{0, 1, 2},
		})
		require.NoError(t, err)

		msg, broken := keeper.AirdropSupplyInvariant(*tk.ClaimKeeper)(ctx)
		require.True(t, broken, msg)
	})

	t.Run("should break with airdrop supply not equal to claimable amounts", func(t *testing.T) {
		ctx, tk, _ := testkeeper.NewTestSetup(t)

		err := tk.ClaimKeeper.InitializeAirdropSupply(ctx, sdk.NewCoin("test", sdkmath.NewInt(10)))
		require.NoError(t, err)

		address := sample.Address(r)
		err = tk.ClaimKeeper.ClaimRecord.Set(ctx, address, types.ClaimRecord{
			Address:           address,
			Claimable:         sdkmath.NewInt(9),
			CompletedMissions: nil,
		})
		require.NoError(t, err)

		msg, broken := keeper.AirdropSupplyInvariant(*tk.ClaimKeeper)(ctx)
		require.True(t, broken, msg)
	})
}
