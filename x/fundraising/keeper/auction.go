package keeper

import (
	"context"
	"fmt"
	"sort"
	"strconv"
	"time"

	"cosmossdk.io/collections"
	sdkerrors "cosmossdk.io/errors"
	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"

	"github.com/ignite/modules/x/fundraising/types"
)

type inOutCoins struct {
	bidder  string
	input   banktypes.Input
	outputs []banktypes.Output
}

// GetAuction returns Auction interface by auction ID.
func (k Keeper) GetAuction(ctx context.Context, auctionID uint64) (types.AuctionI, error) {
	a, err := k.Auction.Get(ctx, auctionID)
	if sdkerrors.IsOf(err, collections.ErrNotFound) {
		return nil, types.ErrAuctionNotFound
	}
	return a, err
}

// Auctions returns all Actions.
func (k Keeper) Auctions(ctx context.Context) ([]types.AuctionI, error) {
	auctions := make([]types.AuctionI, 0)
	err := k.Auction.Walk(ctx, nil, func(_ uint64, auction types.AuctionI) (bool, error) {
		auctions = append(auctions, auction)
		return false, nil
	})
	return auctions, err
}

// AddAllowedBidders is a function that is implemented for an external module.
// An external module uses this function to add allowed bidders in the auction's allowed bidders list.
// It doesn't look up the bidder's previous maximum bid amount. Instead, it overlaps.
// It doesn't have any auctioneer's verification logic because the module is fundamentally designed
// to delegate full authorization to an external module.
// It is up to an external module to freely add necessary verification and operations depending on their use cases.
func (k Keeper) AddAllowedBidders(ctx context.Context, auctionID uint64, allowedBidders []types.AllowedBidder) error {
	if len(allowedBidders) == 0 {
		return types.ErrEmptyAllowedBidders
	}

	auction, err := k.Auction.Get(ctx, auctionID)
	if err != nil {
		return sdkerrors.Wrapf(err, "auction %d is not found", auctionID)
	}

	// Call hook before adding allowed bidders for the auction
	if err := k.BeforeAllowedBiddersAdded(ctx, allowedBidders); err != nil {
		return err
	}

	// Store new allowed bidders
	for _, ab := range allowedBidders {
		if err := ab.Validate(); err != nil {
			return err
		}
		if ab.MaxBidAmount.GT(auction.GetSellingCoin().Amount) {
			return types.ErrInsufficientRemainingAmount
		}

		bidder, err := k.addressCodec.StringToBytes(ab.Bidder)
		if err != nil {
			return sdkerrors.Wrap(err, "invalid address")
		}
		if err := k.AllowedBidder.Set(ctx, collections.Join(auctionID, sdk.AccAddress(bidder)), ab); err != nil {
			return err
		}
	}

	return nil
}

// UpdateAllowedBidder is a function that is implemented for an external module.
// An external module uses this function to update maximum bid amount of particular allowed bidder in the auction.
// It doesn't have any auctioneer's verification logic because the module is fundamentally designed
// to delegate full authorization to an external module.
// It is up to an external module to freely add necessary verification and operations depending on their use cases.
func (k Keeper) UpdateAllowedBidder(ctx context.Context, auctionID uint64, bidder sdk.AccAddress, maxBidAmount math.Int) error {
	_, err := k.Auction.Get(ctx, auctionID)
	if err != nil {
		return sdkerrors.Wrapf(err, "auction %d is not found", auctionID)
	}

	_, err = k.AllowedBidder.Get(ctx, collections.Join(auctionID, bidder))
	if err != nil {
		return sdkerrors.Wrapf(errors.ErrNotFound, "bidder %s is not found", bidder.String())
	}

	allowedBidder := types.NewAllowedBidder(auctionID, bidder, maxBidAmount)

	if err := allowedBidder.Validate(); err != nil {
		return err
	}

	// Call hook before updating the allowed bidders for the auction
	if err := k.BeforeAllowedBidderUpdated(ctx, auctionID, bidder, maxBidAmount); err != nil {
		return err
	}

	if err := k.AllowedBidder.Set(ctx, collections.Join(auctionID, bidder), allowedBidder); err != nil {
		return sdkerrors.Wrapf(errors.ErrNotFound, "allowed bidder %s no set", bidder.String())
	}

	return nil
}

