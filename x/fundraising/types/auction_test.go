package types_test

import (
	"testing"
	"time"

	"cosmossdk.io/math"
	"github.com/cometbft/cometbft/crypto"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/ignite/modules/x/fundraising/types"
)

func TestUnpackAuction(t *testing.T) {
	auction := types.NewFixedPriceAuction(
		types.NewBaseAuction(
			1,
			types.AuctionTypeFixedPrice,
			sdk.AccAddress(crypto.AddressHash([]byte("Auctioneer"))).String(),
			types.SellingReserveAddress(1).String(),
			types.PayingReserveAddress(1).String(),
			math.LegacyMustNewDecFromStr("0.5"),
			sdk.NewInt64Coin("denom3", 1_000_000_000_000),
			"denom4",
			types.VestingReserveAddress(1).String(),
			[]types.VestingSchedule{
				{ReleaseTime: types.MustParseRFC3339("2023-01-01T00:00:00Z"), Weight: math.LegacyOneDec()},
			},
			types.MustParseRFC3339("2022-01-01T00:00:00Z"),
			[]time.Time{types.MustParseRFC3339("2022-02-01T00:00:00Z")},
			types.AuctionStatusStarted,
		),
		sdk.NewInt64Coin("denom3", 1_000_000_000_000),
	)

	a, err := types.PackAuction(auction)
	require.NoError(t, err)

	marshaled, err := a.Marshal()
	require.NoError(t, err)

	var any2 codectypes.Any
	err = any2.Unmarshal(marshaled)
	require.NoError(t, err)

	reMarshal, err := any2.Marshal()
	require.NoError(t, err)
	require.Equal(t, marshaled, reMarshal)

	auction2, err := types.UnpackAuction(&any2)
	require.NoError(t, err)

	require.Equal(t, auction.AuctionId, auction2.GetId())
	require.Equal(t, auction.Type, auction2.GetType())
	require.Equal(t, auction.Auctioneer, auction2.GetAuctioneer())
	require.Equal(t, auction.SellingCoin, auction2.GetSellingCoin())
	require.Equal(t, auction.PayingCoinDenom, auction2.GetPayingCoinDenom())
	require.Equal(t, auction.StartPrice, auction2.GetStartPrice())
	require.Equal(t, auction.SellingReserveAddress, auction2.GetSellingReserveAddress())
	require.Equal(t, auction.SellingReserveAddress, auction2.GetSellingReserveAddress())
	require.Equal(t, auction.PayingReserveAddress, auction2.GetPayingReserveAddress())
	require.Equal(t, auction.VestingReserveAddress, auction2.GetVestingReserveAddress())
	require.Equal(t, auction.VestingSchedules, auction2.GetVestingSchedules())
	require.Equal(t, auction.StartTime.UTC(), auction2.GetStartTime().UTC())
	require.Equal(t, auction.EndTime[0].UTC(), auction2.GetEndTime()[0].UTC())
	require.Equal(t, auction.Status, auction2.GetStatus())

	require.NoError(t, auction2.SetId(5))
	require.NoError(t, auction2.SetType(types.AuctionTypeBatch))
	require.NoError(t, auction2.SetAuctioneer(sdk.AccAddress(crypto.AddressHash([]byte("Auctioneer2"))).String()))
	require.NoError(t, auction2.SetSellingReserveAddress(types.SellingReserveAddress(5).String()))
	require.NoError(t, auction2.SetPayingReserveAddress(types.PayingReserveAddress(5).String()))
	require.NoError(t, auction2.SetVestingReserveAddress(types.VestingReserveAddress(5).String()))
	require.NoError(t, auction2.SetStartPrice(math.LegacyOneDec()))
	require.NoError(t, auction2.SetSellingCoin(sdk.NewInt64Coin("denom5", 1_000_000_000_000)))
	require.NoError(t, auction2.SetPayingCoinDenom("denom6"))
	require.NoError(t, auction2.SetStartTime(types.MustParseRFC3339("2022-10-01T00:00:00Z")))
	require.NoError(t, auction2.SetVestingSchedules([]types.VestingSchedule{{ReleaseTime: types.MustParseRFC3339("2023-01-01T00:00:00Z"), Weight: math.LegacyOneDec()}}))
	require.NoError(t, auction2.SetEndTime([]time.Time{types.MustParseRFC3339("2022-11-01T00:00:00Z")}))
	require.NoError(t, auction2.SetStatus(types.AuctionStatusStarted))

	require.True(t, auction2.GetId() == 5)
	require.True(t, auction2.GetType() == types.AuctionTypeBatch)
	require.Equal(t, auction2.GetAuctioneer(), sdk.AccAddress(crypto.AddressHash([]byte("Auctioneer2"))).String())
	require.Equal(t, auction2.GetSellingReserveAddress(), types.SellingReserveAddress(5).String())
	require.Equal(t, auction2.GetPayingReserveAddress(), types.PayingReserveAddress(5).String())
	require.Equal(t, auction2.GetVestingReserveAddress(), types.VestingReserveAddress(5).String())
	require.True(t, auction2.GetStartPrice().Equal(math.LegacyOneDec()))
}

