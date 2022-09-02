package types_test

import (
	"testing"

	sdkmath "cosmossdk.io/math"
	"github.com/stretchr/testify/require"

	"github.com/ignite/modules/testutil/sample"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ignite/modules/x/claim/types"
)

func TestCheckAirdropSupply(t *testing.T) {
	sampleAddr := sample.Address(r)

	for _, tc := range []struct {
		desc          string
		airdropSupply sdk.Coin
		missionMap    map[uint64]types.Mission
		claimRecords  []types.ClaimRecord
		valid         bool
	}{
		{
			desc:          "valid airdrop supply",
			airdropSupply: sdk.NewCoin("test", sdkmath.NewInt(10)),
			missionMap: map[uint64]types.Mission{
				0: {
					MissionID:   0,
					Description: "",
					Weight:      sdk.ZeroDec(),
				},
			},
			claimRecords: []types.ClaimRecord{
				{
					Address:           sampleAddr,
					Claimable:         sdkmath.NewInt(10),
					CompletedMissions: []uint64{0},
				},
			},
			valid: true,
		}, {
			desc:          "should be invalid with duplicated address in claim record",
			airdropSupply: sdk.NewCoin("test", sdkmath.NewInt(10)),
			missionMap: map[uint64]types.Mission{
				0: {
					MissionID:   0,
					Description: "",
					Weight:      sdk.ZeroDec(),
				},
			},
			claimRecords: []types.ClaimRecord{
				{
					Address:           sampleAddr,
					Claimable:         sdkmath.NewInt(5),
					CompletedMissions: []uint64{0},
				},
				{
					Address:           sampleAddr,
					Claimable:         sdkmath.NewInt(5),
					CompletedMissions: []uint64{},
				},
			},
			valid: false,
		}, {
			desc:          "should be invalid with address completing non existing mission",
			airdropSupply: sdk.NewCoin("test", sdkmath.NewInt(10)),
			missionMap: map[uint64]types.Mission{
				0: {
					MissionID:   0,
					Description: "",
					Weight:      sdk.ZeroDec(),
				},
			},
			claimRecords: []types.ClaimRecord{
				{
					Address:           sampleAddr,
					Claimable:         sdkmath.NewInt(10),
					CompletedMissions: []uint64{0, 2, 3},
				},
			},
			valid: false,
		}, {
			desc:          "should be invalid with airdrop supply not equal to claimable amounts",
			airdropSupply: sdk.NewCoin("test", sdkmath.NewInt(10)),
			missionMap: map[uint64]types.Mission{
				0: {
					MissionID:   0,
					Description: "",
					Weight:      sdk.ZeroDec(),
				},
			},
			claimRecords: []types.ClaimRecord{
				{
					Address:           sampleAddr,
					Claimable:         sdkmath.NewInt(9),
					CompletedMissions: []uint64{0},
				},
			},
			valid: false,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			err := types.CheckAirdropSupply(
				tc.airdropSupply,
				tc.missionMap,
				tc.claimRecords,
			)

			if !tc.valid {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}

		})
	}
}
