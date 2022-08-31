package keeper

import (
	sdkmath "cosmossdk.io/math"
	"fmt"
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

		// check missions
		missionMap := make(map[uint64]types.Mission)
		missions := k.GetAllMission(ctx)

		for _, mission := range missions {
			missionMap[mission.MissionID] = mission
		}

		// check claim records
		claimSum := sdkmath.ZeroInt()
		claimRecordMap := make(map[string]struct{})
		claimRecords := k.GetAllClaimRecord(ctx)

		for _, claimRecord := range claimRecords {
			// check claim record completed missions
			claimable := claimRecord.Claimable
			for _, completedMission := range claimRecord.CompletedMissions {
				mission, ok := missionMap[completedMission]
				if !ok {
					return fmt.Sprintf("address %s completed a non existing mission %d",
						claimRecord.Address,
						completedMission,
					), true
				}

				// reduce claimable with already claimed funds
				claimable = claimable.Sub(claimRecord.ClaimableFromMission(mission))
			}

			claimSum = claimSum.Add(claimable)
			if _, ok := claimRecordMap[claimRecord.Address]; ok {
				return "duplicated address for claim record", true
			}
			claimRecordMap[claimRecord.Address] = struct{}{}
		}

		airdropSupply, _ := k.GetAirdropSupply(ctx)

		// verify airdropSupply == sum of claimRecords
		if !airdropSupply.Amount.Equal(claimSum) {
			return "airdrop supply amount not equal to sum of claimable amounts", true
		}

		return "", false
	}
}