// AllocateSellingCoin allocates allocated selling coin for all matched bids in MatchingInfo and
// releases them from the selling reserve account.
func (k Keeper) AllocateSellingCoin(ctx context.Context, auction types.AuctionI, mInfo MatchingInfo) error {
	// Call hook before selling coin allocation
	if err := k.BeforeSellingCoinsAllocated(ctx, auction.GetId(), mInfo.AllocationMap, mInfo.RefundMap); err != nil {
		return err
	}

	sellingReserveAddr, err := k.addressCodec.StringToBytes(auction.GetSellingReserveAddress())
	if err != nil {
		return sdkerrors.Wrap(err, "invalid address")
	}
	sellingCoinDenom := auction.GetSellingCoin().Denom

	ioCoins := make(map[string]inOutCoins)

	// Sort bidders to reserve determinism
	var bidders []string
	for bidder := range mInfo.AllocationMap {
		bidders = append(bidders, bidder)
	}
	sort.Strings(bidders)

	// Allocate coins to all matched bidders in AllocationMap and
	// set the amounts in transaction inputs and outputs from the selling reserve account
	for _, bidder := range bidders {
		if mInfo.AllocationMap[bidder].IsZero() {
			continue
		}
		allocateCoins := sdk.NewCoins(sdk.NewCoin(sellingCoinDenom, mInfo.AllocationMap[bidder]))
		bidderAddr, err := k.addressCodec.StringToBytes(bidder)
		if err != nil {
			return sdkerrors.Wrap(err, "invalid address")
		}

		if _, ok := ioCoins[bidder]; !ok {
			ioCoins[bidder] = inOutCoins{
				bidder:  bidder,
				outputs: []banktypes.Output{banktypes.NewOutput(bidderAddr, allocateCoins)},
				input:   banktypes.NewInput(sellingReserveAddr, allocateCoins),
			}
			continue
		}

		inout := ioCoins[bidder]
		inout.input.Coins = inout.input.Coins.Add(allocateCoins...)
		inout.outputs = append(inout.outputs, banktypes.NewOutput(bidderAddr, allocateCoins))
		ioCoins[bidder] = inout
	}

	// Send all inputs
	for _, inout := range ioCoins {
		if err := k.bankKeeper.InputOutputCoins(ctx, inout.input, inout.outputs); err != nil {
			return err
		}
	}

	return nil
}

