package types_test

import (
	"errors"
	"testing"
	"time"

	"cosmossdk.io/math"
	"github.com/cometbft/cometbft/crypto"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/ignite/modules/x/fundraising/types"
)

func TestShouldRelease(t *testing.T) {
	now := types.MustParseRFC3339("2021-12-10T00:00:00Z")

	testCases := []struct {
		name      string
		vq        types.VestingQueue
		expResult bool
	}{
		{
			"the release time is already passed the current block time",
			types.VestingQueue{
				AuctionId:   1,
				Auctioneer:  sdk.AccAddress(crypto.AddressHash([]byte("Auctioneer"))).String(),
				PayingCoin:  sdk.NewInt64Coin("denom1", 10000000),
				ReleaseTime: types.MustParseRFC3339("2021-11-01T00:00:00Z"),
				Released:    false,
			},
			true,
		},
		{
			"the release time is exactly the same time as the current block time",
			types.VestingQueue{
				AuctionId:   1,
				Auctioneer:  sdk.AccAddress(crypto.AddressHash([]byte("Auctioneer"))).String(),
				PayingCoin:  sdk.NewInt64Coin("denom1", 10000000),
				ReleaseTime: now,
				Released:    false,
			},
			true,
		},
		{
			"the release time has not passed the current block time",
			types.VestingQueue{
				AuctionId:   1,
				Auctioneer:  sdk.AccAddress(crypto.AddressHash([]byte("Auctioneer"))).String(),
				PayingCoin:  sdk.NewInt64Coin("denom1", 10000000),
				ReleaseTime: types.MustParseRFC3339("2022-01-30T00:00:00Z"),
				Released:    false,
			},
			false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			require.Equal(t, tc.expResult, tc.vq.ShouldRelease(now))
		})
	}
}

func TestValidateVestingSchedules(t *testing.T) {
	for _, tc := range []struct {
		name      string
		schedules []types.VestingSchedule
		endTime   time.Time
		err       error
	}{
		{
			name: "happy case",
			schedules: []types.VestingSchedule{
				{ReleaseTime: types.MustParseRFC3339("9999-01-01T00:00:00Z"), Weight: math.LegacyMustNewDecFromStr("1.0")},
			},
			endTime: types.MustParseRFC3339("2022-03-01T00:00:00Z"),
		},
		{
			name: "invalid case #1",
			schedules: []types.VestingSchedule{
				{ReleaseTime: types.MustParseRFC3339("9999-01-01T00:00:00Z"), Weight: math.LegacyMustNewDecFromStr("-1.0")},
			},
			endTime: types.MustParseRFC3339("2022-03-01T00:00:00Z"),
			err:     errors.New("vesting weight must be positive: invalid vesting schedules"),
		},
		{
			name: "invalid case #2",
			schedules: []types.VestingSchedule{
				{ReleaseTime: types.MustParseRFC3339("2022-01-01T00:00:00Z"), Weight: math.LegacyMustNewDecFromStr("1.0")},
			},
			endTime: types.MustParseRFC3339("2022-03-01T00:00:00Z"),
			err:     errors.New("release time must be set after the end time: invalid vesting schedules"),
		},
		{
			name: "invalid case #3",
			schedules: []types.VestingSchedule{
				{ReleaseTime: types.MustParseRFC3339("9999-01-01T00:00:00Z"), Weight: math.LegacyMustNewDecFromStr("2.0")},
			},
			endTime: types.MustParseRFC3339("2022-03-01T00:00:00Z"),
			err:     errors.New("vesting weight must not be greater than 1: invalid vesting schedules"),
		},
		{
			name: "invalid case #4",
			schedules: []types.VestingSchedule{
				{ReleaseTime: types.MustParseRFC3339("2022-06-01T00:00:00Z"), Weight: math.LegacyMustNewDecFromStr("0.25")},
				{ReleaseTime: types.MustParseRFC3339("2022-04-01T00:00:00Z"), Weight: math.LegacyMustNewDecFromStr("0.25")},
				{ReleaseTime: types.MustParseRFC3339("2022-09-01T00:00:00Z"), Weight: math.LegacyMustNewDecFromStr("0.25")},
				{ReleaseTime: types.MustParseRFC3339("2022-12-01T00:00:00Z"), Weight: math.LegacyMustNewDecFromStr("0.25")},
			},
			endTime: types.MustParseRFC3339("2022-03-01T00:00:00Z"),
			err:     errors.New("release time must be chronological: invalid vesting schedules"),
		},
		{
			name: "invalid case #5",
			schedules: []types.VestingSchedule{
				{ReleaseTime: types.MustParseRFC3339("2022-05-01T00:00:00Z"), Weight: math.LegacyMustNewDecFromStr("0.5")},
				{ReleaseTime: types.MustParseRFC3339("2022-06-01T00:00:00Z"), Weight: math.LegacyMustNewDecFromStr("0.5")},
				{ReleaseTime: types.MustParseRFC3339("2022-07-01T00:00:00Z"), Weight: math.LegacyMustNewDecFromStr("0.5")},
				{ReleaseTime: types.MustParseRFC3339("2022-08-01T00:00:00Z"), Weight: math.LegacyMustNewDecFromStr("0.5")},
			},
			endTime: types.MustParseRFC3339("2022-03-01T00:00:00Z"),
			err:     errors.New("total vesting weight must be equal to 1: invalid vesting schedules"),
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			err := types.ValidateVestingSchedules(tc.schedules, tc.endTime)
			if tc.err != nil {
				require.Error(t, err)
				require.EqualError(t, err, tc.err.Error())
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestSetReleased(t *testing.T) {
	vestingQueue := types.NewVestingQueue(
		1,
		sdk.AccAddress(crypto.AddressHash([]byte("Auctioneer"))),
		sdk.NewInt64Coin("denom1", 10000000),
		types.MustParseRFC3339("2021-11-01T00:00:00Z"),
		false,
	)
	require.False(t, vestingQueue.Released)

	vestingQueue.SetReleased(true)
	require.True(t, vestingQueue.Released)
}
