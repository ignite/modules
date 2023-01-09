package keeper_test

import (
	"testing"
	"time"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	errorsignite "github.com/ignite/modules/pkg/errors"
	tc "github.com/ignite/modules/testutil/constructor"
	testkeeper "github.com/ignite/modules/testutil/keeper"
	"github.com/ignite/modules/testutil/sample"
	"github.com/ignite/modules/x/claim/types"
)

func TestMsgClaim(t *testing.T) {
	sdkCtx, tk, ts := testkeeper.NewTestSetup(t)
	ctx := sdk.WrapSDKContext(sdkCtx)

	// prepare addresses
	var addr []string
	for i := 0; i < 20; i++ {
		addr = append(addr, sample.Address(r))
	}

	type inputState struct {
		noAirdropSupply bool
		noMission       bool
		noInitialClaim  bool
		noClaimRecord   bool
		airdropSupply   sdk.Coin
		mission         types.Mission
		initialClaim    types.InitialClaim
		claimRecord     types.ClaimRecord
		params          types.Params
		airdropStart    time.Time
		blockTime       time.Time
	}
	tests := []struct {
		name            string
		inputState      inputState
		msg             types.MsgClaim
		expectedBalance sdk.Coin
		err             error
	}{
		{
			name: "should fail if no airdrop supply",
			inputState: inputState{
				noInitialClaim:  true,
				noAirdropSupply: true,
				mission: types.Mission{
					MissionID: 1,
					Weight:    sdk.OneDec(),
				},
				claimRecord: types.ClaimRecord{
					Address:           addr[0],
					Claimable:         sdkmath.OneInt(),
					CompletedMissions: []uint64{1},
				},
				params: types.DefaultParams(),
			},
			msg: types.MsgClaim{
				Claimer:   addr[0],
				MissionID: 1,
			},
			err: types.ErrAirdropSupplyNotFound,
		},
		{
			name: "should fail if no claim record",
			inputState: inputState{
				noInitialClaim: true,
				noClaimRecord:  true,
				airdropSupply:  sample.Coin(r),
				mission:        sample.Mission(r),
				params:         types.DefaultParams(),
			},
			msg: types.MsgClaim{
				Claimer:   sample.Address(r),
				MissionID: 1,
			},
			err: types.ErrClaimRecordNotFound,
		},
		{
			name: "should fail if no mission",
			inputState: inputState{
				noMission:     true,
				airdropSupply: sample.Coin(r),
				claimRecord: types.ClaimRecord{
					Address:           addr[1],
					Claimable:         sdkmath.OneInt(),
					CompletedMissions: []uint64{1},
				},
				params: types.DefaultParams(),
			},
			msg: types.MsgClaim{
				Claimer:   addr[1],
				MissionID: 1,
			},
			err: types.ErrMissionNotFound,
		},
		{
			name: "should fail if already claimed",
			inputState: inputState{
				noInitialClaim: true,
				airdropSupply:  sample.Coin(r),
				mission: types.Mission{
					MissionID: 1,
					Weight:    sdk.OneDec(),
				},
				claimRecord: types.ClaimRecord{
					Address:           addr[2],
					Claimable:         sdkmath.OneInt(),
					CompletedMissions: []uint64{1},
					ClaimedMissions:   []uint64{1},
				},
				params: types.DefaultParams(),
			},
			msg: types.MsgClaim{
				Claimer:   addr[2],
				MissionID: 1,
			},
			err: types.ErrMissionAlreadyClaimed,
		},
		{
			name: "should fail if mission not completed",
			inputState: inputState{
				noInitialClaim: true,
				airdropSupply:  sample.Coin(r),
				mission: types.Mission{
					MissionID: 1,
					Weight:    sdk.OneDec(),
				},
				claimRecord: types.ClaimRecord{
					Address:   addr[3],
					Claimable: sdkmath.OneInt(),
				},
				params: types.DefaultParams(),
			},
			msg: types.MsgClaim{
				Claimer:   addr[3],
				MissionID: 1,
			},
			err: types.ErrMissionNotCompleted,
		},
		{
			name: "should fail with critical if claimable amount is greater than module supply",
			inputState: inputState{
				noInitialClaim: true,
				airdropSupply:  tc.Coin(t, "1000foo"),
				mission: types.Mission{
					MissionID: 1,
					Weight:    sdk.OneDec(),
				},
				claimRecord: types.ClaimRecord{
					Address:           addr[4],
					Claimable:         sdkmath.NewIntFromUint64(10000),
					CompletedMissions: []uint64{1},
				},
				params: types.DefaultParams(),
			},
			msg: types.MsgClaim{
				Claimer:   addr[4],
				MissionID: 1,
			},
			err: errorsignite.ErrCritical,
		},
		{
			name: "should fail with critical if claimer address is not bech32",
			inputState: inputState{
				noInitialClaim: true,
				airdropSupply:  tc.Coin(t, "1000foo"),
				mission: types.Mission{
					MissionID: 1,
					Weight:    sdk.OneDec(),
				},
				claimRecord: types.ClaimRecord{
					Address:           "invalid",
					Claimable:         sdkmath.OneInt(),
					CompletedMissions: []uint64{1},
				},
				params: types.DefaultParams(),
			},
			msg: types.MsgClaim{
				Claimer:   "invalid",
				MissionID: 1,
			},
			err: errorsignite.ErrCritical,
		},
		{
			name: "should fail if airdrop start not reached",
			inputState: inputState{
				noInitialClaim: true,
				airdropSupply:  sample.Coin(r),
				mission:        sample.Mission(r),
				claimRecord: types.ClaimRecord{
					Address:   addr[5],
					Claimable: sdkmath.NewIntFromUint64(1000),
				},
				blockTime: time.Unix(0, 0),
				params: types.NewParams(
					types.NewDisabledDecay(),
					time.Unix(20000, 0),
				),
			},
			msg: types.MsgClaim{
				Claimer:   addr[5],
				MissionID: 1,
			},
			err: types.ErrAirdropStartNotReached,
		},
		{
			name: "should allow distributing full airdrop to one account, one mission",
			inputState: inputState{
				noInitialClaim: true,
				airdropSupply:  tc.Coin(t, "1000foo"),
				mission: types.Mission{
					MissionID: 1,
					Weight:    sdk.OneDec(),
				},
				claimRecord: types.ClaimRecord{
					Address:           addr[6],
					Claimable:         sdkmath.NewIntFromUint64(1000),
					CompletedMissions: []uint64{1},
				},
				params: types.DefaultParams(),
			},
			msg: types.MsgClaim{
				Claimer:   addr[6],
				MissionID: 1,
			},
			expectedBalance: tc.Coin(t, "1000foo"),
		},
		{
			name: "should prevent distributing fund for mission with 0 weight",
			inputState: inputState{
				noInitialClaim: true,
				airdropSupply:  tc.Coin(t, "1000foo"),
				mission: types.Mission{
					MissionID: 1,
					Weight:    sdk.ZeroDec(),
				},
				claimRecord: types.ClaimRecord{
					Address:           addr[7],
					Claimable:         sdkmath.NewIntFromUint64(1000),
					CompletedMissions: []uint64{1},
				},
				params: types.DefaultParams(),
			},
			msg: types.MsgClaim{
				Claimer:   addr[7],
				MissionID: 1,
			},
			err: types.ErrNoClaimable,
		},
		{
			name: "should allow distributing half for mission with 0.5 weight",
			inputState: inputState{
				noInitialClaim: true,
				airdropSupply:  tc.Coin(t, "1000foo"),
				mission: types.Mission{
					MissionID: 1,
					Weight:    tc.Dec(t, "0.5"),
				},
				claimRecord: types.ClaimRecord{
					Address:           addr[8],
					Claimable:         sdkmath.NewIntFromUint64(500),
					CompletedMissions: []uint64{1},
				},
				params: types.DefaultParams(),
			},
			msg: types.MsgClaim{
				Claimer:   addr[8],
				MissionID: 1,
			},
			expectedBalance: tc.Coin(t, "250foo"),
		},
		{
			name: "should allow distributing half for mission with 0.5 weight and truncate decimal",
			inputState: inputState{
				noInitialClaim: true,
				airdropSupply:  tc.Coin(t, "1000foo"),
				mission: types.Mission{
					MissionID: 1,
					Weight:    tc.Dec(t, "0.5"),
				},
				claimRecord: types.ClaimRecord{
					Address:           addr[9],
					Claimable:         sdkmath.NewIntFromUint64(201),
					CompletedMissions: []uint64{1},
				},
				params: types.DefaultParams(),
			},
			msg: types.MsgClaim{
				Claimer:   addr[9],
				MissionID: 1,
			},
			expectedBalance: tc.Coin(t, "100foo"),
		},
		{
			name: "should prevent distributing fund for empty claim record",
			inputState: inputState{
				noInitialClaim: true,
				airdropSupply:  tc.Coin(t, "1000foo"),
				mission: types.Mission{
					MissionID: 1,
					Weight:    tc.Dec(t, "0.5"),
				},
				claimRecord: types.ClaimRecord{
					Address:           addr[10],
					Claimable:         sdkmath.ZeroInt(),
					CompletedMissions: []uint64{1},
				},
				params: types.DefaultParams(),
			},
			msg: types.MsgClaim{
				Claimer:   addr[10],
				MissionID: 1,
			},
			err: types.ErrNoClaimable,
		},
		{
			name: "should allow distributing airdrop with other already completed missions",
			inputState: inputState{
				noInitialClaim: true,
				airdropSupply:  tc.Coin(t, "10000bar"),
				mission: types.Mission{
					MissionID: 3,
					Weight:    tc.Dec(t, "0.3"),
				},
				claimRecord: types.ClaimRecord{
					Address:           addr[11],
					Claimable:         sdkmath.NewIntFromUint64(10000),
					CompletedMissions: []uint64{0, 1, 3, 2, 4, 5, 6},
				},
				params: types.DefaultParams(),
			},
			msg: types.MsgClaim{
				Claimer:   addr[11],
				MissionID: 3,
			},
			expectedBalance: tc.Coin(t, "3000bar"),
		},
		{
			name: "should allow applying decay factor if enabled",
			inputState: inputState{
				noInitialClaim: true,
				airdropSupply:  tc.Coin(t, "1000foo"),
				mission: types.Mission{
					MissionID: 1,
					Weight:    tc.Dec(t, "0.5"),
				},
				claimRecord: types.ClaimRecord{
					Address:           addr[12],
					Claimable:         sdkmath.NewIntFromUint64(1000),
					CompletedMissions: []uint64{1},
				},
				params: types.NewParams(types.NewEnabledDecay(
					time.Unix(1000, 0),
					time.Unix(2000, 0),
				), time.Unix(0, 0)),
				blockTime: time.Unix(1500, 0),
			},
			msg: types.MsgClaim{
				Claimer:   addr[12],
				MissionID: 1,
			},
			expectedBalance: tc.Coin(t, "250foo"),
		},
		{
			name: "should allow distributing all funds if decay factor if enabled and decay not started",
			inputState: inputState{
				noInitialClaim: true,
				airdropSupply:  tc.Coin(t, "1000foo"),
				mission: types.Mission{
					MissionID: 1,
					Weight:    tc.Dec(t, "0.5"),
				},
				claimRecord: types.ClaimRecord{
					Address:           addr[13],
					Claimable:         sdkmath.NewIntFromUint64(1000),
					CompletedMissions: []uint64{1},
				},
				params: types.NewParams(types.NewEnabledDecay(
					time.Unix(1000, 0),
					time.Unix(2000, 0),
				), time.Unix(0, 0)),
				blockTime: time.Unix(999, 0),
			},
			msg: types.MsgClaim{
				Claimer:   addr[13],
				MissionID: 1,
			},
			expectedBalance: tc.Coin(t, "500foo"),
		},
		{
			name: "should prevent distributing funds if decay ended",
			inputState: inputState{
				noInitialClaim: true,
				airdropSupply:  tc.Coin(t, "1000foo"),
				mission: types.Mission{
					MissionID: 1,
					Weight:    tc.Dec(t, "0.5"),
				},
				claimRecord: types.ClaimRecord{
					Address:           addr[14],
					Claimable:         sdkmath.NewIntFromUint64(1000),
					CompletedMissions: []uint64{1},
				},
				params: types.NewParams(types.NewEnabledDecay(
					time.Unix(1000, 0),
					time.Unix(2000, 0),
				), time.Unix(0, 0)),
				blockTime: time.Unix(2001, 0),
			},
			msg: types.MsgClaim{
				Claimer:   addr[14],
				MissionID: 1,
			},
			err: types.ErrNoClaimable,
		},
		{
			name: "should allow to claim initial for an existing mission and claim record",
			inputState: inputState{
				airdropSupply: tc.Coin(t, "100000foo"),
				initialClaim: types.InitialClaim{
					Enabled:   true,
					MissionID: 1,
				},
				mission: types.Mission{
					MissionID: 1,
					Weight:    sdk.OneDec(),
				},
				claimRecord: types.ClaimRecord{
					Address:   addr[15],
					Claimable: sdkmath.NewIntFromUint64(100),
				},
			},
			msg: types.MsgClaim{
				Claimer:   addr[15],
				MissionID: 1,
			},
			expectedBalance: tc.Coin(t, "100foo"),
		},
		{
			name: "should prevent claiming initial if initial claim not enabled",
			inputState: inputState{
				airdropSupply: tc.Coin(t, "100000foo"),
				initialClaim: types.InitialClaim{
					Enabled:   false,
					MissionID: 1,
				},
				mission: types.Mission{
					MissionID: 1,
					Weight:    sdk.OneDec(),
				},
				claimRecord: types.ClaimRecord{
					Address:   addr[16],
					Claimable: sdkmath.NewIntFromUint64(100),
				},
			},
			msg: types.MsgClaim{
				Claimer:   addr[16],
				MissionID: 1,
			},
			err: types.ErrInitialClaimNotEnabled,
		},
		{
			name: "should prevent claiming initial already claimed",
			inputState: inputState{
				airdropSupply: tc.Coin(t, "100000foo"),
				initialClaim: types.InitialClaim{
					Enabled:   true,
					MissionID: 1,
				},
				mission: types.Mission{
					MissionID: 1,
					Weight:    sdk.OneDec(),
				},
				claimRecord: types.ClaimRecord{
					Address:           addr[17],
					Claimable:         sdkmath.NewIntFromUint64(100),
					CompletedMissions: []uint64{1},
					ClaimedMissions:   []uint64{1},
				},
			},
			msg: types.MsgClaim{
				Claimer:   addr[17],
				MissionID: 1,
			},
			err: types.ErrMissionAlreadyClaimed,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// initialize input state
			require.NoError(t, tt.inputState.params.Validate())
			tk.ClaimKeeper.SetParams(sdkCtx, tt.inputState.params)
			if !tt.inputState.noAirdropSupply {
				err := tk.ClaimKeeper.InitializeAirdropSupply(sdkCtx, tt.inputState.airdropSupply)
				require.NoError(t, err)
			}
			if !tt.inputState.noInitialClaim {
				tk.ClaimKeeper.SetInitialClaim(sdkCtx, tt.inputState.initialClaim)
			}
			if !tt.inputState.noMission {
				tk.ClaimKeeper.SetMission(sdkCtx, tt.inputState.mission)
			}
			if !tt.inputState.noClaimRecord {
				tk.ClaimKeeper.SetClaimRecord(sdkCtx, tt.inputState.claimRecord)
			}
			if !tt.inputState.blockTime.IsZero() {
				ctx = sdkCtx.WithBlockTime(tt.inputState.blockTime)
			}

			res, err := ts.ClaimSrv.Claim(ctx, &tt.msg)
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
			} else {
				require.NoError(t, err)

				// funds are distributed to the user
				sdkAddr, err := sdk.AccAddressFromBech32(tt.msg.Claimer)
				require.NoError(t, err)

				require.Equal(t, tt.expectedBalance.Amount, res.Claimed)

				balance := tk.BankKeeper.GetBalance(sdkCtx, sdkAddr, tt.inputState.airdropSupply.Denom)
				require.True(t, balance.IsEqual(tt.expectedBalance),
					"expected balance after mission complete: %s, actual balance: %s",
					tt.expectedBalance.String(),
					balance.String(),
				)

				// completed mission is added in claim record
				claimRecord, found := tk.ClaimKeeper.GetClaimRecord(sdkCtx, tt.msg.Claimer)
				require.True(t, found)
				require.True(t, claimRecord.IsMissionCompleted(tt.msg.MissionID))

				// airdrop supply is updated with distributed balance
				airdropSupply, found := tk.ClaimKeeper.GetAirdropSupply(sdkCtx)
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
				tk.ClaimKeeper.RemoveAirdropSupply(sdkCtx)
			}
			if !tt.inputState.noMission {
				tk.ClaimKeeper.RemoveMission(sdkCtx, tt.inputState.mission.MissionID)
			}
			if !tt.inputState.noClaimRecord {
				tk.ClaimKeeper.RemoveClaimRecord(sdkCtx, tt.inputState.claimRecord.Address)
			}
			if !tt.inputState.noInitialClaim {
				tk.ClaimKeeper.RemoveInitialClaim(sdkCtx)
			}
		})
	}
}
