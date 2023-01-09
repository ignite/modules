package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ignite/modules/x/claim/types"
)

const (
	airdropSupplyRoute      = "airdrop-supply"
	claimRecordRoute        = "claim-record"
	claimRecordMissionRoute = "claim-record-mission"
)

// RegisterInvariants registers all module invariants
func RegisterInvariants(ir sdk.InvariantRegistry, k Keeper) {
	ir.RegisterRoute(types.ModuleName, airdropSupplyRoute,
		AirdropSupplyInvariant(k))
	ir.RegisterRoute(types.ModuleName, claimRecordRoute,
		ClaimRecordInvariant(k))
	ir.RegisterRoute(types.ModuleName, claimRecordMissionRoute,
		ClaimRecordMissionInvariant(k))
}

// AllInvariants runs all invariants of the module.
func AllInvariants(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		res, stop := AirdropSupplyInvariant(k)(ctx)
		if stop {
			return res, stop
		}
		res, stop = ClaimRecordMissionInvariant(k)(ctx)
		if stop {
			return res, stop
		}
		return AirdropSupplyInvariant(k)(ctx)
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

// ClaimRecordMissionInvariant invariant checks that claim record completed missions exist
func ClaimRecordMissionInvariant(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		missions := k.GetAllMission(ctx)
		claimRecords := k.GetAllClaimRecord(ctx)

		missionMap := make(map[uint64]struct{})
		for _, mission := range missions {
			missionMap[mission.MissionID] = struct{}{}
		}
		for _, claimRecord := range claimRecords {
			for _, mission := range claimRecord.CompletedMissions {
				if _, ok := missionMap[mission]; !ok {
					return fmt.Sprintf("mission %d not exist", mission), true
				}
			}
		}

		return "", false
	}
}
