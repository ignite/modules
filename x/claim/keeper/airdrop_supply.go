package keeper

import (
	"context"

	"cosmossdk.io/collections"
	sdkerrors "cosmossdk.io/errors"
	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ignite/modules/pkg/errors"
	"github.com/ignite/modules/x/claim/types"
)

// InitializeAirdropSupply set the airdrop supply in the store and set the module balance
func (k Keeper) InitializeAirdropSupply(ctx context.Context, airdropSupply sdk.Coin) error {
	// get the eventual existing balance of the module for the airdrop supply
	moduleBalance := k.bankKeeper.GetBalance(
		ctx,
		k.accountKeeper.GetModuleAddress(types.ModuleName),
		airdropSupply.Denom,
	)

	// if the module has an existing balance, we burn the entire balance
	if moduleBalance.IsPositive() {
		if err := k.bankKeeper.BurnCoins(ctx, types.ModuleName, sdk.NewCoins(moduleBalance)); err != nil {
			return errors.Criticalf("can't burn module balance %s", err.Error())
		}
	}

	// set the module balance with the airdrop supply
	if err := k.bankKeeper.MintCoins(ctx, types.ModuleName, sdk.NewCoins(airdropSupply)); err != nil {
		return errors.Criticalf("can't mint airdrop supply into module balance %s", err.Error())
	}

	return k.AirdropSupply.Set(ctx, types.AirdropSupply{Supply: airdropSupply})
}

func (k Keeper) EndAirdrop(ctx context.Context) error {
	airdropSupply, err := k.AirdropSupply.Get(ctx)
	if err != nil && !sdkerrors.IsOf(err, collections.ErrNotFound) {
		return err
	}
	if sdkerrors.IsOf(err, collections.ErrNotFound) || !airdropSupply.Supply.IsPositive() {
		return nil
	}

	params, err := k.Params.Get(ctx)
	if err != nil {
		return err
	}

	decayInfo := params.DecayInformation
	blockTime := sdk.UnwrapSDKContext(ctx).BlockTime()
	if decayInfo.Enabled && blockTime.After(decayInfo.DecayEnd) {
		err := k.distrKeeper.FundCommunityPool(
			ctx,
			sdk.NewCoins(airdropSupply.Supply),
			k.accountKeeper.GetModuleAddress(types.ModuleName))
		if err != nil {
			return err
		}

		airdropSupply.Supply.Amount = sdkmath.ZeroInt()
		if err := k.AirdropSupply.Set(ctx, airdropSupply); err != nil {
			return err
		}
	}

	// TODO
	// handle other options:
	// https://github.com/ignite/modules/issues/53
	return nil
}