func TestUnpackAuctionJSON(t *testing.T) {
	auction := types.NewFixedPriceAuction(
		types.NewBaseAuction(
			1,
			types.AuctionTypeFixedPrice,
			sdk.AccAddress(crypto.AddressHash([]byte("Auctioneer"))).String(),
			types.SellingReserveAddress(1).String(),
			types.PayingReserveAddress(1).String(),
			math.LegacyMustNewDecFromStr("0.5"),
			sdk.NewInt64Coin("denom1", 1_000_000_000_000),
			"denom2",
			types.VestingReserveAddress(1).String(),
			[]types.VestingSchedule{},
			time.Now().AddDate(0, 0, -1),
			[]time.Time{time.Now().AddDate(0, 1, -1)},
			types.AuctionStatusStarted,
		),
		sdk.NewInt64Coin("denom2", 1_000_000_000_000),
	)

	a, err := types.PackAuction(auction)
	require.NoError(t, err)

	registry := codectypes.NewInterfaceRegistry()
	types.RegisterInterfaces(registry)
	cdc := codec.NewProtoCodec(registry)

	bz := cdc.MustMarshalJSON(a)

	var any2 codectypes.Any
	err = cdc.UnmarshalJSON(bz, &any2)
	require.NoError(t, err)

	auction2, err := types.UnpackAuction(&any2)
	require.NoError(t, err)

	require.Equal(t, uint64(1), auction2.GetId())
}

func TestUnpackAuctions(t *testing.T) {
	auction := []types.AuctionI{
		types.NewFixedPriceAuction(
			types.NewBaseAuction(
				1,
				types.AuctionTypeFixedPrice,
				sdk.AccAddress(crypto.AddressHash([]byte("Auctioneer"))).String(),
				types.SellingReserveAddress(1).String(),
				types.PayingReserveAddress(1).String(),
				math.LegacyMustNewDecFromStr("0.5"),
				sdk.NewInt64Coin("denom1", 1_000_000_000_000),
				"denom2",
				types.VestingReserveAddress(1).String(),
				[]types.VestingSchedule{},
				time.Now().AddDate(0, 0, -1),
				[]time.Time{time.Now().AddDate(0, 1, -1)},
				types.AuctionStatusStarted,
			),
			sdk.NewInt64Coin("denom2", 1_000_000_000_000),
		),
		types.NewBatchAuction(
			types.NewBaseAuction(
				2,
				types.AuctionTypeFixedPrice,
				sdk.AccAddress(crypto.AddressHash([]byte("Auctioneer"))).String(),
				types.SellingReserveAddress(1).String(),
				types.PayingReserveAddress(1).String(),
				math.LegacyMustNewDecFromStr("0.5"),
				sdk.NewInt64Coin("denom3", 1_000_000_000_000),
				"denom4",
				types.VestingReserveAddress(1).String(),
				[]types.VestingSchedule{},
				time.Now().AddDate(0, 0, -1),
				[]time.Time{time.Now().AddDate(0, 1, -1)},
				types.AuctionStatusStarted,
			),
			math.LegacyMustNewDecFromStr("0.1"),
			math.LegacyZeroDec(),
			uint32(3),
			math.LegacyMustNewDecFromStr("0.15"),
		),
	}

	a, err := types.PackAuction(auction[0])
	require.NoError(t, err)

	a2, err := types.PackAuction(auction[1])
	require.NoError(t, err)

	anyAuctions := []*codectypes.Any{a, a2}
	auctions, err := types.UnpackAuctions(anyAuctions)
	require.NoError(t, err)

	registry := codectypes.NewInterfaceRegistry()
	types.RegisterInterfaces(registry)
	cdc := codec.NewProtoCodec(registry)

	bz1 := types.MustMarshalAuction(cdc, auctions[0])
	auction1 := types.MustUnmarshalAuction(cdc, bz1)
	_, ok := auction1.(*types.FixedPriceAuction)
	require.True(t, ok)

	bz2 := types.MustMarshalAuction(cdc, auctions[1])
	auction2 := types.MustUnmarshalAuction(cdc, bz2)
	_, ok = auction2.(*types.BatchAuction)
	require.True(t, ok)
}

