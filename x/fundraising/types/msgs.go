package types

import (
	"time"

	sdkerrors "cosmossdk.io/errors"
	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
)

func NewMsgCancelAuction(auctioneer string, auctionID uint64) *MsgCancelAuction {
	return &MsgCancelAuction{
		Auctioneer: auctioneer,
		AuctionId:  auctionID,
	}
}

func (msg MsgCancelAuction) Type() string {
	return sdk.MsgTypeURL(&MsgCancelAuction{})
}

func NewMsgCreateBatchAuction(
	auctioneer string,
	startPrice math.LegacyDec,
	minBidPrice math.LegacyDec,
	sellingCoin sdk.Coin,
	payingCoinDenom string,
	vestingSchedules []VestingSchedule,
	maxExtendedRound uint32,
	extendedRoundRate math.LegacyDec,
	startTime time.Time,
	endTime time.Time,
) *MsgCreateBatchAuction {
	return &MsgCreateBatchAuction{
		Auctioneer:        auctioneer,
		StartPrice:        startPrice,
		MinBidPrice:       minBidPrice,
		SellingCoin:       sellingCoin,
		PayingCoinDenom:   payingCoinDenom,
		VestingSchedules:  vestingSchedules,
		MaxExtendedRound:  maxExtendedRound,
		ExtendedRoundRate: extendedRoundRate,
		StartTime:         startTime,
		EndTime:           endTime,
	}
}

func (msg MsgCreateBatchAuction) Type() string {
	return sdk.MsgTypeURL(&MsgCreateBatchAuction{})
}

func (msg MsgCreateBatchAuction) Validate() error {
	if !msg.StartPrice.IsPositive() {
		return sdkerrors.Wrapf(errors.ErrInvalidRequest, "start price must be positive")
	}
	if !msg.MinBidPrice.IsPositive() {
		return sdkerrors.Wrapf(errors.ErrInvalidRequest, "minimum price must be positive")
	}
	if err := msg.SellingCoin.Validate(); err != nil {
		return sdkerrors.Wrapf(errors.ErrInvalidRequest, "invalid selling coin: %v", err)
	}
	if !msg.SellingCoin.Amount.IsPositive() {
		return sdkerrors.Wrapf(errors.ErrInvalidRequest, "selling coin amount must be positive")
	}
	if msg.SellingCoin.Denom == msg.PayingCoinDenom {
		return sdkerrors.Wrapf(errors.ErrInvalidRequest, "selling coin denom must not be the same as paying coin denom")
	}
	if err := sdk.ValidateDenom(msg.PayingCoinDenom); err != nil {
		return sdkerrors.Wrapf(errors.ErrInvalidRequest, "invalid paying coin denom: %v", err)
	}
	if !msg.EndTime.After(msg.StartTime) {
		return sdkerrors.Wrapf(errors.ErrInvalidRequest, "end time must be set after start time")
	}
	if !msg.ExtendedRoundRate.IsPositive() {
		return sdkerrors.Wrapf(errors.ErrInvalidRequest, "extend rate must be positive")
	}
	return ValidateVestingSchedules(msg.VestingSchedules, msg.EndTime)
}

func NewMsgCreateFixedPriceAuction(
	auctioneer string,
	startPrice math.LegacyDec,
	sellingCoin sdk.Coin,
	payingCoinDenom string,
	vestingSchedules []VestingSchedule,
	startTime time.Time,
	endTime time.Time,
) *MsgCreateFixedPriceAuction {
	return &MsgCreateFixedPriceAuction{
		Auctioneer:       auctioneer,
		StartPrice:       startPrice,
		SellingCoin:      sellingCoin,
		PayingCoinDenom:  payingCoinDenom,
		VestingSchedules: vestingSchedules,
		StartTime:        startTime,
		EndTime:          endTime,
	}
}

func (msg MsgCreateFixedPriceAuction) Type() string {
	return sdk.MsgTypeURL(&MsgCreateFixedPriceAuction{})
}

