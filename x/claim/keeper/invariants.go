package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ignite/modules/x/claim/types"
)

const (
	airdropSupplyRoute = "airdrop-supply"
	claimRecordRoute   = "claim-record"
)

// RegisterInvariants registers all module invariants
func RegisterInvariants(ir sdk.InvariantRegistry, k Keeper) {
	ir.RegisterRoute(types.ModuleName, airdropSupplyRoute,
		AirdropSupplyInvariant(k))
	ir.RegisterRoute(types.ModuleName, claimRecordRoute,
		ClaimRecordInvariant(k))
}

// AllInvariants runs all invariants of the module.
func AllInvariants(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		res, stop := AirdropSupplyInvariant(k)(ctx)
		if stop {
			return res, stop
		}
		return ClaimRecordInvariant(k)(ctx)
	}
}

// AirdropSupplyInvariant invariant checks that airdrop supply is equal to the remaining claimable
// amounts in claim records
func AirdropSupplyInvariant(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		missions := k.GetAllMission(ctx)
		claimRecords := k.GetAllClaimRecord(ctx)
		airdropSupply, _ := k.GetAirdropSupply(ctx)

		missionMap := make(map[uint64]types.Mission)
		for _, mission := range missions {
			missionMap[mission.MissionID] = mission
		}

		err := types.CheckAirdropSupply(airdropSupply, missionMap, claimRecords)
		if err != nil {
			return err.Error(), true
		}

		return "", false
	}
}

// ClaimRecordInvariant invariant checks that claim record was claimed but not completed
func ClaimRecordInvariant(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		missions := k.GetAllMission(ctx)
		claimRecords := k.GetAllClaimRecord(ctx)

		for _, claimRecord := range claimRecords {
			for _, mission := range missions {
				if !claimRecord.IsMissionCompleted(mission.MissionID) &&
					claimRecord.IsMissionClaimed(mission.MissionID) {
					return fmt.Sprintf("mission %d claimed but not completed", mission.MissionID), true
				}
			}
		}
		return "", false
	}
}
