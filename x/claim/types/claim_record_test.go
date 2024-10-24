package types_test

import (
	"testing"

	sdkmath "cosmossdk.io/math"
	"github.com/stretchr/testify/require"

	tc "github.com/ignite/modules/testutil/constructor"
	"github.com/ignite/modules/testutil/sample"
	claim "github.com/ignite/modules/x/claim/types"
)

func TestClaimRecord_Validate(t *testing.T) {
	for _, tc := range []struct {
		name        string
		claimRecord claim.ClaimRecord
		valid       bool
	}{
		{
			name:        "should validate claim record",
			claimRecord: sample.ClaimRecord(r),
			valid:       true,
		},
		{
			name: "should validate claim record with no completed mission",
			claimRecord: claim.ClaimRecord{
				Address:           sample.Address(r),
				Claimable:         sdkmath.OneInt(),
				CompletedMissions: []uint64{},
			},
			valid: true,
		},
		{
			name: "should prevent zero claimable amount",
			claimRecord: claim.ClaimRecord{
				Address:           sample.Address(r),
				Claimable:         sdkmath.ZeroInt(),
				CompletedMissions: []uint64{0, 1, 2},
			},
			valid: false,
		},
		{
			name: "should prevent negative claimable amount",
			claimRecord: claim.ClaimRecord{
				Address:           sample.Address(r),
				Claimable:         sdkmath.NewInt(-1),
				CompletedMissions: []uint64{0, 1, 2},
			},
			valid: false,
		},
		{
			name: "should prevent duplicate completed mission IDs",
			claimRecord: claim.ClaimRecord{
				Address:           sample.Address(r),
				Claimable:         sdkmath.OneInt(),
				CompletedMissions: []uint64{0, 1, 2, 0},
			},
			valid: false,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			require.EqualValues(t, tc.valid, tc.claimRecord.Validate() == nil)
		})
	}
}

func TestClaimRecord_IsMissionCompleted(t *testing.T) {
	for _, tc := range []struct {
		name        string
		claimRecord claim.ClaimRecord
		missionID   uint64
		completed   bool
	}{
		{
			name: "should show completed mission if in list 1",
			claimRecord: claim.ClaimRecord{
				Address:           sample.Address(r),
				Claimable:         sdkmath.OneInt(),
				CompletedMissions: []uint64{0, 1, 2, 3},
			},
			completed: true,
			missionID: 0,
		},
		{
			name: "should show completed mission if in list 2",
			claimRecord: claim.ClaimRecord{
				Address:           sample.Address(r),
				Claimable:         sdkmath.OneInt(),
				CompletedMissions: []uint64{0, 1, 2, 3},
			},
			completed: true,
			missionID: 3,
		},
		{
			name: "should prevent claimRecord with no completed missions",
			claimRecord: claim.ClaimRecord{
				Address:           sample.Address(r),
				Claimable:         sdkmath.OneInt(),
				CompletedMissions: []uint64{},
			},
			completed: false,
			missionID: 0,
		},
		{
			name: "should prevent claimRecord without requested mission",
			claimRecord: claim.ClaimRecord{
				Address:           sample.Address(r),
				Claimable:         sdkmath.OneInt(),
				CompletedMissions: []uint64{1, 2, 3},
			},
			completed: false,
			missionID: 0,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			require.EqualValues(t, tc.completed, tc.claimRecord.IsMissionCompleted(tc.missionID))
		})
	}
}

func TestClaimRecord_ClaimableFromMission(t *testing.T) {
	for _, tt := range []struct {
		name        string
		claimRecord claim.ClaimRecord
		mission     claim.Mission
		expected    sdkmath.Int
	}{
		{
			name: "should allow get claimable from mission with full weight",
			claimRecord: claim.ClaimRecord{
				Claimable: sdkmath.NewIntFromUint64(100),
			},
			mission: claim.Mission{
				Weight: sdkmath.LegacyOneDec(),
			},
			expected: sdkmath.NewIntFromUint64(100),
		},
		{
			name: "should allow get claimable from mission with empty weight",
			claimRecord: claim.ClaimRecord{
				Claimable: sdkmath.NewIntFromUint64(100),
			},
			mission: claim.Mission{
				Weight: sdkmath.LegacyZeroDec(),
			},
			expected: sdkmath.ZeroInt(),
		},
		{
			name: "should allow get claimable from mission with half weight",
			claimRecord: claim.ClaimRecord{
				Claimable: sdkmath.NewIntFromUint64(100),
			},
			mission: claim.Mission{
				Weight: tc.Dec(t, "0.5"),
			},
			expected: sdkmath.NewIntFromUint64(50),
		},
		{
			name: "should allow get claimable and cut decimal",
			claimRecord: claim.ClaimRecord{
				Claimable: sdkmath.NewIntFromUint64(201),
			},
			mission: claim.Mission{
				Weight: tc.Dec(t, "0.5"),
			},
			expected: sdkmath.NewIntFromUint64(100),
		},
		{
			name: "should return no claimable if decimal cut",
			claimRecord: claim.ClaimRecord{
				Claimable: sdkmath.NewIntFromUint64(1),
			},
			mission: claim.Mission{
				Weight: tc.Dec(t, "0.99"),
			},
			expected: sdkmath.NewIntFromUint64(0),
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.claimRecord.ClaimableFromMission(tt.mission)
			require.True(t, got.Equal(tt.expected),
				"expected: %s, got %s",
				tt.expected.String(),
				got.String(),
			)
		})
	}
}
