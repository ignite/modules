package types

import (
	"errors"
	"fmt"

	sdkmath "cosmossdk.io/math"
)

const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		ClaimRecordList: []ClaimRecord{},
		MissionList:     []Mission{},
		InitialClaim:    InitialClaim{},
		AirdropSupply:   AirdropSupply{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// check airdrop supply
	err := gs.AirdropSupply.Supply.Validate()
	if err != nil {
		return err
	}

	// Check for duplicated index in claimRecord
	claimRecordIndexMap := make(map[string]struct{})

	for _, elem := range gs.ClaimRecordList {
		index := fmt.Sprint(elem.Address)
		if _, ok := claimRecordIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for claimRecord")
		}
		claimRecordIndexMap[index] = struct{}{}
	}

	// Check for duplicated ID in mission
	missionMap := make(map[uint64]Mission)
	missionCount := gs.GetMissionCount()
	weightSum := sdkmath.LegacyZeroDec()
	for _, elem := range gs.MissionList {
		err := elem.Validate()
		if err != nil {
			return err
		}

		weightSum = weightSum.Add(elem.Weight)

		if _, ok := missionMap[elem.MissionID]; ok {
			return fmt.Errorf("duplicated id for mission")
		}
		if elem.MissionID >= missionCount {
			return fmt.Errorf("mission id should be lower or equal than the last id")
		}
		missionMap[elem.MissionID] = elem
	}

	// ensure mission weight sum is 1
	if len(gs.MissionList) > 0 {
		if !weightSum.Equal(sdkmath.LegacyOneDec()) {
			return errors.New("sum of mission weights must be 1")
		}
	}

	// check initial claim mission exist if enabled
	if gs.InitialClaim.Enabled {
		if _, ok := missionMap[gs.InitialClaim.MissionID]; !ok {
			return errors.New("initial claim mission doesn't exist")
		}
	}

	for _, claimRecord := range gs.ClaimRecordList {
		err := claimRecord.Validate()
		if err != nil {
			return err
		}

	}

	if err := CheckAirdropSupply(gs.AirdropSupply.Supply, missionMap, gs.ClaimRecordList); err != nil {
		return err
	}

	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
