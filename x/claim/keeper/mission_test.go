package keeper_test

import (
	"testing"
	"time"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	errorsignite "github.com/ignite/modules/pkg/errors"
	"github.com/ignite/modules/testutil/nullify"
	"github.com/ignite/modules/testutil/sample"
	"github.com/ignite/modules/x/claim/keeper"
	"github.com/ignite/modules/x/claim/types"
)

func createNMission(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Mission {
	items := make([]types.Mission, n)
	for i := range items {
		items[i].MissionID = uint64(i)
		items[i].Weight = sdkmath.LegacyNewDec(r.Int63())
		keeper.SetMission(ctx, items[i])
	}
	return items
}

func TestMissionGet(t *testing.T) {
	ctx, tk := createClaimKeeper(t)

	t.Run("should allow get", func(t *testing.T) {
		items := createNMission(tk, ctx, 10)
		for _, item := range items {
			got, found := tk.GetMission(ctx, item.MissionID)
			require.True(t, found)
			require.Equal(t,
				nullify.Fill(&item),
				nullify.Fill(&got),
			)
		}
	})
}

func TestMissionGetAll(t *testing.T) {
	ctx, tk := createClaimKeeper(t)

	t.Run("should allow get all", func(t *testing.T) {
		items := createNMission(tk, ctx, 10)
		require.ElementsMatch(t,
			nullify.Fill(items),
			nullify.Fill(tk.GetAllMission(ctx)),
		)
	})
}

func TestMissionRemove(t *testing.T) {
	ctx, tk := createClaimKeeper(t)

	t.Run("should allow remove", func(t *testing.T) {
		items := createNMission(tk, ctx, 10)
		for _, item := range items {
			tk.RemoveMission(ctx, item.MissionID)
			_, found := tk.GetMission(ctx, item.MissionID)
			require.False(t, found)
		}
	})
}

func TestKeeper_ClaimMission(t *testing.T) {
	ctx, tk := createClaimKeeper(t)

	// prepare addresses
	addr := make([]string, 20)
	for i := 0; i < len(addr); i++ {
		addr[i] = sample.AccAddress()
	}

	type inputState struct {
		noAirdropSupply bool
		noMission       bool
		noClaimRecord   bool
		airdropSupply   sdk.Coin
		mission         types.Mission
		claimRecord     types.ClaimRecord
		params          types.Params
		blockTime       time.Time
	}
	tests := []struct {
		name            string
		inputState      inputState
		missionID       uint64
		address         string
		expectedBalance sdk.Coin
		err             error
	}{
		{
			name: "should fail if no airdrop supply",
			inputState: inputState{
				noAirdropSupply: true,
				claimRecord: types.ClaimRecord{
					Address:           addr[0],
					Claimable:         sdkmath.NewInt(r.Int63n(100000)),
					CompletedMissions: []uint64{1, 2, 3},
				},
				mission: types.Mission{
					MissionID:   1,
					Description: "dummy mission",
					Weight:      sdkmath.LegacyNewDec(r.Int63n(1_000_000)).Quo(sdkmath.LegacyNewDec(1_000_000)),
				},
				params: types.DefaultParams(),
			},
			missionID: 1,
			address:   sample.AccAddress(),
			err:       types.ErrAirdropSupplyNotFound,
		},
		{
			name: "should fail if no mission",
			inputState: inputState{
				noMission:     true,
				airdropSupply: sdk.NewCoin("foo", sdkmath.NewInt(1000)),
				claimRecord: types.ClaimRecord{
					Address:           addr[0],
					Claimable:         sdkmath.OneInt(),
					CompletedMissions: []uint64{1},
				},
				params: types.DefaultParams(),
			},
			missionID: 1,
			address:   addr[0],
			err:       types.ErrMissionNotFound,
		},
		{
			name: "should fail if already claimed",
			inputState: inputState{
				airdropSupply: sdk.NewCoin("foo", sdkmath.NewInt(1000)),
				mission: types.Mission{
					MissionID: 1,
					Weight:    sdkmath.LegacyOneDec(),
				},
				claimRecord: types.ClaimRecord{
					Address:           addr[1],
					Claimable:         sdkmath.OneInt(),
					CompletedMissions: []uint64{1},
					ClaimedMissions:   []uint64{1},
				},
				params: types.DefaultParams(),
			},
			missionID: 1,
			address:   addr[1],
			err:       types.ErrMissionAlreadyClaimed,
		},
		{
			name: "should fail if mission not completed",
			inputState: inputState{
				airdropSupply: sdk.NewCoin("foo", sdkmath.NewInt(1000)),
				mission: types.Mission{
					MissionID: 1,
					Weight:    sdkmath.LegacyOneDec(),
				},
				claimRecord: types.ClaimRecord{
					Address:   addr[1],
					Claimable: sdkmath.OneInt(),
				},
				params: types.DefaultParams(),
			},
			missionID: 1,
			address:   addr[1],
			err:       types.ErrMissionNotCompleted,
		},
		{
			name: "should fail with critical if claimable amount is greater than module supply",
			inputState: inputState{
				airdropSupply: sdk.NewCoin("foo", sdkmath.NewInt(1000)),
				mission: types.Mission{
					MissionID: 1,
					Weight:    sdkmath.LegacyOneDec(),
				},
				claimRecord: types.ClaimRecord{
					Address:           addr[2],
					Claimable:         sdkmath.NewIntFromUint64(10000),
					CompletedMissions: []uint64{1},
				},
				params: types.DefaultParams(),
			},
			missionID: 1,
			address:   addr[2],
			err:       errorsignite.ErrCritical,
		},
		{
			name: "should fail with critical if claimer address is not bech32",
			inputState: inputState{
				airdropSupply: sdk.NewCoin("foo", sdkmath.NewInt(1000)),
				mission: types.Mission{
					MissionID: 1,
					Weight:    sdkmath.LegacyOneDec(),
				},
				claimRecord: types.ClaimRecord{
					Address:           "invalid",
					Claimable:         sdkmath.OneInt(),
					CompletedMissions: []uint64{1},
				},
				params: types.DefaultParams(),
			},
			missionID: 1,
			address:   "invalid",
			err:       errorsignite.ErrCritical,
		},
		{
			name: "should allow distributing full airdrop to one account, one mission",
			inputState: inputState{
				airdropSupply: sdk.NewCoin("foo", sdkmath.NewInt(1000)),
				mission: types.Mission{
					MissionID: 1,
					Weight:    sdkmath.LegacyOneDec(),
				},
				claimRecord: types.ClaimRecord{
					Address:           addr[3],
					Claimable:         sdkmath.NewIntFromUint64(1000),
					CompletedMissions: []uint64{1},
				},
				params: types.DefaultParams(),
			},
			missionID:       1,
			address:         addr[3],
			expectedBalance: sdk.NewCoin("foo", sdkmath.NewInt(1000)),
		},
		{
			name: "should prevent distributing fund for mission with 0 weight",
			inputState: inputState{
				airdropSupply: sdk.NewCoin("foo", sdkmath.NewInt(1000)),
				mission: types.Mission{
					MissionID: 1,
					Weight:    sdkmath.LegacyZeroDec(),
				},
				claimRecord: types.ClaimRecord{
					Address:           addr[4],
					Claimable:         sdkmath.NewIntFromUint64(1000),
					CompletedMissions: []uint64{1},
				},
				params: types.DefaultParams(),
			},
			missionID: 1,
			address:   addr[4],
			err:       types.ErrNoClaimable,
		},
		{
			name: "should allow distributing half for mission with 0.5 weight",
			inputState: inputState{
				airdropSupply: sdk.NewCoin("foo", sdkmath.NewInt(1000)),
				mission: types.Mission{
					MissionID: 1,
					Weight:    sdkmath.LegacyMustNewDecFromStr("0.5"),
				},
				claimRecord: types.ClaimRecord{
					Address:           addr[5],
					Claimable:         sdkmath.NewIntFromUint64(500),
					CompletedMissions: []uint64{1},
				},
				params: types.DefaultParams(),
			},
			missionID:       1,
			address:         addr[5],
			expectedBalance: sdk.NewCoin("foo", sdkmath.NewInt(250)),
		},
		{
			name: "should allow distributing half for mission with 0.5 weight and truncate decimal",
			inputState: inputState{
				airdropSupply: sdk.NewCoin("foo", sdkmath.NewInt(1000)),
				mission: types.Mission{
					MissionID: 1,
					Weight:    sdkmath.LegacyMustNewDecFromStr("0.5"),
				},
				claimRecord: types.ClaimRecord{
					Address:           addr[6],
					Claimable:         sdkmath.NewIntFromUint64(201),
					CompletedMissions: []uint64{1},
				},
				params: types.DefaultParams(),
			},
			missionID:       1,
			address:         addr[6],
			expectedBalance: sdk.NewCoin("foo", sdkmath.NewInt(100)),
		},
		{
			name: "should prevent distributing fund for empty claim record",
			inputState: inputState{
				airdropSupply: sdk.NewCoin("foo", sdkmath.NewInt(1000)),
				mission: types.Mission{
					MissionID: 1,
					Weight:    sdkmath.LegacyMustNewDecFromStr("0.5"),
				},
				claimRecord: types.ClaimRecord{
					Address:           addr[7],
					Claimable:         sdkmath.ZeroInt(),
					CompletedMissions: []uint64{1},
				},
				params: types.DefaultParams(),
			},
			missionID: 1,
			address:   addr[7],
			err:       types.ErrNoClaimable,
		},
		{
			name: "should allow distributing airdrop with other already completed missions",
			inputState: inputState{
				airdropSupply: sdk.NewCoin("bar", sdkmath.NewInt(1000)),
				mission: types.Mission{
					MissionID: 3,
					Weight:    sdkmath.LegacyMustNewDecFromStr("0.3"),
				},
				claimRecord: types.ClaimRecord{
					Address:           addr[8],
					Claimable:         sdkmath.NewIntFromUint64(10000),
					CompletedMissions: []uint64{0, 1, 3, 2, 4, 5, 6},
				},
				params: types.DefaultParams(),
			},
			missionID:       3,
			address:         addr[8],
			expectedBalance: sdk.NewCoin("bar", sdkmath.NewInt(3000)),
		},
		{
			name: "should allow applying decay factor if enabled",
			inputState: inputState{
				airdropSupply: sdk.NewCoin("foo", sdkmath.NewInt(1000)),
				mission: types.Mission{
					MissionID: 1,
					Weight:    sdkmath.LegacyMustNewDecFromStr("0.5"),
				},
				claimRecord: types.ClaimRecord{
					Address:           addr[9],
					Claimable:         sdkmath.NewIntFromUint64(1000),
					CompletedMissions: []uint64{1},
				},
				params: types.NewParams(types.NewEnabledDecay(
					time.Unix(1000, 0),
					time.Unix(2000, 0),
				), time.Unix(0, 0)),
				blockTime: time.Unix(1500, 0),
			},
			missionID:       1,
			address:         addr[9],
			expectedBalance: sdk.NewCoin("foo", sdkmath.NewInt(250)),
		},
		{
			name: "should allow distributing all funds if decay factor if enabled and decay not started",
			inputState: inputState{
				airdropSupply: sdk.NewCoin("foo", sdkmath.NewInt(1000)),
				mission: types.Mission{
					MissionID: 1,
					Weight:    sdkmath.LegacyMustNewDecFromStr("0.5"),
				},
				claimRecord: types.ClaimRecord{
					Address:           addr[10],
					Claimable:         sdkmath.NewIntFromUint64(1000),
					CompletedMissions: []uint64{1},
				},
				params: types.NewParams(types.NewEnabledDecay(
					time.Unix(1000, 0),
					time.Unix(2000, 0),
				), time.Unix(0, 0)),
				blockTime: time.Unix(999, 0),
			},
			missionID:       1,
			address:         addr[10],
			expectedBalance: sdk.NewCoin("foo", sdkmath.NewInt(500)),
		},
		{
			name: "should prevent distributing funds if decay ended",
			inputState: inputState{
				airdropSupply: sdk.NewCoin("foo", sdkmath.NewInt(1000)),
				mission: types.Mission{
					MissionID: 1,
					Weight:    sdkmath.LegacyMustNewDecFromStr("0.5"),
				},
				claimRecord: types.ClaimRecord{
					Address:           addr[11],
					Claimable:         sdkmath.NewIntFromUint64(1000),
					CompletedMissions: []uint64{1},
				},
				params: types.NewParams(types.NewEnabledDecay(
					time.Unix(1000, 0),
					time.Unix(2000, 0),
				), time.Unix(0, 0)),
				blockTime: time.Unix(2001, 0),
			},
			missionID: 1,
			address:   addr[11],
			err:       types.ErrNoClaimable,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// initialize input state
			require.NoError(t, tt.inputState.params.Validate())
			tk.SetParams(ctx, tt.inputState.params)
			if !tt.inputState.noAirdropSupply {
				err := tk.InitializeAirdropSupply(ctx, tt.inputState.airdropSupply)
				require.NoError(t, err)
			}
			if !tt.inputState.noMission {
				tk.SetMission(ctx, tt.inputState.mission)
			}
			if !tt.inputState.noClaimRecord {
				tk.SetClaimRecord(ctx, tt.inputState.claimRecord)
			}
			if !tt.inputState.blockTime.IsZero() {
				ctx = ctx.WithBlockTime(tt.inputState.blockTime)
			}

			claimed, err := tk.ClaimMission(ctx, tt.inputState.claimRecord, tt.missionID)
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
			} else {
				require.NoError(t, err)

				// funds are distributed to the user
				sdkAddr, err := sdk.AccAddressFromBech32(tt.address)
				require.NoError(t, err)

				require.Equal(t, tt.expectedBalance.Amount, claimed)

				balance := tk.BankKeeper.GetBalance(ctx, sdkAddr, tt.inputState.airdropSupply.Denom)
				require.True(t, balance.IsEqual(tt.expectedBalance),
					"expected balance after mission complete: %s, actual balance: %s",
					tt.expectedBalance.String(),
					balance.String(),
				)

				// completed mission is added in claim record
				claimRecord, found := tk.GetClaimRecord(ctx, tt.address)
				require.True(t, found)
				require.True(t, claimRecord.IsMissionCompleted(tt.missionID))

				// airdrop supply is updated with distributed balance
				airdropSupply, found := tk.GetAirdropSupply(ctx)
				require.True(t, found)
				expectedAidropSupply := tt.inputState.airdropSupply.Sub(tt.expectedBalance)

				require.True(t, airdropSupply.IsEqual(expectedAidropSupply),
					"expected airdrop supply after mission complete: %s, actual supply: %s",
					expectedAidropSupply,
					airdropSupply,
				)
			}

			// clear input state
			if !tt.inputState.noAirdropSupply {
				tk.RemoveAirdropSupply(ctx)
			}
			if !tt.inputState.noMission {
				tk.RemoveMission(ctx, tt.inputState.mission.MissionID)
			}
			if !tt.inputState.noClaimRecord {
				tk.RemoveClaimRecord(ctx, tt.inputState.claimRecord.Address)
			}
		})
	}
}

