package types

import (
	"errors"

	sdkmath "cosmossdk.io/math"
)

const (
	// MissionIDStaking is the mission ID for staking mission to claim airdrop
	MissionIDStaking = 1
	// MissionIDVoting is the mission ID for voting mission to claim airdrop
	MissionIDVoting = 2
)

// Validate checks the mission is valid
func (m Mission) Validate() error {
	if m.Weight.LT(sdkmath.LegacyZeroDec()) || m.Weight.GT(sdkmath.LegacyOneDec()) {
		return errors.New("mission weight must be in range [0:1]")
	}

	return nil
}
