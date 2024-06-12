package keeper

import (
	"context"

	"cosmossdk.io/log"
	sdkmath "cosmossdk.io/math"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	errorsignite "github.com/ignite/modules/pkg/errors"
	"github.com/ignite/modules/x/mint/types"
)

// Keeper of the mint store
type Keeper struct {
	cdc           codec.BinaryCodec
	storeKey      storetypes.StoreKey
	stakingKeeper types.StakingKeeper
	accountKeeper types.AccountKeeper
	bankKeeper    types.BankKeeper
	distrKeeper   types.DistrKeeper

	authority        string
	feeCollectorName string
}

// NewKeeper creates a new mint Keeper instance
func NewKeeper(cdc codec.BinaryCodec, key storetypes.StoreKey, authority string, sk types.StakingKeeper, ak types.AccountKeeper, bk types.BankKeeper, dk types.DistrKeeper, feeCollectorName string) Keeper {
	// ensure mint module account is set
	if addr := ak.GetModuleAddress(types.ModuleName); addr == nil {
		panic("the mint module account has not been set")
	}

	return Keeper{
		cdc:              cdc,
		storeKey:         key,
		stakingKeeper:    sk,
		accountKeeper:    ak,
		bankKeeper:       bk,
		distrKeeper:      dk,
		authority:        authority,
		feeCollectorName: feeCollectorName,
	}
}

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx context.Context) log.Logger {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	return sdkCtx.Logger().With("module", "x/"+types.ModuleName)
}

func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	store := ctx.KVStore(k.storeKey)
	b := k.cdc.MustMarshal(&params)
	store.Set(types.ParamsKey, b)
}

func (k Keeper) GetParams(ctx sdk.Context) (params types.Params) {
	store := ctx.KVStore(k.storeKey)
	b := store.Get(types.ParamsKey)
	if b == nil {
		panic("stored mint params should not have been nil")
	}

	k.cdc.MustUnmarshal(b, &params)
	return
}

// GetMinter gets the minter
func (k Keeper) GetMinter(ctx sdk.Context) (minter types.Minter) {
	store := ctx.KVStore(k.storeKey)
	b := store.Get(types.MinterKey)
	if b == nil {
		panic("stored minter should not have been nil")
	}

	k.cdc.MustUnmarshal(b, &minter)
	return
}

// SetMinter sets the minter
func (k Keeper) SetMinter(ctx sdk.Context, minter types.Minter) {
	store := ctx.KVStore(k.storeKey)
	b := k.cdc.MustMarshal(&minter)
	store.Set(types.MinterKey, b)
}

// StakingTokenSupply implements an alias call to the underlying staking keeper's
// StakingTokenSupply to be used in BeginBlocker.
func (k Keeper) StakingTokenSupply(ctx context.Context) (sdkmath.Int, error) {
	return k.stakingKeeper.StakingTokenSupply(ctx)
}

// BondedRatio implements an alias call to the underlying staking keeper's
// BondedRatio to be used in BeginBlocker.
func (k Keeper) BondedRatio(ctx context.Context) (sdkmath.LegacyDec, error) {
	return k.stakingKeeper.BondedRatio(ctx)
}

// MintCoin implements an alias call to the underlying supply keeper's
// MintCoin to be used in BeginBlocker.
func (k Keeper) MintCoin(ctx sdk.Context, coin sdk.Coin) error {
	return k.bankKeeper.MintCoins(ctx, types.ModuleName, sdk.NewCoins(coin))
}

// GetProportion gets the balance of the `MintedDenom` from minted coins and returns coins according to the `AllocationRatio`.
func (k Keeper) GetProportion(mintedCoin sdk.Coin, ratio sdkmath.LegacyDec) sdk.Coin {
	return sdk.NewCoin(mintedCoin.Denom, sdkmath.LegacyNewDecFromInt(mintedCoin.Amount).Mul(ratio).TruncateInt())
}

// DistributeMintedCoin implements distribution of minted coins from mint
// to be used in BeginBlocker.
func (k Keeper) DistributeMintedCoin(goCtx context.Context, mintedCoin sdk.Coin) error {
	ctx := sdk.UnwrapSDKContext(goCtx)

	params := k.GetParams(ctx)
	proportions := params.DistributionProportions

	// allocate staking rewards into fee collector account to be moved to on next begin blocker by staking module
	stakingRewardsCoins := sdk.NewCoins(k.GetProportion(mintedCoin, proportions.Staking))
	err := k.bankKeeper.SendCoinsFromModuleToModule(ctx, types.ModuleName, k.feeCollectorName, stakingRewardsCoins)
	if err != nil {
		return err
	}

	fundedAddrsCoin := k.GetProportion(mintedCoin, proportions.FundedAddresses)
	fundedAddrsCoins := sdk.NewCoins(fundedAddrsCoin)
	if len(params.FundedAddresses) == 0 {
		// fund community pool when rewards address is empty
		if err = k.distrKeeper.FundCommunityPool(
			ctx,
			fundedAddrsCoins,
			k.accountKeeper.GetModuleAddress(types.ModuleName),
		); err != nil {
			return err
		}
	} else {
		// allocate developer rewards to developer addresses by weight
		for _, w := range params.FundedAddresses {
			fundedAddrCoins := sdk.NewCoins(k.GetProportion(fundedAddrsCoin, w.Weight))
			devAddr, err := sdk.AccAddressFromBech32(w.Address)
			if err != nil {
				return errorsignite.Critical(err.Error())
			}
			err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, devAddr, fundedAddrCoins)
			if err != nil {
				return err
			}
		}
	}

	// subtract from original provision to ensure no coins left over after the allocations
	communityPoolCoins := sdk.NewCoins(mintedCoin).Sub(stakingRewardsCoins...).Sub(fundedAddrsCoins...)
	err = k.distrKeeper.FundCommunityPool(ctx, communityPoolCoins, k.accountKeeper.GetModuleAddress(types.ModuleName))
	if err != nil {
		return err
	}

	return err
}
