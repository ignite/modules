package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ignite/modules/x/claim/types"
)

// SetParams set the params
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	store := ctx.KVStore(k.storeKey)
	b := k.cdc.MustMarshal(&params)
	store.Set(types.ParamsKey, b)
}

// GetParams get all parameters as types.Params
func (k Keeper) GetParams(ctx sdk.Context) (params types.Params) {
	store := ctx.KVStore(k.storeKey)
	b := store.Get(types.ParamsKey)
	if b == nil {
		panic("stored mint params should not have been nil")
	}

	k.cdc.MustUnmarshal(b, &params)
	return
}
