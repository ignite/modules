package keeper

import (
	"context"
	"fmt"

	"cosmossdk.io/collections"
	"cosmossdk.io/core/address"
	"cosmossdk.io/core/store"
	"cosmossdk.io/log"
	"cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ignite/modules/pkg/errors"
	"github.com/ignite/modules/x/mint/types"
)

type (
	Keeper struct {
		cdc              codec.BinaryCodec
		addressCodec     address.Codec
		storeService     store.KVStoreService
		logger           log.Logger
		feeCollectorName string

		// the address capable of executing a MsgUpdateParams message.
		// Typically, this should be the x/gov module account.
		authority string

		Schema collections.Schema
		Params collections.Item[types.Params]
		Minter collections.Item[types.Minter]
		// this line is used by starport scaffolding # collection/type

		accountKeeper types.AccountKeeper
		bankKeeper    types.BankKeeper
		stakingKeeper types.StakingKeeper
		distrKeeper   types.DistrKeeper
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	addressCodec address.Codec,
	storeService store.KVStoreService,
	logger log.Logger,
	authority string,
	accountKeeper types.AccountKeeper,
	bankKeeper types.BankKeeper,
	stakingKeeper types.StakingKeeper,
	distrKeeper types.DistrKeeper,
	feeCollectorName string,
) Keeper {
	if _, err := addressCodec.StringToBytes(authority); err != nil {
		panic(fmt.Sprintf("invalid authority address %s: %s", authority, err))
	}

	sb := collections.NewSchemaBuilder(storeService)

	k := Keeper{
		cdc:          cdc,
		addressCodec: addressCodec,
		storeService: storeService,
		authority:    authority,
		logger:       logger,

		accountKeeper:    accountKeeper,
		bankKeeper:       bankKeeper,
		stakingKeeper:    stakingKeeper,
		distrKeeper:      distrKeeper,
		Params:           collections.NewItem(sb, types.ParamsKey, "params", codec.CollValue[types.Params](cdc)),
		Minter:           collections.NewItem(sb, types.MinterKey, "minter", codec.CollValue[types.Minter](cdc)),
		feeCollectorName: feeCollectorName,
		// this line is used by starport scaffolding # collection/instantiate
	}

	schema, err := sb.Build()
	if err != nil {
		panic(err)
	}
	k.Schema = schema

	return k
}

// GetAuthority returns the module's authority.
func (k Keeper) GetAuthority() string {
	return k.authority
}

// Logger returns a module-specific logger.
func (k Keeper) Logger() log.Logger {
	return k.logger.With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// StakingTokenSupply implements an alias call to the underlying staking keeper's
// StakingTokenSupply to be used in BeginBlocker.
func (k Keeper) StakingTokenSupply(ctx context.Context) (math.Int, error) {
	return k.stakingKeeper.StakingTokenSupply(ctx)
}

// BondedRatio implements an alias call to the underlying staking keeper's
// BondedRatio to be used in BeginBlocker.
func (k Keeper) BondedRatio(ctx context.Context) (math.LegacyDec, error) {
	return k.stakingKeeper.BondedRatio(ctx)
}

// MintCoin implements an alias call to the underlying supply keeper's
// MintCoin to be used in BeginBlocker.
func (k Keeper) MintCoin(ctx context.Context, coin sdk.Coin) error {
	return k.bankKeeper.MintCoins(ctx, types.ModuleName, sdk.NewCoins(coin))
}

// GetProportion gets the balance of the `MintedDenom` from minted coins and returns coins according to the `AllocationRatio`.
func (k Keeper) GetProportion(_ context.Context, mintedCoin sdk.Coin, ratio math.LegacyDec) sdk.Coin {
	return sdk.NewCoin(mintedCoin.Denom, math.LegacyNewDecFromInt(mintedCoin.Amount).Mul(ratio).TruncateInt())
}

// DistributeMintedCoin implements distribution of minted coins from mint
// to be used in BeginBlocker.
func (k Keeper) DistributeMintedCoin(ctx context.Context, mintedCoin sdk.Coin) error {
	params, err := k.Params.Get(ctx)
	if err != nil {
		return err
	}
	proportions := params.DistributionProportions

	// allocate staking rewards into fee collector account to be moved to on next begin blocker by staking module
	stakingRewardsCoins := sdk.NewCoins(k.GetProportion(ctx, mintedCoin, proportions.Staking))
	err = k.bankKeeper.SendCoinsFromModuleToModule(ctx, types.ModuleName, k.feeCollectorName, stakingRewardsCoins)
	if err != nil {
		return err
	}

	fundedAddrsCoin := k.GetProportion(ctx, mintedCoin, proportions.FundedAddresses)
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
			fundedAddrCoins := sdk.NewCoins(k.GetProportion(ctx, fundedAddrsCoin, w.Weight))
			devAddr, err := k.addressCodec.StringToBytes(w.Address)
			if err != nil {
				return errors.Critical(err.Error())
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
