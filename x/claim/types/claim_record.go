package types

import (
	"errors"
	"fmt"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Validate checks the claimRecord is valid
func (m ClaimRecord) Validate() error {
	if _, err := sdk.AccAddressFromBech32(m.Address); err != nil {
		return err
	}

	if !m.Claimable.IsPositive() {
		return errors.New("claimable amount must be positive")
	}

	missionIDMap := make(map[uint64]struct{})
	for _, elem := range m.CompletedMissions {
		if _, ok := missionIDMap[elem]; ok {
			return fmt.Errorf("duplicated id for completed mission")
		}
		missionIDMap[elem] = struct{}{}
	}

	return nil
}

// IsMissionCompleted checks if the specified mission ID is completed for the claim record
func (m ClaimRecord) IsMissionCompleted(missionID uint64) bool {
	for _, completed := range m.CompletedMissions {
		if completed == missionID {
			return true
		}
	}
	return false
}

// IsMissionClaimed checks if the specified mission ID is claimed for the claim record
func (m ClaimRecord) IsMissionClaimed(missionID uint64) bool {
	for _, claimed := range m.ClaimedMissions {
		if claimed == missionID {
			return true
		}
	}
	return false
}

// ClaimableFromMission returns the amount claimable for this claim record from the provided mission completion
func (m ClaimRecord) ClaimableFromMission(mission Mission) sdkmath.Int {
	return mission.Weight.Mul(sdkmath.LegacyNewDecFromInt(m.Claimable)).TruncateInt()
}
