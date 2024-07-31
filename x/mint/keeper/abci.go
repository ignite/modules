package keeper

import (
	"context"
	"time"

	"github.com/cosmos/cosmos-sdk/telemetry"

	"github.com/ignite/modules/x/mint/types"
)

// BeginBlocker mints new coins for the previous block.
func (k Keeper) BeginBlocker(ctx context.Context) error {
	defer telemetry.ModuleMeasureSince(types.ModuleName, time.Now(), telemetry.MetricKeyBeginBlocker)

	// fetch stored minter & params
	minter, err := k.Minter.Get(ctx)
	if err != nil {
		return err
	}
	params, err := k.Params.Get(ctx)
	if err != nil {
		return err
	}

	// recalculate inflation rate
	totalStakingSupply, err := k.StakingTokenSupply(ctx)
	if err != nil {
		return err
	}
	bondedRatio, err := k.BondedRatio(ctx)
	if err != nil {
		return err
	}
	minter.Inflation = minter.NextInflationRate(params, bondedRatio)
	minter.AnnualProvisions = minter.NextAnnualProvisions(params, totalStakingSupply)

	if err := k.Minter.Set(ctx, minter); err != nil {
		return err
	}

	// mint coins, update supply
	mintedCoin := minter.BlockProvision(params)
	err = k.MintCoin(ctx, mintedCoin)
	if err != nil {
		return err
	}

	// distribute minted coins according to the defined proportions
	err = k.DistributeMintedCoin(ctx, mintedCoin)
	if err != nil {
		return err
	}

	if mintedCoin.Amount.IsInt64() {
		defer telemetry.ModuleSetGauge(types.ModuleName, float32(mintedCoin.Amount.Int64()), "minted_tokens")
	}

	return nil
}
