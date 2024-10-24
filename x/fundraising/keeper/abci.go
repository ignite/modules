package keeper

import (
	"context"
	"time"

	"github.com/cosmos/cosmos-sdk/telemetry"

	"github.com/ignite/modules/x/fundraising/types"
)

func (k Keeper) BeginBlocker(ctx context.Context) error {
	defer telemetry.ModuleMeasureSince(types.ModuleName, time.Now(), telemetry.MetricKeyEndBlocker)

	// Get all auctions from the store and execute operations depending on auction status.
	auctions, err := k.Auctions(ctx)
	if err != nil {
		return err
	}
	for _, auction := range auctions {
		switch auction.GetStatus() {
		case types.AuctionStatusStandBy:
			if err := k.ExecuteStandByStatus(ctx, auction); err != nil {
				return err
			}
		case types.AuctionStatusStarted:
			if err := k.ExecuteStartedStatus(ctx, auction); err != nil {
				return err
			}
		case types.AuctionStatusVesting:
			if err := k.ExecuteVestingStatus(ctx, auction); err != nil {
				return err
			}
		}
	}
	return nil
}
