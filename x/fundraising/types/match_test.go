package types_test

import (
	"encoding/binary"
	"testing"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/ignite/modules/x/fundraising/types"
)

// These helper functions below are taken from keeper_test package
// move these to separate file

func testAddr(addrNum int) sdk.AccAddress {
	addr := make(sdk.AccAddress, 20)
	binary.PutVarint(addr, int64(addrNum))
	return addr
}

// parseDec parses string and returns math.LegacyDec.
func parseDec(s string) math.LegacyDec {
	return math.LegacyMustNewDecFromStr(s)
}

func TestMatch(t *testing.T) {
	const (
		payingCoinDenom  = "paying"
		sellingCoinDenom = "selling"
	)

	newBid := func(id uint64, typ types.BidType, bidder string, price math.LegacyDec, bidAmt math.Int) types.Bid {
		var coin sdk.Coin
		switch typ {
		case types.BidTypeBatchWorth:
			coin = sdk.NewCoin(payingCoinDenom, price.MulInt(bidAmt).Ceil().TruncateInt())
		case types.BidTypeBatchMany:
			coin = sdk.NewCoin(sellingCoinDenom, bidAmt)
		}
		return types.Bid{
			// Omitted fields are not important when testing types.Match
			BidId:     id,
			Bidder:    bidder,
			Type:      typ,
			Price:     price,
			Coin:      coin,
			IsMatched: false,
		}
	}

	var bidders []string
	for i := 0; i < 10; i++ {
		bidders = append(bidders, testAddr(i).String())
	}

	for _, tc := range []struct {
		name                string
		allowedBidders      map[string]math.Int
		sellingCoinAmt      math.Int
		bids                []types.Bid
		matchPrice          math.LegacyDec
		matched             bool
		matchedAmt          math.Int
		matchedBidIDs       []uint64 // should be sorted
		matchResultByBidder map[string]*types.BidderMatchResult
	}{
		{
			"basic case",
			map[string]math.Int{
				bidders[0]: math.NewInt(100_000000),
			},
			math.NewInt(100_000000),
			[]types.Bid{
				newBid(1, types.BidTypeBatchWorth, bidders[0], parseDec("1.0"), math.NewInt(100_000000)),
			},
			parseDec("1.0"),
			true,
			math.NewInt(100_000000),
			[]uint64{1},
			map[string]*types.BidderMatchResult{
				bidders[0]: {
					PayingAmount:  math.NewInt(100_000000),
					MatchedAmount: math.NewInt(100_000000),
				},
			},
		},
		{
			"partial match",
			map[string]math.Int{
				bidders[0]: math.NewInt(50_000000),
			},
			math.NewInt(100_000000),
			[]types.Bid{
				newBid(1, types.BidTypeBatchWorth, bidders[0], parseDec("1.0"), math.NewInt(100_000000)),
			},
			parseDec("1.0"),
			true,
			math.NewInt(50_000000),
			[]uint64{1},
			map[string]*types.BidderMatchResult{
				bidders[0]: {
					PayingAmount:  math.NewInt(50_000000),
					MatchedAmount: math.NewInt(50_000000),
				},
			},
		},
		{
			"no match",
			map[string]math.Int{
				bidders[0]: math.NewInt(100_000000),
			},
			math.NewInt(100_000000),
			[]types.Bid{
				newBid(1, types.BidTypeBatchWorth, bidders[0], parseDec("1.0"), math.NewInt(100_000000)),
			},
			parseDec("1.1"),
			false,
			math.Int{},
			nil, nil,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			var allowedBidders []types.AllowedBidder
			for bidder, maxBidAmt := range tc.allowedBidders {
				allowedBidders = append(allowedBidders, types.AllowedBidder{
					Bidder:       bidder,
					MaxBidAmount: maxBidAmt,
				})
			}
			prices, bidsByPrice := types.BidsByPrice(tc.bids)
			matchRes, matched := types.Match(tc.matchPrice, prices, bidsByPrice, tc.sellingCoinAmt, allowedBidders)
			require.Equal(t, tc.matched, matched)
			if matched {
				require.True(math.IntEq(t, tc.matchedAmt, matchRes.MatchedAmount))
				var matchedBidIDs []uint64
				for _, bid := range matchRes.MatchedBids {
					matchedBidIDs = append(matchedBidIDs, bid.BidId)
				}
				require.Equal(t, tc.matchedBidIDs, matchedBidIDs)
				require.Equal(t, tc.matchResultByBidder, matchRes.MatchResultByBidder)
			}
		})
	}
}
