package types

import (
	"errors"
	"fmt"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		AirdropSupply:   AirdropSupply{Supply: sdk.NewCoin("utest", sdkmath.ZeroInt())},
		ClaimRecordList: []ClaimRecord{},
		MissionCount:    0,
		MissionList:     []Mission{},
		InitialClaim:    InitialClaim{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// check airdrop supply
	if err := gs.AirdropSupply.Supply.Validate(); err != nil {
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

	// check missions
	weightSum := sdkmath.LegacyZeroDec()
	missionCount := gs.GetMissionCount()
	missionMap := make(map[uint64]Mission)
	for _, mission := range gs.MissionList {
		err := mission.Validate()
		if err != nil {
			return err
		}

		weightSum = weightSum.Add(mission.Weight)
		if _, ok := missionMap[mission.MissionId]; ok {
			return errors.New("duplicated id for mission")
		}
		if mission.MissionId >= missionCount {
			return fmt.Errorf("mission id should be lower or equal than the last id")
		}
		missionMap[mission.MissionId] = mission
	}

	// ensure mission weight sum is 1
	if len(gs.MissionList) > 0 {
		if !weightSum.Equal(sdkmath.LegacyOneDec()) {
			return errors.New("sum of mission weights must be 1")
		}
	}

	// check initial claim mission exist if enabled
	if gs.InitialClaim.Enabled {
		if _, ok := missionMap[gs.InitialClaim.MissionId]; !ok {
			return errors.New("initial claim mission doesn't exist")
		}
	}

	for _, claimRecord := range gs.ClaimRecordList {
		if err := claimRecord.Validate(); err != nil {
			return err
		}
	}

	if err := CheckAirdropSupply(gs.AirdropSupply, missionMap, gs.ClaimRecordList); err != nil {
		return err
	}

	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
