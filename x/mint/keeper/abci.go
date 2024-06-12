package keeper

import (
	"context"
	"time"

	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ignite/modules/x/mint/types"
)

// BeginBlocker mints new coins for the previous block.
func (k Keeper) BeginBlocker(goCtx context.Context) error {
	defer telemetry.ModuleMeasureSince(types.ModuleName, time.Now(), telemetry.MetricKeyBeginBlocker)

	ctx := sdk.UnwrapSDKContext(goCtx)

	// fetch stored minter & params
	minter := k.GetMinter(ctx)
	params := k.GetParams(ctx)

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
	k.SetMinter(ctx, minter)

	// mint coins, update supply
	mintedCoin := minter.BlockProvision(params)
	if err = k.MintCoin(ctx, mintedCoin); err != nil {
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

	return ctx.EventManager().EmitTypedEvent(&types.EventMint{
		BondedRatio:      bondedRatio,
		Inflation:        minter.Inflation,
		AnnualProvisions: minter.AnnualProvisions,
		Amount:           mintedCoin.Amount,
	})
}