func (msg MsgCreateFixedPriceAuction) Validate() error {
	if !msg.StartPrice.IsPositive() {
		return sdkerrors.Wrapf(errors.ErrInvalidRequest, "start price must be positive")
	}
	if err := msg.SellingCoin.Validate(); err != nil {
		return sdkerrors.Wrapf(errors.ErrInvalidRequest, "invalid selling coin: %v", err)
	}
	if !msg.SellingCoin.Amount.IsPositive() {
		return sdkerrors.Wrapf(errors.ErrInvalidRequest, "selling coin amount must be positive")
	}
	if msg.SellingCoin.Denom == msg.PayingCoinDenom {
		return sdkerrors.Wrapf(errors.ErrInvalidRequest, "selling coin denom must not be the same as paying coin denom")
	}
	if err := sdk.ValidateDenom(msg.PayingCoinDenom); err != nil {
		return sdkerrors.Wrapf(errors.ErrInvalidRequest, "invalid paying coin denom: %v", err)
	}
	if !msg.EndTime.After(msg.StartTime) {
		return sdkerrors.Wrapf(errors.ErrInvalidRequest, "end time must be set after start time")
	}
	return ValidateVestingSchedules(msg.VestingSchedules, msg.EndTime)
}

func NewMsgPlaceBid(
	auctionID uint64,
	bidder string,
	bidType BidType,
	price math.LegacyDec,
	coin sdk.Coin,
) *MsgPlaceBid {
	return &MsgPlaceBid{
		Bidder:    bidder,
		AuctionId: auctionID,
		BidType:   bidType,
		Price:     price,
		Coin:      coin,
	}
}

func (msg MsgPlaceBid) Type() string {
	return sdk.MsgTypeURL(&MsgPlaceBid{})
}

func (msg MsgPlaceBid) Validate() error {
	if !msg.Price.IsPositive() {
		return sdkerrors.Wrapf(errors.ErrInvalidRequest, "bid price must be positive value")
	}
	if err := msg.Coin.Validate(); err != nil {
		return sdkerrors.Wrapf(errors.ErrInvalidRequest, "invalid bid coin: %v", err)
	}
	if !msg.Coin.Amount.IsPositive() {
		return sdkerrors.Wrapf(errors.ErrInvalidRequest, "invalid coin amount: %s", msg.Coin.Amount.String())
	}
	if msg.BidType != BidTypeFixedPrice && msg.BidType != BidTypeBatchWorth &&
		msg.BidType != BidTypeBatchMany {
		return sdkerrors.Wrapf(errors.ErrInvalidRequest, "invalid bid type: %T", msg.BidType.String())
	}
	return nil
}

func NewMsgAddAllowedBidder(auctionID uint64, bidder string, maxBidAmount math.Int) *MsgAddAllowedBidder {
	return &MsgAddAllowedBidder{
		AuctionId:    auctionID,
		Bidder:       bidder,
		MaxBidAmount: maxBidAmount,
	}
}

func (msg MsgAddAllowedBidder) Type() string {
	return sdk.MsgTypeURL(&MsgAddAllowedBidder{})
}

func NewMsgModifyBid(
	auctionID uint64,
	bidder string,
	bidID uint64,
	price math.LegacyDec,
	coin sdk.Coin,
) *MsgModifyBid {
	return &MsgModifyBid{
		Bidder:    bidder,
		AuctionId: auctionID,
		BidId:     bidID,
		Price:     price,
		Coin:      coin,
	}
}

func (msg MsgModifyBid) Type() string {
	return sdk.MsgTypeURL(&MsgModifyBid{})
}

func (msg MsgModifyBid) Validate() error {
	if !msg.Price.IsPositive() {
		return sdkerrors.Wrapf(errors.ErrInvalidRequest, "bid price must be positive value")
	}
	if err := msg.Coin.Validate(); err != nil {
		return sdkerrors.Wrapf(errors.ErrInvalidRequest, "invalid bid coin: %v", err)
	}
	if !msg.Coin.Amount.IsPositive() {
		return sdkerrors.Wrapf(errors.ErrInvalidRequest, "invalid coin amount: %s", msg.Coin.Amount.String())
	}
	return nil
}
