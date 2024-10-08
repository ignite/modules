package types

import (
	"errors"
	"fmt"

	sdkmath "cosmossdk.io/math"
)

func CheckAirdropSupply(airdropSupply AirdropSupply, missionMap map[uint64]Mission, claimRecords []ClaimRecord) error {
	claimSum := sdkmath.ZeroInt()
	claimRecordMap := make(map[string]struct{})

	for _, claimRecord := range claimRecords {

		// check claim record completed missions
		claimable := claimRecord.Claimable
		for _, completedMission := range claimRecord.CompletedMissions {
			mission, ok := missionMap[completedMission]
			if !ok {
				return fmt.Errorf("address %s completed a non existing mission %d",
					claimRecord.Address,
					completedMission,
				)
			}

			// reduce claimable with already claimed funds
			claimable = claimable.Sub(claimRecord.ClaimableFromMission(mission))
		}

		claimSum = claimSum.Add(claimable)
		if _, ok := claimRecordMap[claimRecord.Address]; ok {
			return errors.New("duplicated address for claim record")
		}
		claimRecordMap[claimRecord.Address] = struct{}{}
	}

	// verify airdropSupply == sum of claimRecords
	if !airdropSupply.Supply.Amount.Equal(claimSum) {
		return fmt.Errorf("airdrop supply amount %v not equal to sum of claimable amounts %v",
			airdropSupply.Supply.Amount,
			claimSum,
		)
	}

	return nil
}
