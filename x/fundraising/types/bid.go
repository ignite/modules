package types

import (
	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// NewBid returns a new Bid.
func NewBid(auctionID uint64, bidder sdk.AccAddress, bidID uint64, bidType BidType, price math.LegacyDec, coin sdk.Coin, isMatched bool) Bid {
	return Bid{
		AuctionId: auctionID,
		Bidder:    bidder.String(),
		BidId:     bidID,
		Type:      bidType,
		Price:     price,
		Coin:      coin,
		IsMatched: isMatched,
	}
}

func (b *Bid) SetMatched(status bool) {
	b.IsMatched = status
}

// ConvertToSellingAmount converts to selling amount depending on the bid coin denom.
// Note that we take as little coins as possible to prevent from overflowing the remaining selling coin.
func (b Bid) ConvertToSellingAmount(denom string) (amount math.Int) {
	if b.Coin.Denom == denom {
		return math.LegacyNewDecFromInt(b.Coin.Amount).QuoTruncate(b.Price).TruncateInt() // BidAmount / BidPrice
	}
	return b.Coin.Amount
}

// ConvertToPayingAmount converts to paying amount depending on the bid coin denom.
// Note that we take as many coins as possible by ceiling numbers from bidder.
func (b Bid) ConvertToPayingAmount(denom string) (amount math.Int) {
	if b.Coin.Denom == denom {
		return b.Coin.Amount
	}
	return math.LegacyNewDecFromInt(b.Coin.Amount).Mul(b.Price).Ceil().TruncateInt() // BidAmount * BidPrice
}
