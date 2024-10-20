// Package networksuite provides base test suite for tests that need a local network instance
package networksuite

import (
	"math/rand"

	sdkmath "cosmossdk.io/math"
	"github.com/gogo/protobuf/proto"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/ignite/modules/testutil/network"
	"github.com/ignite/modules/testutil/nullify"
	"github.com/ignite/modules/testutil/sample"
	claim "github.com/ignite/modules/x/claim/types"
)

// NetworkTestSuite is a test suite for query tests that initializes a network instance
type NetworkTestSuite struct {
	suite.Suite
	Network    *network.Network
	ClaimState claim.GenesisState
}

// SetupSuite setups the local network with a genesis state
func (nts *NetworkTestSuite) SetupSuite() {
	r := sample.Rand()
	cfg := network.DefaultConfig()

	updateConfigGenesisState := func(moduleName string, moduleState proto.Message) {
		buf, err := cfg.Codec.MarshalJSON(moduleState)
		require.NoError(nts.T(), err)
		cfg.GenesisState[moduleName] = buf
	}

	// initialize claim
	require.NoError(nts.T(), cfg.Codec.UnmarshalJSON(cfg.GenesisState[claim.ModuleName], &nts.ClaimState))
	nts.ClaimState = populateClaim(r, nts.ClaimState)
	updateConfigGenesisState(claim.ModuleName, &nts.ClaimState)

	nts.Network = network.New(nts.T(), cfg)
}

func populateClaim(r *rand.Rand, claimState claim.GenesisState) claim.GenesisState {
	claimState.AirdropSupply = claim.AirdropSupply{Supply: sample.Coin(r)}
	totalSupply := sdkmath.ZeroInt()
	for i := 0; i < 5; i++ {
		// fill claim records
		accSupply := sdkmath.NewIntFromUint64(r.Uint64() % 1000)
		claimRecord := claim.ClaimRecord{
			Claimable: accSupply,
			Address:   sample.Address(r),
		}
		totalSupply = totalSupply.Add(accSupply)
		nullify.Fill(&claimRecord)
		claimState.ClaimRecordList = append(claimState.ClaimRecordList, claimRecord)
	}
	claimState.AirdropSupply.Supply.Amount = totalSupply

	// add missions
	for i := 0; i < 5; i++ {
		mission := claim.Mission{
			MissionId: uint64(i),
			Weight:    sdkmath.LegacyNewDec(r.Int63()),
		}
		nullify.Fill(&mission)
		claimState.MissionList = append(claimState.MissionList, mission)
	}

	return claimState
}