func TestShouldAuctionStarted(t *testing.T) {
	auction := types.BaseAuction{
		AuctionId:             1,
		Type:                  types.AuctionTypeFixedPrice,
		Auctioneer:            sdk.AccAddress(crypto.AddressHash([]byte("Auctioneer"))).String(),
		SellingReserveAddress: types.SellingReserveAddress(1).String(),
		PayingReserveAddress:  types.PayingReserveAddress(1).String(),
		StartPrice:            math.LegacyMustNewDecFromStr("0.5"),
		SellingCoin:           sdk.NewInt64Coin("denom3", 1_000_000_000_000),
		PayingCoinDenom:       "denom4",
		VestingReserveAddress: types.VestingReserveAddress(1).String(),
		VestingSchedules:      []types.VestingSchedule{},
		StartTime:             types.MustParseRFC3339("2021-12-01T00:00:00Z"),
		EndTime:               []time.Time{types.MustParseRFC3339("2021-12-15T00:00:00Z")},
		Status:                types.AuctionStatusStandBy,
	}

	for _, tc := range []struct {
		name string
		time string
		want bool
	}{
		{name: "Test case before auction start", time: "2021-11-01T00:00:00Z", want: false},
		{name: "Test case just before auction start", time: "2021-11-15T23:59:59Z", want: false},
		{name: "Test case well before auction start", time: "2021-11-20T00:00:00Z", want: false},
		{name: "Test case at auction start", time: "2021-12-01T00:00:00Z", want: true},
		{name: "Test case just after auction start", time: "2021-12-01T00:00:01Z", want: true},
		{name: "Test case during auction", time: "2021-12-10T00:00:00Z", want: true},
		{name: "Test case after auction start", time: "2022-01-01T00:00:00Z", want: true},
	} {
		t.Run(tc.name, func(t *testing.T) {
			require.Equal(t, tc.want, auction.ShouldAuctionStarted(types.MustParseRFC3339(tc.time)))
		})
	}
}

func TestShouldAuctionClosed(t *testing.T) {
	auction := types.BaseAuction{
		AuctionId:             1,
		Type:                  types.AuctionTypeFixedPrice,
		Auctioneer:            sdk.AccAddress(crypto.AddressHash([]byte("Auctioneer"))).String(),
		SellingReserveAddress: types.SellingReserveAddress(1).String(),
		PayingReserveAddress:  types.PayingReserveAddress(1).String(),
		StartPrice:            math.LegacyMustNewDecFromStr("0.5"),
		SellingCoin:           sdk.NewInt64Coin("denom3", 1_000_000_000_000),
		PayingCoinDenom:       "denom4",
		VestingReserveAddress: types.VestingReserveAddress(1).String(),
		VestingSchedules:      []types.VestingSchedule{},
		StartTime:             types.MustParseRFC3339("2021-12-01T00:00:00Z"),
		EndTime:               []time.Time{types.MustParseRFC3339("2021-12-15T00:00:00Z")},
		Status:                types.AuctionStatusStandBy,
	}

	for _, tc := range []struct {
		name string
		time string
		want bool
	}{
		{name: "Test case well before auction", time: "2021-11-01T00:00:00Z", want: false},
		{name: "Test case just before auction", time: "2021-11-15T23:59:59Z", want: false},
		{name: "Test case well before auction", time: "2021-11-20T00:00:00Z", want: false},
		{name: "Test case at auction end", time: "2021-12-15T00:00:00Z", want: true},
		{name: "Test case just after auction end", time: "2021-12-15T00:00:01Z", want: true},
		{name: "Test case after auction end", time: "2021-12-30T00:00:00Z", want: true},
		{name: "Test case well after auction end", time: "2022-01-01T00:00:00Z", want: true},
	} {
		require.Equal(t, tc.want, auction.ShouldAuctionClosed(types.MustParseRFC3339(tc.time)))
	}
}

func TestSellingReserveAddress(t *testing.T) {
	for _, tc := range []struct {
		name      string
		auctionID uint64
		want      string
	}{
		{name: "Test case for auctionID 1", auctionID: 1, want: "cosmos1wl90665mfk3pgg095qhmlgha934exjvv437acgq42zw0sg94flestth4zu"},
		{name: "Test case for auctionID 2", auctionID: 2, want: "cosmos197ewwasd96k2fh3nx5m76zvqxpzjcxuyq65rwgw0aa2edmwafgfqfa5qqz"},
	} {
		t.Run("", func(t *testing.T) {
			require.Equal(t, tc.want, types.SellingReserveAddress(tc.auctionID).String())
		})
	}
}

func TestPayingReserveAddress(t *testing.T) {
	for _, tc := range []struct {
		name      string
		auctionID uint64
		want      string
	}{
		{name: "Test case for auctionID 1", auctionID: 1, want: "cosmos17gk7a5ys8pxuexl7tvyk3pc9tdmqjjek03zjemez4eqvqdxlu92qdhphm2"},
		{name: "Test case for auctionID 2", auctionID: 2, want: "cosmos1s3cspws3lsqfvtjcz9jvpx7kjm93npmwjq8p4xfu3fcjj5jz9pks20uja6"},
	} {
		t.Run("", func(t *testing.T) {
			require.Equal(t, tc.want, types.PayingReserveAddress(tc.auctionID).String())
		})
	}
}

func TestVestingReserveAddress(t *testing.T) {
	for _, tc := range []struct {
		name      string
		auctionID uint64
		want      string
	}{
		{name: "Test case for auctionID 1", auctionID: 1, want: "cosmos1q4x4k4qsr4jwrrugnplhlj52mfd9f8jn5ck7r4ykdpv9wczvz4dqe8vrvt"},
		{name: "Test case for auctionID 2", auctionID: 2, want: "cosmos1pye9kv5f8s9n8uxnr0uznsn3klq57vqz8h2ya6u0v4w5666lqdfqjrw0qu"},
	} {
		t.Run("", func(t *testing.T) {
			require.Equal(t, tc.want, types.VestingReserveAddress(tc.auctionID).String())
		})
	}
}