func TestKeeper_CompleteMission(t *testing.T) {
	ctx, tk := createClaimKeeper(t)

	addr := make([]string, 7)
	for i := 0; i < len(addr); i++ {
		addr[i] = sample.AccAddress()
	}

	type inputState struct {
		airdropSupply sdk.Coin
		mission       types.Mission
		claimRecord   types.ClaimRecord
		params        types.Params
		blockTime     time.Time
	}
	tests := []struct {
		name            string
		inputState      inputState
		missionID       uint64
		address         string
		isClaimed       bool
		expectedClaimed sdkmath.Int
		err             error
	}{
		{
			name: "should fail if mission id not found",
			inputState: inputState{
				mission: types.Mission{
					MissionID: 1,
					Weight:    sdkmath.LegacyMustNewDecFromStr("0.5"),
				},
				claimRecord: types.ClaimRecord{
					Address:           addr[0],
					Claimable:         sdkmath.NewIntFromUint64(1000),
					CompletedMissions: []uint64{1},
					ClaimedMissions:   []uint64{1},
				},
				params: types.NewParams(types.NewEnabledDecay(
					time.Unix(1000, 0),
					time.Unix(2000, 0),
				), time.Unix(2001, 0)),
				blockTime: time.Unix(0, 0),
			},
			missionID: 10,
			address:   addr[0],
			err:       types.ErrMissionNotFound,
		},
		{
			name: "should fail if claim record id not found",
			inputState: inputState{
				mission: types.Mission{
					MissionID: 1,
					Weight:    sdkmath.LegacyMustNewDecFromStr("0.5"),
				},
				claimRecord: types.ClaimRecord{
					Address:           addr[1],
					Claimable:         sdkmath.NewIntFromUint64(1000),
					CompletedMissions: []uint64{1},
					ClaimedMissions:   []uint64{1},
				},
				params: types.NewParams(types.NewEnabledDecay(
					time.Unix(1000, 0),
					time.Unix(2000, 0),
				), time.Unix(2001, 0)),
				blockTime: time.Unix(0, 0),
			},
			missionID: 1,
			address:   sample.AccAddress(),
			err:       types.ErrClaimRecordNotFound,
		},
		{
			name: "should fail if mission already completed",
			inputState: inputState{
				mission: types.Mission{
					MissionID: 1,
					Weight:    sdkmath.LegacyMustNewDecFromStr("0.5"),
				},
				claimRecord: types.ClaimRecord{
					Address:           addr[2],
					Claimable:         sdkmath.NewIntFromUint64(1000),
					CompletedMissions: []uint64{1},
				},
				params: types.NewParams(types.NewEnabledDecay(
					time.Unix(1000, 0),
					time.Unix(2000, 0),
				), time.Unix(2001, 0)),
				blockTime: time.Unix(0, 0),
			},
			missionID: 1,
			address:   addr[2],
			err:       types.ErrMissionCompleted,
		},
		{
			name: "should success",
			inputState: inputState{
				mission: types.Mission{
					MissionID: 1,
					Weight:    sdkmath.LegacyMustNewDecFromStr("0.5"),
				},
				claimRecord: types.ClaimRecord{
					Address:   addr[4],
					Claimable: sdkmath.NewIntFromUint64(1000),
				},
				params: types.NewParams(types.NewEnabledDecay(
					time.Unix(1000, 0),
					time.Unix(2000, 0),
				), time.Unix(2001, 0)),
				blockTime: time.Unix(0, 0),
			},
			missionID: 1,
			address:   addr[4],
		},
		{
			name: "should success and fail to claim balance",
			inputState: inputState{
				mission: types.Mission{
					MissionID: 1,
					Weight:    sdkmath.LegacyMustNewDecFromStr("0.5"),
				},
				claimRecord: types.ClaimRecord{
					Address:   addr[5],
					Claimable: sdkmath.NewIntFromUint64(1000),
				},
				params: types.NewParams(types.NewEnabledDecay(
					time.Unix(1000, 0),
					time.Unix(2000, 0),
				), time.Unix(0, 0)),
				blockTime: time.Unix(50000, 0),
			},
			missionID: 1,
			address:   addr[5],
			err:       types.ErrAirdropSupplyNotFound,
		},
		{
			name: "should complete mission and allow to claim",
			inputState: inputState{
				airdropSupply: sdk.NewCoin("foo", sdkmath.NewInt(1000)),
				mission: types.Mission{
					MissionID: 1,
					Weight:    sdkmath.LegacyOneDec(),
				},
				claimRecord: types.ClaimRecord{
					Address:   addr[6],
					Claimable: sdkmath.NewIntFromUint64(1000),
				},
				params: types.DefaultParams(),
			},
			missionID:       1,
			address:         addr[6],
			expectedClaimed: sdkmath.NewInt(1000),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.NoError(t, tt.inputState.params.Validate())
			if !tt.inputState.airdropSupply.IsNil() && !tt.inputState.airdropSupply.IsZero() {
				err := tk.InitializeAirdropSupply(ctx, tt.inputState.airdropSupply)
				require.NoError(t, err)
			}
			tk.SetParams(ctx, tt.inputState.params)
			tk.SetMission(ctx, tt.inputState.mission)
			tk.SetClaimRecord(ctx, tt.inputState.claimRecord)
			if !tt.inputState.blockTime.IsZero() {
				ctx = ctx.WithBlockTime(tt.inputState.blockTime)
			}

			claimed, err := tk.CompleteMission(ctx, tt.missionID, tt.address)
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
			require.Equal(t, tt.expectedClaimed, claimed)

			claimRecord, found := tk.GetClaimRecord(ctx, tt.address)
			require.True(t, found)
			require.True(t, claimRecord.IsMissionCompleted(tt.missionID))
		})
	}
}
