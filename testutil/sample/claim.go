package sample

import (
	"math/rand"

	sdkmath "cosmossdk.io/math"

	claim "github.com/ignite/modules/x/claim/types"
)

func ClaimRecord(r *rand.Rand) claim.ClaimRecord {
	return claim.ClaimRecord{
		Address:           Address(r),
		Claimable:         sdkmath.NewInt(r.Int63n(100000)),
		CompletedMissions: uint64Sequence(r),
	}
}

func Mission(r *rand.Rand) claim.Mission {
	const max = 1_000_000
	maxInt := sdkmath.LegacyNewDec(max)
	weight := sdkmath.LegacyNewDec(r.Int63n(max)).Quo(maxInt)

	return claim.Mission{
		MissionId:   Uint64(r),
		Description: String(r, 20),
		Weight:      weight,
	}
}

func uint64Sequence(r *rand.Rand) []uint64 {
	listLen := r.Int63n(10)
	list := make([]uint64, int(listLen))

	for i := range list {
		list[i] = r.Uint64()
	}

	return list
}