// ReleaseVestingPayingCoin releases the vested selling coin to the auctioneer from the vesting reserve account.
func (k Keeper) ReleaseVestingPayingCoin(ctx context.Context, auction types.AuctionI) error {
	vestingQueues, err := k.GetVestingQueuesByAuctionID(ctx, auction.GetId())
	if err != nil {
		return err
	}
	vestingQueuesLen := len(vestingQueues)

	for i, vestingQueue := range vestingQueues {
		blockTime := sdk.UnwrapSDKContext(ctx).BlockTime()
		if vestingQueue.ShouldRelease(blockTime) {
			vestingReserve, err := k.addressCodec.StringToBytes(auction.GetVestingReserveAddress())
			if err != nil {
				return sdkerrors.Wrap(err, "invalid address")
			}

			auctioneer, err := k.addressCodec.StringToBytes(auction.GetAuctioneer())
			if err != nil {
				return sdkerrors.Wrap(err, "invalid address")
			}

			payingCoins := sdk.NewCoins(vestingQueue.PayingCoin)

			if err := k.bankKeeper.SendCoins(ctx, vestingReserve, auctioneer, payingCoins); err != nil {
				return sdkerrors.Wrap(err, "failed to release paying coin to the auctioneer")
			}

			vestingQueue.SetReleased(true)
			if err := k.VestingQueue.Set(ctx, collections.Join(
				vestingQueue.AuctionId,
				vestingQueue.ReleaseTime,
			), vestingQueue); err != nil {
				return err
			}

			// Update status when all the amounts are released
			if i == vestingQueuesLen-1 {
				if err := auction.SetStatus(types.AuctionStatusFinished); err != nil {
					return err
				}
				if err := k.Auction.Set(ctx, auction.GetId(), auction); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

// RefundRemainingSellingCoin refunds the remaining selling coin to the auctioneer.
func (k Keeper) RefundRemainingSellingCoin(ctx context.Context, auction types.AuctionI) error {
	sellingReserveAddr, err := k.addressCodec.StringToBytes(auction.GetSellingReserveAddress())
	if err != nil {
		return sdkerrors.Wrap(err, "invalid address")
	}
	sellingCoinDenom := auction.GetSellingCoin().Denom
	spendableCoins := k.bankKeeper.SpendableCoins(ctx, sellingReserveAddr)
	releaseCoins := sdk.NewCoins(sdk.NewCoin(sellingCoinDenom, spendableCoins.AmountOf(sellingCoinDenom)))
	auctioneer, err := k.addressCodec.StringToBytes(auction.GetAuctioneer())
	if err != nil {
		return sdkerrors.Wrap(err, "invalid address")
	}

	if err := k.bankKeeper.SendCoins(ctx, sellingReserveAddr, auctioneer, releaseCoins); err != nil {
		return err
	}
	return nil
}

// RefundPayingCoin refunds paying coin to the corresponding bidders.
func (k Keeper) RefundPayingCoin(ctx context.Context, auction types.AuctionI, mInfo MatchingInfo) error {
	payingReserveAddr, err := k.addressCodec.StringToBytes(auction.GetPayingReserveAddress())
	if err != nil {
		return sdkerrors.Wrap(err, "invalid address")
	}
	payingCoinDenom := auction.GetPayingCoinDenom()

	ioCoins := make(map[string]inOutCoins)

	// Sort bidders to reserve determinism
	var bidders []string
	for bidder := range mInfo.RefundMap {
		bidders = append(bidders, bidder)
	}
	sort.Strings(bidders)

	// Refund the unmatched bid amount back to the bidder
	for _, bidder := range bidders {
		if mInfo.RefundMap[bidder].IsZero() {
			continue
		}

		bidderAddr, err := k.addressCodec.StringToBytes(bidder)
		if err != nil {
			return sdkerrors.Wrap(err, "invalid address")
		}
		refundCoins := sdk.NewCoins(sdk.NewCoin(payingCoinDenom, mInfo.RefundMap[bidder]))

		if _, ok := ioCoins[bidder]; !ok {
			ioCoins[bidder] = inOutCoins{
				bidder:  bidder,
				outputs: []banktypes.Output{banktypes.NewOutput(bidderAddr, refundCoins)},
				input:   banktypes.NewInput(payingReserveAddr, refundCoins),
			}
			continue
		}

		inout := ioCoins[bidder]
		inout.input.Coins = inout.input.Coins.Add(refundCoins...)
		inout.outputs = append(inout.outputs, banktypes.NewOutput(bidderAddr, refundCoins))
		ioCoins[bidder] = inout
	}

	// Send all inputs.
	for _, inout := range ioCoins {
		if err := k.bankKeeper.InputOutputCoins(ctx, inout.input, inout.outputs); err != nil {
			return err
		}
	}

	return nil
}

// ExtendRound extends another round of ExtendedPeriod value for the auction.
func (k Keeper) ExtendRound(ctx context.Context, ba *types.BatchAuction) error {
	params, err := k.Params.Get(ctx)
	if err != nil {
		return err
	}
	extendedPeriod := params.ExtendedPeriod
	nextEndTime := ba.GetEndTime()[len(ba.GetEndTime())-1].AddDate(0, 0, int(extendedPeriod))
	endTime := append(ba.GetEndTime(), nextEndTime)

	_ = ba.SetEndTime(endTime)

	return k.Auction.Set(ctx, ba.GetId(), ba)
}

// CloseFixedPriceAuction closes a fixed price auction.
func (k Keeper) CloseFixedPriceAuction(ctx context.Context, auction types.AuctionI) error {
	mInfo, err := k.CalculateFixedPriceAllocation(ctx, auction)
	if err != nil {
		return err
	}

	if err := k.AllocateSellingCoin(ctx, auction, mInfo); err != nil {
		return err
	}

	if err := k.RefundRemainingSellingCoin(ctx, auction); err != nil {
		return err
	}

	if err := k.ApplyVestingSchedules(ctx, auction); err != nil {
		return err
	}

	return nil
}

// CloseBatchAuction closes a batch auction.
func (k Keeper) CloseBatchAuction(ctx context.Context, auction types.AuctionI) error {
	ba, ok := auction.(*types.BatchAuction)
	if !ok {
		return fmt.Errorf("unable to close auction that is not a batch auction: %T", auction)
	}

	// Extend round since there is no last matched length to compare with
	lastMatchedLen, err := k.GetLastMatchedBidsLen(ctx, ba.GetId())
	if err != nil {
		return err
	}
	mInfo, err := k.CalculateBatchAllocation(ctx, auction)
	if err != nil {
		return err
	}

	// Close the auction when maximum extended round + 1 is the same as the length of end times
	// If the value of MaxExtendedRound is 0, it means that an auctioneer does not want have an extended round
	if ba.MaxExtendedRound+1 == uint32(len(auction.GetEndTime())) {
		if err := k.AllocateSellingCoin(ctx, auction, mInfo); err != nil {
			return err
		}

		if err := k.RefundRemainingSellingCoin(ctx, auction); err != nil {
			return err
		}

		if err := k.RefundPayingCoin(ctx, auction, mInfo); err != nil {
			return err
		}

		if err := k.ApplyVestingSchedules(ctx, auction); err != nil {
			return err
		}

		return nil
	}

	if lastMatchedLen == 0 {
		return k.ExtendRound(ctx, ba)
	}

	currDec := math.LegacyNewDec(mInfo.MatchedLen)
	lastDec := math.LegacyNewDec(lastMatchedLen)
	diff := math.LegacyOneDec().Sub(currDec.Quo(lastDec)) // 1 - (CurrentMatchedLenDec / LastMatchedLenDec)

	// To prevent from auction sniping technique, compare the extended round rate with
	// the current and the last length of matched bids to determine
	// if the auction needs another extended round
	if diff.GTE(ba.ExtendedRoundRate) {
		return k.ExtendRound(ctx, ba)
	}

	if err := k.AllocateSellingCoin(ctx, auction, mInfo); err != nil {
		return err
	}

	if err := k.RefundRemainingSellingCoin(ctx, auction); err != nil {
		return err
	}

	if err := k.RefundPayingCoin(ctx, auction, mInfo); err != nil {
		return err
	}

	if err := k.ApplyVestingSchedules(ctx, auction); err != nil {
		return err
	}

	return nil
}

// CreateFixedPriceAuction handles types.MsgCreateFixedPriceAuction and create a fixed price auction.
// Note that the module is designed to delegate authorization to an external module to add allowed bidders for the auction.
func (k Keeper) CreateFixedPriceAuction(ctx context.Context, msg *types.MsgCreateFixedPriceAuction) (types.AuctionI, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	blockTime := sdkCtx.BlockTime()
	if blockTime.After(msg.EndTime) { // EndTime < CurrentTime
		return nil, sdkerrors.Wrapf(errors.ErrInvalidRequest, "end time (%d) must be set after the current time", msg.EndTime.Unix())
	}

	if len(msg.VestingSchedules) > types.MaxNumVestingSchedules {
		return nil, sdkerrors.Wrap(errors.ErrInvalidRequest, "exceed maximum number of vesting schedules")
	}

	nextId, err := k.AuctionSeq.Next(ctx)
	if err != nil {
		return nil, err
	}

	auctioneer, err := k.addressCodec.StringToBytes(msg.GetAuctioneer())
	if err != nil {
		return nil, sdkerrors.Wrap(err, "invalid address")
	}

	if err := k.PayCreationFee(ctx, auctioneer); err != nil {
		return nil, sdkerrors.Wrap(err, "failed to pay auction creation fee")
	}

	if err := k.ReserveSellingCoin(ctx, nextId, auctioneer, msg.SellingCoin); err != nil {
		return nil, sdkerrors.Wrap(err, "failed to reserve selling coin")
	}

	ba := types.NewBaseAuction(
		nextId,
		types.AuctionTypeFixedPrice,
		msg.Auctioneer,
		types.SellingReserveAddress(nextId).String(),
		types.PayingReserveAddress(nextId).String(),
		msg.StartPrice,
		msg.SellingCoin,
		msg.PayingCoinDenom,
		types.VestingReserveAddress(nextId).String(),
		msg.VestingSchedules,
		msg.StartTime,
		[]time.Time{msg.EndTime}, // it is an array data type to handle BatchAuction
		types.AuctionStatusStandBy,
	)

	// Update status if the start time is already passed over the current time
	if ba.ShouldAuctionStarted(blockTime) {
		_ = ba.SetStatus(types.AuctionStatusStarted)
	}

	auction := types.NewFixedPriceAuction(ba, msg.SellingCoin)

	// Call hook before storing an auction
	if err := k.BeforeFixedPriceAuctionCreated(
		ctx,
		auction.Auctioneer,
		auction.StartPrice,
		auction.SellingCoin,
		auction.PayingCoinDenom,
		auction.VestingSchedules,
		auction.StartTime,
		auction.EndTime[0],
	); err != nil {
		return nil, err
	}

	if err := k.Auction.Set(ctx, ba.GetId(), auction); err != nil {
		return nil, err
	}

	// Call hook after storing an auction
	if err := k.AfterFixedPriceAuctionCreated(
		ctx,
		auction.AuctionId,
		auction.Auctioneer,
		auction.StartPrice,
		auction.SellingCoin,
		auction.PayingCoinDenom,
		auction.VestingSchedules,
		auction.StartTime,
		auction.EndTime[0],
	); err != nil {
		return nil, err
	}

	sdkCtx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeCreateFixedPriceAuction,
			sdk.NewAttribute(types.AttributeKeyAuctionID, strconv.FormatUint(nextId, 10)),
			sdk.NewAttribute(types.AttributeKeyAuctioneerAddress, auction.GetAuctioneer()),
			sdk.NewAttribute(types.AttributeKeySellingReserveAddress, auction.GetSellingReserveAddress()),
			sdk.NewAttribute(types.AttributeKeyPayingReserveAddress, auction.GetPayingReserveAddress()),
			sdk.NewAttribute(types.AttributeKeyStartPrice, auction.GetStartPrice().String()),
			sdk.NewAttribute(types.AttributeKeySellingCoin, auction.GetSellingCoin().String()),
			sdk.NewAttribute(types.AttributeKeyPayingCoinDenom, auction.GetPayingCoinDenom()),
			sdk.NewAttribute(types.AttributeKeyVestingReserveAddress, auction.GetVestingReserveAddress()),
			sdk.NewAttribute(types.AttributeKeyRemainingSellingCoin, auction.RemainingSellingCoin.String()),
			sdk.NewAttribute(types.AttributeKeyStartTime, auction.GetStartTime().String()),
			sdk.NewAttribute(types.AttributeKeyEndTime, msg.EndTime.String()),
			sdk.NewAttribute(types.AttributeKeyAuctionStatus, auction.GetStatus().String()),
		),
	})

	return auction, nil
}

// CreateBatchAuction handles types.MsgCreateBatchAuction and create a batch auction.
// Note that the module is designed to delegate authorization to an external module to add allowed bidders for the auction.
func (k Keeper) CreateBatchAuction(ctx context.Context, msg *types.MsgCreateBatchAuction) (types.AuctionI, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	blockTime := sdkCtx.BlockTime()
	if blockTime.After(msg.EndTime) { // EndTime < CurrentTime
		return nil, sdkerrors.Wrapf(errors.ErrInvalidRequest, "end time (%d) must be set after the current time", msg.EndTime.Unix())
	}

	if len(msg.VestingSchedules) > types.MaxNumVestingSchedules {
		return nil, sdkerrors.Wrap(errors.ErrInvalidRequest, "exceed maximum number of vesting schedules")
	}

	if msg.MaxExtendedRound > types.MaxExtendedRound {
		return nil, sdkerrors.Wrap(errors.ErrInvalidRequest, "exceed maximum extended round")
	}

	nextId, err := k.AuctionSeq.Next(ctx)
	if err != nil {
		return nil, err
	}

	auctioneer, err := k.addressCodec.StringToBytes(msg.GetAuctioneer())
	if err != nil {
		return nil, sdkerrors.Wrap(err, "invalid address")
	}

	if err := k.PayCreationFee(ctx, auctioneer); err != nil {
		return nil, sdkerrors.Wrap(err, "failed to pay auction creation fee")
	}

	if err := k.ReserveSellingCoin(ctx, nextId, auctioneer, msg.SellingCoin); err != nil {
		return nil, sdkerrors.Wrap(err, "failed to reserve selling coin")
	}

	endTime := []time.Time{msg.EndTime} // it is an array data type to handle BatchAuction
	ba := types.NewBaseAuction(
		nextId,
		types.AuctionTypeBatch,
		msg.Auctioneer,
		types.SellingReserveAddress(nextId).String(),
		types.PayingReserveAddress(nextId).String(),
		msg.StartPrice,
		msg.SellingCoin,
		msg.PayingCoinDenom,
		types.VestingReserveAddress(nextId).String(),
		msg.VestingSchedules,
		msg.StartTime,
		endTime,
		types.AuctionStatusStandBy,
	)

	// Update status if the start time is already passed the current time
	if ba.ShouldAuctionStarted(blockTime) {
		_ = ba.SetStatus(types.AuctionStatusStarted)
	}

	auction := types.NewBatchAuction(
		ba,
		msg.MinBidPrice,
		math.LegacyZeroDec(),
		msg.MaxExtendedRound,
		msg.ExtendedRoundRate,
	)

	// Call hook before storing an auction
	if err := k.BeforeBatchAuctionCreated(
		ctx,
		auction.Auctioneer,
		auction.StartPrice,
		auction.MinBidPrice,
		auction.SellingCoin,
		auction.PayingCoinDenom,
		auction.VestingSchedules,
		auction.MaxExtendedRound,
		auction.ExtendedRoundRate,
		auction.StartTime,
		auction.EndTime[0],
	); err != nil {
		return nil, err
	}

	if err := k.Auction.Set(ctx, ba.GetId(), auction); err != nil {
		return nil, err
	}

	// Call hook after storing an auction
	if err := k.AfterBatchAuctionCreated(
		ctx,
		auction.AuctionId,
		auction.Auctioneer,
		auction.StartPrice,
		auction.MinBidPrice,
		auction.SellingCoin,
		auction.PayingCoinDenom,
		auction.VestingSchedules,
		auction.MaxExtendedRound,
		auction.ExtendedRoundRate,
		auction.StartTime,
		auction.EndTime[0],
	); err != nil {
		return nil, err
	}

	sdkCtx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeCreateBatchAuction,
			sdk.NewAttribute(types.AttributeKeyAuctionID, strconv.FormatUint(nextId, 10)),
			sdk.NewAttribute(types.AttributeKeyAuctioneerAddress, auction.GetAuctioneer()),
			sdk.NewAttribute(types.AttributeKeySellingReserveAddress, auction.GetSellingReserveAddress()),
			sdk.NewAttribute(types.AttributeKeyPayingReserveAddress, auction.GetPayingReserveAddress()),
			sdk.NewAttribute(types.AttributeKeyStartPrice, auction.GetStartPrice().String()),
			sdk.NewAttribute(types.AttributeKeySellingCoin, auction.GetSellingCoin().String()),
			sdk.NewAttribute(types.AttributeKeyPayingCoinDenom, auction.GetPayingCoinDenom()),
			sdk.NewAttribute(types.AttributeKeyVestingReserveAddress, auction.GetVestingReserveAddress()),
			sdk.NewAttribute(types.AttributeKeyStartTime, auction.GetStartTime().String()),
			sdk.NewAttribute(types.AttributeKeyEndTime, msg.EndTime.String()),
			sdk.NewAttribute(types.AttributeKeyAuctionStatus, auction.GetStatus().String()),
			sdk.NewAttribute(types.AttributeKeyMinBidPrice, auction.MinBidPrice.String()),
			sdk.NewAttribute(types.AttributeKeyMaxExtendedRound, fmt.Sprint(auction.MaxExtendedRound)),
			sdk.NewAttribute(types.AttributeKeyExtendedRoundRate, auction.ExtendedRoundRate.String()),
		),
	})

	return auction, nil
}

