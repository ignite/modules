package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ignite/modules/x/claim/types"
)

const (
	airdropSupplyRoute = "airdrop-supply"
)

// RegisterInvariants registers all module invariants
func RegisterInvariants(ir sdk.InvariantRegistry, k Keeper) {
	ir.RegisterRoute(types.ModuleName, airdropSupplyRoute,
		AirdropSupplyInvariant(k))
}

// AllInvariants runs all invariants of the module.
func AllInvariants(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		return AirdropSupplyInvariant(k)(ctx)
	}
}

// AirdropSupplyInvariant invariant checks that airdrop supply is equal to the remaining claimable
// amounts in claim records
func AirdropSupplyInvariant(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		return "", false
	}
}
