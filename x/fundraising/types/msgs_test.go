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

func TestMsgCreateFixedPriceAuction(t *testing.T) {
	testCases := []struct {
		desc string
		msg  *types.MsgCreateFixedPriceAuction
		err  error
	}{
		{
			desc: "valid fixed price auction",
			msg: types.NewMsgCreateFixedPriceAuction(
				sdk.AccAddress(crypto.AddressHash([]byte("Auctioneer"))).String(),
				math.LegacyMustNewDecFromStr("0.5"),
				sdk.NewInt64Coin("denom2", 10_000_000_000_000),
				"denom1",
				[]types.VestingSchedule{},
				time.Now(),
				time.Now().AddDate(0, 1, 0),
			),
		},
		{
			desc: "start price must be positive",
			err:  errors.New("start price must be positive: invalid request"),
			msg: types.NewMsgCreateFixedPriceAuction(
				sdk.AccAddress(crypto.AddressHash([]byte("Auctioneer"))).String(),
				math.LegacyMustNewDecFromStr("0"),
				sdk.NewInt64Coin("denom2", 10_000_000_000_000),
				"denom1",
				[]types.VestingSchedule{},
				time.Now(),
				time.Now().AddDate(0, 1, 0),
			),
		},
		{
			desc: "selling coin amount must be positive",
			err:  errors.New("selling coin amount must be positive: invalid request"),
			msg: types.NewMsgCreateFixedPriceAuction(
				sdk.AccAddress(crypto.AddressHash([]byte("Auctioneer"))).String(),
				math.LegacyMustNewDecFromStr("0.5"),
				sdk.NewInt64Coin("denom2", 0),
				"denom1",
				[]types.VestingSchedule{},
				time.Now(),
				time.Now().AddDate(0, 1, 0),
			),
		},
		{
			desc: "selling coin denom must not be the same as paying coin denom",
			err:  errors.New("selling coin denom must not be the same as paying coin denom: invalid request"),
			msg: types.NewMsgCreateFixedPriceAuction(
				sdk.AccAddress(crypto.AddressHash([]byte("Auctioneer"))).String(),
				math.LegacyMustNewDecFromStr("0.5"),
				sdk.NewInt64Coin("denom2", 10_000_000_000_000),
				"denom2",
				[]types.VestingSchedule{},
				time.Now(),
				time.Now().AddDate(0, 1, 0),
			),
		},
		{
			desc: "end time must be set after start time",
			err:  errors.New("end time must be set after start time: invalid request"),
			msg: types.NewMsgCreateFixedPriceAuction(
				sdk.AccAddress(crypto.AddressHash([]byte("Auctioneer"))).String(),
				math.LegacyMustNewDecFromStr("0.5"),
				sdk.NewInt64Coin("denom2", 10_000_000_000_000),
				"denom1",
				[]types.VestingSchedule{},
				time.Now(),
				time.Now().AddDate(-1, 0, 0),
			),
		},
		{
			desc: "vesting weight must be positive",
			err:  errors.New("vesting weight must be positive: invalid vesting schedules"),
			msg: types.NewMsgCreateFixedPriceAuction(
				sdk.AccAddress(crypto.AddressHash([]byte("Auctioneer"))).String(),
				math.LegacyMustNewDecFromStr("0.5"),
				sdk.NewInt64Coin("denom2", 10_000_000_000_000),
				"denom1",
				[]types.VestingSchedule{
					{
						ReleaseTime: time.Now().AddDate(0, 1, 0).AddDate(0, 6, 0),
						Weight:      math.LegacyZeroDec(),
					},
				},
				time.Now(),
				time.Now().AddDate(0, 1, 0),
			),
		},
		{
			desc: "vesting weight must not be greater than 1",
			err:  errors.New("vesting weight must not be greater than 1: invalid vesting schedules"),
			msg: types.NewMsgCreateFixedPriceAuction(
				sdk.AccAddress(crypto.AddressHash([]byte("Auctioneer"))).String(),
				math.LegacyMustNewDecFromStr("0.5"),
				sdk.NewInt64Coin("denom2", 10_000_000_000_000),
				"denom1",
				[]types.VestingSchedule{
					{
						ReleaseTime: time.Now().AddDate(0, 1, 0).AddDate(0, 6, 0),
						Weight:      math.LegacyMustNewDecFromStr("1.1"),
					},
				},
				time.Now(),
				time.Now().AddDate(0, 1, 0),
			),
		},
		{
			desc: "release time must be set after the end time",
			err:  errors.New("release time must be set after the end time: invalid vesting schedules"),
			msg: types.NewMsgCreateFixedPriceAuction(
				sdk.AccAddress(crypto.AddressHash([]byte("Auctioneer"))).String(),
				math.LegacyMustNewDecFromStr("0.5"),
				sdk.NewInt64Coin("denom2", 10_000_000_000_000),
				"denom1",
				[]types.VestingSchedule{
					{
						ReleaseTime: types.MustParseRFC3339("2022-06-01T22:08:41+00:00"),
						Weight:      math.LegacyMustNewDecFromStr("1.0"),
					},
				},
				time.Now(),
				time.Now().AddDate(1, 0, 0),
			),
		},
		{
			desc: "release time must be chronological",
			err:  errors.New("release time must be chronological: invalid vesting schedules"),
			msg: types.NewMsgCreateFixedPriceAuction(
				sdk.AccAddress(crypto.AddressHash([]byte("Auctioneer"))).String(),
				math.LegacyMustNewDecFromStr("0.5"),
				sdk.NewInt64Coin("denom2", 10_000_000_000_000),
				"denom1",
				[]types.VestingSchedule{
					{
						ReleaseTime: time.Now().AddDate(0, 1, 0).AddDate(0, 6, 0),
						Weight:      math.LegacyMustNewDecFromStr("0.5"),
					},
					{
						ReleaseTime: time.Now().AddDate(0, 1, 0).AddDate(0, 3, 0),
						Weight:      math.LegacyMustNewDecFromStr("0.5"),
					},
				},
				time.Now(),
				time.Now().AddDate(0, 1, 0),
			),
		},
		{
			desc: "total vesting weight must be equal to 1",
			err:  errors.New("total vesting weight must be equal to 1: invalid vesting schedules"),
			msg: types.NewMsgCreateFixedPriceAuction(
				sdk.AccAddress(crypto.AddressHash([]byte("Auctioneer"))).String(),
				math.LegacyMustNewDecFromStr("0.5"),
				sdk.NewInt64Coin("denom2", 10_000_000_000_000),
				"denom1",
				[]types.VestingSchedule{
					{
						ReleaseTime: time.Now().AddDate(0, 1, 0).AddDate(0, 6, 0),
						Weight:      math.LegacyMustNewDecFromStr("0.5"),
					},
					{
						ReleaseTime: time.Now().AddDate(0, 1, 0).AddDate(1, 0, 0),
						Weight:      math.LegacyMustNewDecFromStr("0.3"),
					},
				},
				time.Now(),
				time.Now().AddDate(0, 1, 0),
			),
		},
		{
			desc: "start price must be positive",
			err:  errors.New("start price must be positive: invalid request"),
			msg: types.NewMsgCreateFixedPriceAuction(
				sdk.AccAddress(crypto.AddressHash([]byte("Auctioneer"))).String(),
				math.LegacyMustNewDecFromStr("0"),
				sdk.NewInt64Coin("denom2", 10_000_000_000_000),
				"denom1",
				[]types.VestingSchedule{},
				time.Now(),
				time.Now().AddDate(0, 1, 0),
			),
		},
		{
			desc: "selling coin amount must be positive",
			err:  errors.New("selling coin amount must be positive: invalid request"),
			msg: types.NewMsgCreateFixedPriceAuction(
				sdk.AccAddress(crypto.AddressHash([]byte("Auctioneer"))).String(),
				math.LegacyMustNewDecFromStr("0.5"),
				sdk.NewInt64Coin("denom2", 0),
				"denom1",
				[]types.VestingSchedule{},
				time.Now(),
				time.Now().AddDate(0, 1, 0),
			),
		},
		{
			desc: "selling coin denom must not be the same as paying coin denom",
			err:  errors.New("selling coin denom must not be the same as paying coin denom: invalid request"),
			msg: types.NewMsgCreateFixedPriceAuction(
				sdk.AccAddress(crypto.AddressHash([]byte("Auctioneer"))).String(),
				math.LegacyMustNewDecFromStr("0.5"),
				sdk.NewInt64Coin("denom2", 10_000_000_000_000),
				"denom2",
				[]types.VestingSchedule{},
				time.Now(),
				time.Now().AddDate(0, 1, 0),
			),
		},
		{
			desc: "end time must be set after start time",
			err:  errors.New("end time must be set after start time: invalid request"),
			msg: types.NewMsgCreateFixedPriceAuction(
				sdk.AccAddress(crypto.AddressHash([]byte("Auctioneer"))).String(),
				math.LegacyMustNewDecFromStr("0.5"),
				sdk.NewInt64Coin("denom2", 10_000_000_000_000),
				"denom1",
				[]types.VestingSchedule{},
				time.Now(),
				time.Now().AddDate(-1, 0, 0),
			),
		},
		{
			desc: "vesting weight must be positive",
			err:  errors.New("vesting weight must be positive: invalid vesting schedules"),
			msg: types.NewMsgCreateFixedPriceAuction(
				sdk.AccAddress(crypto.AddressHash([]byte("Auctioneer"))).String(),
				math.LegacyMustNewDecFromStr("0.5"),
				sdk.NewInt64Coin("denom2", 10_000_000_000_000),
				"denom1",
				[]types.VestingSchedule{
					{
						ReleaseTime: time.Now().AddDate(0, 1, 0).AddDate(0, 6, 0),
						Weight:      math.LegacyZeroDec(),
					},
				},
				time.Now(),
				time.Now().AddDate(0, 1, 0),
			),
		},
		{
			desc: "vesting weight must not be greater than 1",
			err:  errors.New("vesting weight must not be greater than 1: invalid vesting schedules"),
			msg: types.NewMsgCreateFixedPriceAuction(
				sdk.AccAddress(crypto.AddressHash([]byte("Auctioneer"))).String(),
				math.LegacyMustNewDecFromStr("0.5"),
				sdk.NewInt64Coin("denom2", 10_000_000_000_000),
				"denom1",
				[]types.VestingSchedule{
					{
						ReleaseTime: time.Now().AddDate(0, 1, 0).AddDate(0, 6, 0),
						Weight:      math.LegacyMustNewDecFromStr("1.1"),
					},
				},
				time.Now(),
				time.Now().AddDate(0, 1, 0),
			),
		},
		{
			desc: "release time must be set after the end time",
			err:  errors.New("release time must be set after the end time: invalid vesting schedules"),
			msg: types.NewMsgCreateFixedPriceAuction(
				sdk.AccAddress(crypto.AddressHash([]byte("Auctioneer"))).String(),
				math.LegacyMustNewDecFromStr("0.5"),
				sdk.NewInt64Coin("denom2", 10_000_000_000_000),
				"denom1",
				[]types.VestingSchedule{
					{
						ReleaseTime: types.MustParseRFC3339("2022-06-01T22:08:41+00:00"),
						Weight:      math.LegacyMustNewDecFromStr("1.0"),
					},
				},
				time.Now(),
				time.Now().AddDate(1, 0, 0),
			),
		},
		{
			desc: "release time must be chronological",
			err:  errors.New("release time must be chronological: invalid vesting schedules"),
			msg: types.NewMsgCreateFixedPriceAuction(
				sdk.AccAddress(crypto.AddressHash([]byte("Auctioneer"))).String(),
				math.LegacyMustNewDecFromStr("0.5"),
				sdk.NewInt64Coin("denom2", 10_000_000_000_000),
				"denom1",
				[]types.VestingSchedule{
					{
						ReleaseTime: time.Now().AddDate(0, 1, 0).AddDate(0, 6, 0),
						Weight:      math.LegacyMustNewDecFromStr("0.5"),
					},
					{
						ReleaseTime: time.Now().AddDate(0, 1, 0).AddDate(0, 3, 0),
						Weight:      math.LegacyMustNewDecFromStr("0.5"),
					},
				},
				time.Now(),
				time.Now().AddDate(0, 1, 0),
			),
		},
		{
			desc: "total vesting weight must be equal to 1",
			err:  errors.New("total vesting weight must be equal to 1: invalid vesting schedules"),
			msg: types.NewMsgCreateFixedPriceAuction(
				sdk.AccAddress(crypto.AddressHash([]byte("Auctioneer"))).String(),
				math.LegacyMustNewDecFromStr("0.5"),
				sdk.NewInt64Coin("denom2", 10_000_000_000_000),
				"denom1",
				[]types.VestingSchedule{
					{
						ReleaseTime: time.Now().AddDate(0, 1, 0).AddDate(0, 6, 0),
						Weight:      math.LegacyMustNewDecFromStr("0.5"),
					},
					{
						ReleaseTime: time.Now().AddDate(0, 1, 0).AddDate(1, 0, 0),
						Weight:      math.LegacyMustNewDecFromStr("0.3"),
					},
				},
				time.Now(),
				time.Now().AddDate(0, 1, 0),
			),
		},
	}
	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.msg.Validate()
			if tc.err != nil {
				require.Error(t, err)
				require.Equal(t, err.Error(), err.Error())
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestMsgCreateBatchAuction(t *testing.T) {
	testCases := []struct {
		desc string
		msg  *types.MsgCreateBatchAuction
		err  error
	}{
		{
			desc: "valid batch auction should pass",
			msg: types.NewMsgCreateBatchAuction(
				sdk.AccAddress(crypto.AddressHash([]byte("Auctioneer"))).String(),
				math.LegacyMustNewDecFromStr("0.5"),
				math.LegacyMustNewDecFromStr("0.1"),
				sdk.NewInt64Coin("denom2", 10_000_000_000_000),
				"denom1",
				[]types.VestingSchedule{},
				uint32(2),
				math.LegacyMustNewDecFromStr("0.05"),
				time.Now(),
				time.Now().AddDate(0, 1, 0),
			),
		},
		{
			desc: "start price must be positive",
			err:  errors.New("start price must be positive: invalid request"),
			msg: types.NewMsgCreateBatchAuction(
				sdk.AccAddress(crypto.AddressHash([]byte("Auctioneer"))).String(),
				math.LegacyMustNewDecFromStr("0"),
				math.LegacyMustNewDecFromStr("0.1"),
				sdk.NewInt64Coin("denom2", 10_000_000_000_000),
				"denom1",
				[]types.VestingSchedule{},
				uint32(2),
				math.LegacyMustNewDecFromStr("0.05"),
				time.Now(),
				time.Now().AddDate(0, 1, 0),
			),
		},
		{
			desc: "minimum price must be positive",
			err:  errors.New("minimum price must be positive: invalid request"),
			msg: types.NewMsgCreateBatchAuction(
				sdk.AccAddress(crypto.AddressHash([]byte("Auctioneer"))).String(),
				math.LegacyMustNewDecFromStr("0.1"),
				math.LegacyMustNewDecFromStr("0"),
				sdk.NewInt64Coin("denom2", 10_000_000_000_000),
				"denom1",
				[]types.VestingSchedule{},
				uint32(2),
				math.LegacyMustNewDecFromStr("0.05"),
				time.Now(),
				time.Now().AddDate(0, 1, 0),
			),
		},
		{
			desc: "selling coin amount must be positive",
			err:  errors.New("selling coin amount must be positive: invalid request"),
			msg: types.NewMsgCreateBatchAuction(
				sdk.AccAddress(crypto.AddressHash([]byte("Auctioneer"))).String(),
				math.LegacyMustNewDecFromStr("0.5"),
				math.LegacyMustNewDecFromStr("0.1"),
				sdk.NewInt64Coin("denom2", 0),
				"denom1",
				[]types.VestingSchedule{},
				uint32(2),
				math.LegacyMustNewDecFromStr("0.05"),
				time.Now(),
				time.Now().AddDate(0, 1, 0),
			),
		},
		{
			desc: "selling coin denom must not be the same as paying coin denom",
			err:  errors.New("selling coin denom must not be the same as paying coin denom: invalid request"),
			msg: types.NewMsgCreateBatchAuction(
				sdk.AccAddress(crypto.AddressHash([]byte("Auctioneer"))).String(),
				math.LegacyMustNewDecFromStr("0.5"),
				math.LegacyMustNewDecFromStr("0.1"),
				sdk.NewInt64Coin("denom2", 10_000_000_000_000),
				"denom2",
				[]types.VestingSchedule{},
				uint32(2),
				math.LegacyMustNewDecFromStr("0.05"),
				time.Now(),
				time.Now().AddDate(0, 1, 0),
			),
		},
		{
			desc: "end time must be set after start time",
			err:  errors.New("end time must be set after start time: invalid request"),
			msg: types.NewMsgCreateBatchAuction(
				sdk.AccAddress(crypto.AddressHash([]byte("Auctioneer"))).String(),
				math.LegacyMustNewDecFromStr("0.5"),
				math.LegacyMustNewDecFromStr("0.1"),
				sdk.NewInt64Coin("denom2", 10_000_000_000_000),
				"denom1",
				[]types.VestingSchedule{},
				uint32(2),
				math.LegacyMustNewDecFromStr("0.05"),
				time.Now(),
				time.Now().AddDate(-1, 0, 0),
			),
		},
		{
			desc: "vesting weight must be positive",
			err:  errors.New("vesting weight must be positive: invalid vesting schedules"),
			msg: types.NewMsgCreateBatchAuction(
				sdk.AccAddress(crypto.AddressHash([]byte("Auctioneer"))).String(),
				math.LegacyMustNewDecFromStr("0.5"),
				math.LegacyMustNewDecFromStr("0.1"),
				sdk.NewInt64Coin("denom2", 10_000_000_000_000),
				"denom1",
				[]types.VestingSchedule{
					{
						ReleaseTime: time.Now().AddDate(0, 1, 0).AddDate(0, 6, 0),
						Weight:      math.LegacyZeroDec(),
					},
				},
				uint32(2),
				math.LegacyMustNewDecFromStr("0.05"),
				time.Now(),
				time.Now().AddDate(0, 1, 0),
			),
		},
		{
			desc: "vesting weight must not be greater than 1",
			err:  errors.New("vesting weight must not be greater than 1: invalid vesting schedules"),
			msg: types.NewMsgCreateBatchAuction(
				sdk.AccAddress(crypto.AddressHash([]byte("Auctioneer"))).String(),
				math.LegacyMustNewDecFromStr("0.5"),
				math.LegacyMustNewDecFromStr("0.1"),
				sdk.NewInt64Coin("denom2", 10_000_000_000_000),
				"denom1",
				[]types.VestingSchedule{
					{
						ReleaseTime: time.Now().AddDate(0, 1, 0).AddDate(0, 6, 0),
						Weight:      math.LegacyMustNewDecFromStr("1.1"),
					},
				},
				uint32(2),
				math.LegacyMustNewDecFromStr("0.05"),
				time.Now(),
				time.Now().AddDate(0, 1, 0),
			),
		},
		{
			desc: "release time must be set after the end time",
			err:  errors.New("release time must be set after the end time: invalid vesting schedules"),
			msg: types.NewMsgCreateBatchAuction(
				sdk.AccAddress(crypto.AddressHash([]byte("Auctioneer"))).String(),
				math.LegacyMustNewDecFromStr("0.5"),
				math.LegacyMustNewDecFromStr("0.1"),
				sdk.NewInt64Coin("denom2", 10_000_000_000_000),
				"denom1",
				[]types.VestingSchedule{
					{
						ReleaseTime: time.Now(),
						Weight:      math.LegacyMustNewDecFromStr("1.0"),
					},
				},
				uint32(2),
				math.LegacyMustNewDecFromStr("0.05"),
				time.Now(),
				time.Now().AddDate(1, 0, 0),
			),
		},
		{
			desc: "release time must be chronological",
			err:  errors.New("release time must be chronological: invalid vesting schedules"),
			msg: types.NewMsgCreateBatchAuction(
				sdk.AccAddress(crypto.AddressHash([]byte("Auctioneer"))).String(),
				math.LegacyMustNewDecFromStr("0.5"),
				math.LegacyMustNewDecFromStr("0.1"),
				sdk.NewInt64Coin("denom2", 10_000_000_000_000),
				"denom1",
				[]types.VestingSchedule{
					{
						ReleaseTime: time.Now().AddDate(0, 1, 0).AddDate(0, 6, 0),
						Weight:      math.LegacyMustNewDecFromStr("0.5"),
					},
					{
						ReleaseTime: time.Now().AddDate(0, 1, 0).AddDate(0, 3, 0),
						Weight:      math.LegacyMustNewDecFromStr("0.5"),
					},
				},
				uint32(2),
				math.LegacyMustNewDecFromStr("0.05"),
				time.Now(),
				time.Now().AddDate(0, 1, 0),
			),
		},
		{
			desc: "total vesting weight must be equal to 1",
			err:  errors.New("total vesting weight must be equal to 1: invalid vesting schedules"),
			msg: types.NewMsgCreateBatchAuction(
				sdk.AccAddress(crypto.AddressHash([]byte("Auctioneer"))).String(),
				math.LegacyMustNewDecFromStr("0.5"),
				math.LegacyMustNewDecFromStr("0.1"),
				sdk.NewInt64Coin("denom2", 10_000_000_000_000),
				"denom1",
				[]types.VestingSchedule{
					{
						ReleaseTime: time.Now().AddDate(0, 1, 0).AddDate(0, 6, 0),
						Weight:      math.LegacyMustNewDecFromStr("0.5"),
					},
					{
						ReleaseTime: time.Now().AddDate(0, 1, 0).AddDate(1, 0, 0),
						Weight:      math.LegacyMustNewDecFromStr("0.3"),
					},
				},
				uint32(2),
				math.LegacyMustNewDecFromStr("0.05"),
				time.Now(),
				time.Now().AddDate(0, 1, 0),
			),
		},
		{
			desc: "extend rate must be positive",
			err:  errors.New("extend rate must be positive: invalid request"),
			msg: types.NewMsgCreateBatchAuction(
				sdk.AccAddress(crypto.AddressHash([]byte("Auctioneer"))).String(),
				math.LegacyMustNewDecFromStr("0.5"),
				math.LegacyMustNewDecFromStr("0.1"),
				sdk.NewInt64Coin("denom2", 10_000_000_000_000),
				"denom1",
				[]types.VestingSchedule{},
				uint32(2),
				math.LegacyMustNewDecFromStr("-0.05"),
				time.Now(),
				time.Now().AddDate(0, 1, 0),
			),
		},
	}
	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.msg.Validate()
			if tc.err != nil {
				require.Error(t, err)
				require.Equal(t, err.Error(), err.Error())
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestMsgPlaceBid(t *testing.T) {
	testCases := []struct {
		desc string
		msg  *types.MsgPlaceBid
		err  error
	}{
		{
			desc: "valid bid",
			msg: types.NewMsgPlaceBid(
				uint64(1),
				sdk.AccAddress(crypto.AddressHash([]byte("Bidder"))).String(),
				types.BidTypeBatchWorth,
				math.LegacyOneDec(),
				sdk.NewInt64Coin("denom2", 1000000),
			),
		},
		{
			desc: "bid price must be positive value",
			err:  errors.New("bid price must be positive value: invalid request"),
			msg: types.NewMsgPlaceBid(
				uint64(1),
				sdk.AccAddress(crypto.AddressHash([]byte("Bidder"))).String(),
				types.BidTypeBatchWorth,
				math.LegacyZeroDec(),
				sdk.NewInt64Coin("denom2", 1000000),
			),
		},
		{
			desc: "invalid coin amount",
			err:  errors.New("invalid coin amount: 0: invalid request"),
			msg: types.NewMsgPlaceBid(
				uint64(1),
				sdk.AccAddress(crypto.AddressHash([]byte("Bidder"))).String(),
				types.BidTypeBatchWorth,
				math.LegacyOneDec(),
				sdk.NewInt64Coin("denom2", 0),
			),
		},
	}
	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.msg.Validate()
			if tc.err != nil {
				require.Error(t, err)
				require.Equal(t, err.Error(), err.Error())
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestMsgModifyBid(t *testing.T) {
	testCases := []struct {
		desc string
		msg  *types.MsgModifyBid
		err  error
	}{
		{
			desc: "valid bid modification",
			msg: types.NewMsgModifyBid(
				uint64(1),
				sdk.AccAddress(crypto.AddressHash([]byte("Bidder"))).String(),
				uint64(0),
				math.LegacyOneDec(),
				sdk.NewInt64Coin("denom2", 1000000),
			),
		},
		{
			desc: "bid price must be positive value",
			err:  errors.New("bid price must be positive value: invalid request"),
			msg: types.NewMsgModifyBid(
				uint64(1),
				sdk.AccAddress(crypto.AddressHash([]byte("Bidder"))).String(),
				uint64(0),
				math.LegacyZeroDec(),
				sdk.NewInt64Coin("denom2", 1000000),
			),
		},
		{
			desc: "invalid coin amount",
			err:  errors.New("invalid coin amount: 0: invalid request"),
			msg: types.NewMsgModifyBid(
				uint64(1),
				sdk.AccAddress(crypto.AddressHash([]byte("Bidder"))).String(),
				uint64(0),
				math.LegacyOneDec(),
				sdk.NewInt64Coin("denom2", 0),
			),
		},
	}
	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.msg.Validate()
			if tc.err != nil {
				require.Error(t, err)
				require.Equal(t, err.Error(), err.Error())
				return
			}
			require.NoError(t, err)
		})
	}
}
