package keeper

import (
	"context"
	"time"

	"cosmossdk.io/collections"
	sdkerrors "cosmossdk.io/errors"
	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ignite/modules/x/fundraising/types"
)

// GetVestingQueue returns VestingQueue by auction ID and release time.
func (k Keeper) GetVestingQueue(ctx context.Context, auctionID uint64, releaseTime time.Time) (types.VestingQueue, error) {
	v, err := k.VestingQueue.Get(ctx, collections.Join(auctionID, releaseTime))
	if sdkerrors.IsOf(err, collections.ErrNotFound) {
		return types.VestingQueue{}, types.ErrVestingQueueNotFound
	}
	return v, err
}

// VestingQueues returns all VestingQueue.
func (k Keeper) VestingQueues(ctx context.Context) ([]types.VestingQueue, error) {
	vestingQueues := make([]types.VestingQueue, 0)
	err := k.VestingQueue.Walk(ctx, nil, func(_ collections.Pair[uint64, time.Time], bid types.VestingQueue) (bool, error) {
		vestingQueues = append(vestingQueues, bid)
		return false, nil
	})
	return vestingQueues, err
}

// GetVestingQueuesByAuctionID returns all vesting queues associated with the auction id that are registered in the store.
func (k Keeper) GetVestingQueuesByAuctionID(ctx context.Context, auctionID uint64) ([]types.VestingQueue, error) {
	vestingQueues := make([]types.VestingQueue, 0)
	rng := collections.NewPrefixedPairRange[uint64, time.Time](auctionID)
	err := k.VestingQueue.Walk(ctx, rng, func(key collections.Pair[uint64, time.Time], vestingQueue types.VestingQueue) (bool, error) {
		vestingQueues = append(vestingQueues, vestingQueue)
		return false, nil
	})
	return vestingQueues, err
}

// ApplyVestingSchedules stores vesting queues based on the vesting schedules of the auction and
// sets status to vesting.
func (k Keeper) ApplyVestingSchedules(ctx context.Context, auction types.AuctionI) error {
	payingReserveAddr := auction.GetPayingReserveAddress()
	vestingReserveAddr := auction.GetVestingReserveAddress()
	payingCoinDenom := auction.GetPayingCoinDenom()
	spendableCoins := k.bankKeeper.SpendableCoins(ctx, payingReserveAddr)
	reserveCoin := sdk.NewCoin(payingCoinDenom, spendableCoins.AmountOf(payingCoinDenom))

	vsLen := len(auction.GetVestingSchedules())
	if vsLen == 0 {
		// Send reserve coins to the auctioneer from the paying reserve account
		if err := k.bankKeeper.SendCoins(ctx, payingReserveAddr, auction.GetAuctioneer(), sdk.NewCoins(reserveCoin)); err != nil {
			return err
		}

		if err := auction.SetStatus(types.AuctionStatusFinished); err != nil {
			return err
		}
		if err := k.Auction.Set(ctx, auction.GetId(), auction); err != nil {
			return err
		}
	} else {
		// Move reserve coins from the paying reserve to the vesting reserve account
		if err := k.bankKeeper.SendCoins(ctx, payingReserveAddr, vestingReserveAddr, sdk.NewCoins(reserveCoin)); err != nil {
			return err
		}

		remaining := reserveCoin
		for i, schedule := range auction.GetVestingSchedules() {
			payingAmt := math.LegacyNewDecFromInt(reserveCoin.Amount).MulTruncate(schedule.Weight).TruncateInt()

			// All the remaining paying coin goes to the last vesting queue
			if i == vsLen-1 {
				payingAmt = remaining.Amount
			}

			if err := k.VestingQueue.Set(
				ctx,
				collections.Join(
					auction.GetId(),
					schedule.ReleaseTime,
				),
				types.VestingQueue{
					AuctionId:   auction.GetId(),
					Auctioneer:  auction.GetAuctioneer().String(),
					PayingCoin:  sdk.NewCoin(payingCoinDenom, payingAmt),
					ReleaseTime: schedule.ReleaseTime,
					Released:    false,
				},
			); err != nil {
				return err
			}

			remaining = remaining.SubAmount(payingAmt)
		}

		if err := auction.SetStatus(types.AuctionStatusVesting); err != nil {
			return err
		}
		if err := k.Auction.Set(ctx, auction.GetId(), auction); err != nil {
			return err
		}
	}

	return nil
}
