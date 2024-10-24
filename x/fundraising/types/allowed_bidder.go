package types

import (
	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// NewAllowedBidder returns a new AllowedBidder.
func NewAllowedBidder(auctionID uint64, bidderAddr sdk.AccAddress, maxBidAmount math.Int) AllowedBidder {
	return AllowedBidder{
		AuctionId:    auctionID,
		Bidder:       bidderAddr.String(),
		MaxBidAmount: maxBidAmount,
	}
}

// Validate validates allowed bidder object.
func (ab AllowedBidder) Validate() error {
	if ab.MaxBidAmount.IsNil() {
		return ErrInvalidMaxBidAmount
	}
	if !ab.MaxBidAmount.IsPositive() {
		return ErrInvalidMaxBidAmount
	}
	return nil
}
