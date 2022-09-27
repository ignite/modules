package keeper

import (
	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ignite/modules/x/mint/types"
	"time"
)

// BeginBlocker mints new coins for the previous block.
func (k Keeper) BeginBlocker(ctx sdk.Context) error {
	defer telemetry.ModuleMeasureSince(types.ModuleName, time.Now(), telemetry.MetricKeyBeginBlocker)

	// fetch stored minter & params
	minter := k.GetMinter(ctx)
	params := k.GetParams(ctx)

	// recalculate inflation rate
	totalStakingSupply := k.StakingTokenSupply(ctx)
	bondedRatio := k.BondedRatio(ctx)
	minter.Inflation = minter.NextInflationRate(params, bondedRatio)
	minter.AnnualProvisions = minter.NextAnnualProvisions(params, totalStakingSupply)
	k.SetMinter(ctx, minter)

	// mint coins, update supply
	mintedCoin := minter.BlockProvision(params)
	mintedCoins := sdk.NewCoins(mintedCoin)

	err := k.MintCoins(ctx, mintedCoins)
	if err != nil {
		panic(err)
	}

	// distribute minted coins according to the defined proportions
	err = k.DistributeMintedCoins(ctx, mintedCoin)
	if err != nil {
		panic(err)
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
