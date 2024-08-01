package keeper

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/types/query"

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
		missions, _, err := query.CollectionPaginate(ctx, k.Mission, nil,
			func(_ uint64, value types.Mission) (types.Mission, error) {
				return value, nil
			},
		)
		if err != nil {
			return "", false
		}

		claimRecords, _, err := query.CollectionPaginate(ctx, k.ClaimRecord, nil,
			func(_ string, value types.ClaimRecord) (types.ClaimRecord, error) {
				return value, nil
			},
		)
		airdropSupply, err := k.AirdropSupply.Get(ctx)
		if err != nil {
			return "", false
		}

		missionMap := make(map[uint64]types.Mission)
		for _, mission := range missions {
			missionMap[mission.MissionID] = mission
		}

		if err := types.CheckAirdropSupply(airdropSupply.Supply, missionMap, claimRecords); err != nil {
			return err.Error(), true
		}

		return "", false
	}
}

// ClaimRecordInvariant invariant checks that claim record was claimed but not completed
func ClaimRecordInvariant(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		missions, _, err := query.CollectionPaginate(ctx, k.Mission, nil,
			func(_ uint64, value types.Mission) (types.Mission, error) {
				return value, nil
			},
		)
		if err != nil {
			return "", false
		}

		claimRecords, _, err := query.CollectionPaginate(ctx, k.ClaimRecord, nil,
			func(_ string, value types.ClaimRecord) (types.ClaimRecord, error) {
				return value, nil
			},
		)

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
		missions, _, err := query.CollectionPaginate(ctx, k.Mission, nil,
			func(_ uint64, value types.Mission) (types.Mission, error) {
				return value, nil
			},
		)
		if err != nil {
			return "", false
		}

		claimRecords, _, err := query.CollectionPaginate(ctx, k.ClaimRecord, nil,
			func(_ string, value types.ClaimRecord) (types.ClaimRecord, error) {
				return value, nil
			},
		)

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
