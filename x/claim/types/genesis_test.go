package types_test

import (
	"testing"
	"time"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/ignite/modules/testutil/sample"
	"github.com/ignite/modules/x/claim/types"
)

func TestGenesisState_Validate(t *testing.T) {
	fiftyPercent, err := sdkmath.LegacyNewDecFromStr("0.5")
	require.NoError(t, err)

	claimAmts := []sdkmath.Int{
		sdkmath.NewIntFromUint64(10),
		sdkmath.NewIntFromUint64(5020),
	}

	for _, tt := range []struct {
		name     string
		genState *types.GenesisState
		valid    bool
	}{
		{
			name:     "should validate default",
			genState: types.DefaultGenesis(),
			valid:    true,
		},
		{
			name: "should validate airdrop supply sum of claim amounts",
			genState: &types.GenesisState{
				Params: types.DefaultParams(),
				ClaimRecords: []types.ClaimRecord{
					{
						Address:   sample.AccAddress(),
						Claimable: claimAmts[0],
					},
					{
						Address:   sample.AccAddress(),
						Claimable: claimAmts[1],
					},
				},
				Missions: []types.Mission{
					{
						MissionID: 0,
						Weight:    fiftyPercent,
					},
					{
						MissionID: 1,
						Weight:    fiftyPercent,
					},
				},
				AirdropSupply: sdk.NewCoin("foo", claimAmts[0].Add(claimAmts[1])),
				InitialClaim: types.InitialClaim{
					Enabled:   false,
					MissionID: 21,
				},
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			name: "should allow genesis state with no airdrop supply",
			genState: &types.GenesisState{
				Params: types.DefaultParams(),
				Missions: []types.Mission{
					{
						MissionID: 0,
						Weight:    sdkmath.LegacyOneDec(),
					},
				},
				AirdropSupply: sdk.NewCoin("foo", sdkmath.ZeroInt()),
			},
			valid: true,
		},
		{
			name: "should allow genesis state with no mission",
			genState: &types.GenesisState{
				Params: types.DefaultParams(),
				ClaimRecords: []types.ClaimRecord{
					{
						Address:   sample.AccAddress(),
						Claimable: sdkmath.NewIntFromUint64(10),
					},
					{
						Address:   sample.AccAddress(),
						Claimable: sdkmath.NewIntFromUint64(10),
					},
				},
				AirdropSupply: sdk.NewCoin("foo", sdkmath.NewInt(20)),
			},
			valid: true,
		},
		{
			name: "should allow mission with 0 weight",
			genState: &types.GenesisState{
				Params: types.DefaultParams(),
				ClaimRecords: []types.ClaimRecord{
					{
						Address:   sample.AccAddress(),
						Claimable: sdkmath.NewIntFromUint64(10),
					},
					{
						Address:   sample.AccAddress(),
						Claimable: sdkmath.NewIntFromUint64(10),
					},
				},
				Missions: []types.Mission{
					{
						MissionID: 0,
						Weight:    sdkmath.LegacyOneDec(),
					},
					{
						MissionID: 1,
						Weight:    sdkmath.LegacyZeroDec(),
					},
				},
				AirdropSupply: sdk.NewCoin("foo", sdkmath.NewInt(20)),
			},
			valid: true,
		},
		{
			name: "should allow claim record with completed missions",
			genState: &types.GenesisState{
				Params: types.DefaultParams(),
				ClaimRecords: []types.ClaimRecord{
					{
						Address:           sample.AccAddress(),
						Claimable:         sdkmath.NewIntFromUint64(10),
						CompletedMissions: []uint64{0},
					},
					{
						Address:           sample.AccAddress(),
						Claimable:         sdkmath.NewIntFromUint64(10),
						CompletedMissions: []uint64{1},
					},
				},
				Missions: []types.Mission{
					{
						MissionID: 0,
						Weight:    sdkmath.LegacyMustNewDecFromStr("0.4"),
					},
					{
						MissionID: 1,
						Weight:    sdkmath.LegacyMustNewDecFromStr("0.6"),
					},
				},
				AirdropSupply: sdk.NewCoin("foo", sdkmath.NewInt(10)),
			},
			valid: true,
		},
		{
			name: "should allow claim record with missions all completed",
			genState: &types.GenesisState{
				Params: types.DefaultParams(),
				ClaimRecords: []types.ClaimRecord{
					{
						Address:           sample.AccAddress(),
						Claimable:         sdkmath.NewIntFromUint64(10),
						CompletedMissions: []uint64{0},
					},
					{
						Address:           sample.AccAddress(),
						Claimable:         sdkmath.NewIntFromUint64(10),
						CompletedMissions: []uint64{0, 1},
					},
				},
				Missions: []types.Mission{
					{
						MissionID: 0,
						Weight:    sdkmath.LegacyMustNewDecFromStr("0.4"),
					},
					{
						MissionID: 1,
						Weight:    sdkmath.LegacyMustNewDecFromStr("0.6"),
					},
				},
				AirdropSupply: sdk.NewCoin("foo", sdkmath.NewInt(6)),
			},
			valid: true,
		},
		{
			name: "should allow claim record with zero weight mission completed",
			genState: &types.GenesisState{
				Params: types.DefaultParams(),
				ClaimRecords: []types.ClaimRecord{
					{
						Address:           sample.AccAddress(),
						Claimable:         sdkmath.NewIntFromUint64(10),
						CompletedMissions: []uint64{1},
					},
					{
						Address:           sample.AccAddress(),
						Claimable:         sdkmath.NewIntFromUint64(10),
						CompletedMissions: []uint64{1},
					},
				},
				Missions: []types.Mission{
					{
						MissionID: 0,
						Weight:    sdkmath.LegacyOneDec(),
					},
					{
						MissionID: 1,
						Weight:    sdkmath.LegacyZeroDec(),
					},
				},
				AirdropSupply: sdk.NewCoin("foo", sdkmath.NewInt(20)),
			},
			valid: true,
		},
		{
			name: "should validate genesis state with initial claim enabled",
			genState: &types.GenesisState{
				Params: types.DefaultParams(),
				ClaimRecords: []types.ClaimRecord{
					{
						Address:   sample.AccAddress(),
						Claimable: sdkmath.NewIntFromUint64(10),
					},
					{
						Address:   sample.AccAddress(),
						Claimable: sdkmath.NewIntFromUint64(10),
					},
				},
				Missions: []types.Mission{
					{
						MissionID: 0,
						Weight:    sdkmath.LegacyOneDec(),
					},
				},
				AirdropSupply: sdk.NewCoin("foo", sdkmath.NewInt(20)),
				InitialClaim: types.InitialClaim{
					Enabled:   true,
					MissionID: 0,
				},
			},
			valid: true,
		},
		{
			name: "should prevent validate duplicated claimRecord",
			genState: &types.GenesisState{
				Params: types.DefaultParams(),
				ClaimRecords: []types.ClaimRecord{
					{
						Address:   "duplicate",
						Claimable: sdkmath.NewIntFromUint64(10),
					},
					{
						Address:   "duplicate",
						Claimable: sdkmath.NewIntFromUint64(10),
					},
				},
				Missions: []types.Mission{
					{
						MissionID: 0,
						Weight:    sdkmath.LegacyOneDec(),
					},
				},
				AirdropSupply: sdk.NewCoin("foo", sdkmath.NewInt(20)),
			},
			valid: false,
		},
		{
			name: "should prevent validate claim record with non positive allocation",
			genState: &types.GenesisState{
				Params: types.DefaultParams(),
				ClaimRecords: []types.ClaimRecord{
					{
						Address:   sample.AccAddress(),
						Claimable: sdkmath.NewIntFromUint64(20),
					},
					{
						Address:   sample.AccAddress(),
						Claimable: sdkmath.ZeroInt(),
					},
				},
				Missions: []types.Mission{
					{
						MissionID: 0,
						Weight:    sdkmath.LegacyOneDec(),
					},
				},
				AirdropSupply: sdk.NewCoin("foo", sdkmath.NewInt(20)),
			},
			valid: false,
		},
		{
			name: "should prevent validate airdrop supply higher than sum of claim amounts",
			genState: &types.GenesisState{
				Params: types.DefaultParams(),
				ClaimRecords: []types.ClaimRecord{
					{
						Address:   sample.AccAddress(),
						Claimable: sdkmath.NewIntFromUint64(10),
					},
					{
						Address:   sample.AccAddress(),
						Claimable: sdkmath.NewIntFromUint64(9),
					},
				},
				AirdropSupply: sdk.NewCoin("foo", sdkmath.NewInt(20)),
				Missions: []types.Mission{
					{
						MissionID: 0,
						Weight:    sdkmath.LegacyOneDec(),
					},
				},
			},
			valid: false,
		},
		{
			name: "should prevent validate airdrop supply lower than sum of claim amounts",
			genState: &types.GenesisState{
				Params: types.DefaultParams(),
				ClaimRecords: []types.ClaimRecord{
					{
						Address:   sample.AccAddress(),
						Claimable: sdkmath.NewIntFromUint64(10),
					},
					{
						Address:   sample.AccAddress(),
						Claimable: sdkmath.NewIntFromUint64(11),
					},
				},
				AirdropSupply: sdk.NewCoin("foo", sdkmath.NewInt(20)),
				Missions: []types.Mission{
					{
						MissionID: 0,
						Weight:    sdkmath.LegacyOneDec(),
					},
				},
			},
			valid: false,
		},
		{
			name: "should prevent validate invalid airdrop supply with records with completed missions",
			genState: &types.GenesisState{
				Params: types.DefaultParams(),
				ClaimRecords: []types.ClaimRecord{
					{
						Address:           sample.AccAddress(),
						Claimable:         sdkmath.NewIntFromUint64(10),
						CompletedMissions: []uint64{0},
					},
					{
						Address:   sample.AccAddress(),
						Claimable: sdkmath.NewIntFromUint64(10),
					},
				},
				Missions: []types.Mission{
					{
						MissionID: 0,
						Weight:    sdkmath.LegacyMustNewDecFromStr("0.4"),
					},
					{
						MissionID: 1,
						Weight:    sdkmath.LegacyMustNewDecFromStr("0.6"),
					},
				},
				AirdropSupply: sdk.NewCoin("foo", sdkmath.NewInt(20)),
			},
			valid: false,
		},
		{
			name: "should prevent validate claim record with non existing mission",
			genState: &types.GenesisState{
				Params: types.DefaultParams(),
				ClaimRecords: []types.ClaimRecord{
					{
						Address:           sample.AccAddress(),
						Claimable:         sdkmath.NewIntFromUint64(10),
						CompletedMissions: []uint64{0},
					},
					{
						Address:   sample.AccAddress(),
						Claimable: sdkmath.NewIntFromUint64(10),
					},
				},
				Missions: []types.Mission{
					{
						MissionID: 1,
						Weight:    sdkmath.LegacyOneDec(),
					},
				},
				AirdropSupply: sdk.NewCoin("foo", sdkmath.NewInt(20)),
			},
			valid: false,
		},
		{
			name: "should prevent validate invalid genesis supply coin",
			genState: &types.GenesisState{
				Params:        types.DefaultParams(),
				AirdropSupply: sdk.Coin{},
				Missions: []types.Mission{
					{
						MissionID: 0,
						Weight:    sdkmath.LegacyOneDec(),
					},
				},
			},
			valid: false,
		},
		{
			name: "should prevent validate duplicated mission",
			genState: &types.GenesisState{
				Params: types.DefaultParams(),
				Missions: []types.Mission{
					{
						MissionID: 0,
						Weight:    fiftyPercent,
					},
					{
						MissionID: 0,
						Weight:    fiftyPercent,
					},
				},
			},
			valid: false,
		},
		{
			name: "should prevent validate mission list weights are not equal to 1",
			genState: &types.GenesisState{
				Params: types.DefaultParams(),
				Missions: []types.Mission{
					{
						MissionID: 0,
						Weight:    fiftyPercent,
					},
					{
						MissionID: 0,
						Weight:    sdkmath.LegacyZeroDec(),
					},
				},
			},
			valid: false,
		},
		{
			name: "should prevent validate initial claim enabled with non existing mission",
			genState: &types.GenesisState{
				Params: types.DefaultParams(),
				ClaimRecords: []types.ClaimRecord{
					{
						Address:   sample.AccAddress(),
						Claimable: sdkmath.NewIntFromUint64(10),
					},
					{
						Address:   sample.AccAddress(),
						Claimable: sdkmath.NewIntFromUint64(10),
					},
				},
				AirdropSupply: sdk.NewCoin("foo", sdkmath.NewInt(20)),
				InitialClaim: types.InitialClaim{
					Enabled:   true,
					MissionID: 0,
				},
			},
			valid: false,
		},
		{
			name: "should prevent validate genesis state with invalid param",
			genState: &types.GenesisState{
				Params: types.NewParams(types.DecayInformation{
					Enabled:    true,
					DecayStart: time.UnixMilli(1001),
					DecayEnd:   time.UnixMilli(1000),
				}, time.Unix(0, 0)),
				Missions: []types.Mission{
					{
						MissionID: 0,
						Weight:    sdkmath.LegacyOneDec(),
					},
				},
				AirdropSupply: sdk.NewCoin("foo", sdkmath.ZeroInt()),
			},
			valid: false,
		},
		// this line is used by starport scaffolding # types/genesis/testcase
	} {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.genState.Validate()
			if tt.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}