// CancelAuction handles types.MsgCancelAuction and cancels the auction.
// An auction can only be canceled when it is not started yet.
func (k Keeper) CancelAuction(ctx context.Context, msg *types.MsgCancelAuction) error {
	auction, err := k.Auction.Get(ctx, msg.AuctionId)
	if err != nil {
		return err
	}

	if auction.GetAuctioneer() != msg.Auctioneer {
		return sdkerrors.Wrap(errors.ErrUnauthorized, "only the auctioneer can cancel the auction")
	}

	if auction.GetStatus() != types.AuctionStatusStandBy {
		return sdkerrors.Wrap(types.ErrInvalidAuctionStatus, "only the stand by auction can be cancelled")
	}

	sellingReserveAddr, err := k.addressCodec.StringToBytes(auction.GetSellingReserveAddress())
	if err != nil {
		return sdkerrors.Wrap(err, "invalid address")
	}
	sellingCoinDenom := auction.GetSellingCoin().Denom
	spendableCoins := k.bankKeeper.SpendableCoins(ctx, sellingReserveAddr)
	releaseCoin := sdk.NewCoin(sellingCoinDenom, spendableCoins.AmountOf(sellingCoinDenom))

	// Release the selling coin back to the auctioneer
	auctioneer, err := k.addressCodec.StringToBytes(msg.GetAuctioneer())
	if err != nil {
		return sdkerrors.Wrap(err, "invalid address")
	}
	if err := k.bankKeeper.SendCoins(ctx, sellingReserveAddr, auctioneer, sdk.NewCoins(releaseCoin)); err != nil {
		return sdkerrors.Wrap(err, "failed to release the selling coin")
	}

	// Call hook before cancelling the auction
	if err := k.BeforeAuctionCanceled(ctx, msg.AuctionId, msg.Auctioneer); err != nil {
		return err
	}

	if auction.GetType() == types.AuctionTypeFixedPrice {
		fa := auction.(*types.FixedPriceAuction)
		fa.RemainingSellingCoin = sdk.NewCoin(sellingCoinDenom, math.ZeroInt())
		auction = fa
	}

	_ = auction.SetStatus(types.AuctionStatusCancelled)
	if err := k.Auction.Set(ctx, auction.GetId(), auction); err != nil {
		return err
	}

	sdkCtx := sdk.UnwrapSDKContext(ctx)
	sdkCtx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeCancelAuction,
			sdk.NewAttribute(types.AttributeKeyAuctionID, strconv.FormatUint(auction.GetId(), 10)),
		),
	})

	return nil
}
