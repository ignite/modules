package types_test

import (
	"testing"

	sdkmath "cosmossdk.io/math"
	"github.com/stretchr/testify/require"

	"github.com/ignite/modules/x/claim/types"
	claim "github.com/ignite/modules/x/claim/types"
)

func TestMission_Validate(t *testing.T) {
	for _, tc := range []struct {
		name    string
		mission claim.Mission
		valid   bool
	}{
		{
			name: "should validate valid mission",
			mission: types.Mission{
				MissionID:   1,
				Description: "dummy mission",
				Weight:      sdkmath.LegacyNewDec(r.Int63n(1_000_000)).Quo(sdkmath.LegacyNewDec(1_000_000)),
			},
			valid: true,
		},
		{
			name: "should accept weigth 0",
			mission: claim.Mission{
				MissionID:   uint64(r.Intn(10000)),
				Description: "dummy description",
				Weight:      sdkmath.LegacyMustNewDecFromStr("0"),
			},
			valid: true,
		},
		{
			name: "should accept weight 1",
			mission: claim.Mission{
				MissionID:   uint64(r.Intn(10000)),
				Description: "dummy description",
				Weight:      sdkmath.LegacyMustNewDecFromStr("1"),
			},
			valid: true,
		},
		{
			name: "should prevent weight greater than 1",
			mission: claim.Mission{
				MissionID:   uint64(r.Intn(10000)),
				Description: "dummy description",
				Weight:      sdkmath.LegacyMustNewDecFromStr("1.0000001"),
			},
			valid: false,
		},
		{
			name: "should prevent weight less than 0",
			mission: claim.Mission{
				MissionID:   uint64(r.Intn(10000)),
				Description: "dummy description",
				Weight:      sdkmath.LegacyMustNewDecFromStr("-0.0000001"),
			},
			valid: false,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			require.EqualValues(t, tc.valid, tc.mission.Validate() == nil)
		})
	}
}
