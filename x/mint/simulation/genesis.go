package simulation

// DONTCOVER

import (
	"encoding/json"
	"fmt"
	"math/rand"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"

	"github.com/ignite/modules/testutil/sample"
	"github.com/ignite/modules/x/mint/types"
)

// Simulation parameter constants
const (
	Inflation               = "inflation"
	InflationRateChange     = "inflation_rate_change"
	InflationMax            = "inflation_max"
	InflationMin            = "inflation_min"
	GoalBonded              = "goal_bonded"
	DistributionProportions = "distribution_proportions"
	FundedAddresses         = "funded_addresses"
)

// GenInflation randomized Inflation
func GenInflation(r *rand.Rand) sdkmath.LegacyDec {
	return sdkmath.LegacyNewDecWithPrec(int64(r.Intn(99)), 2)
}

// GenInflationRateChange randomized InflationRateChange
func GenInflationRateChange(r *rand.Rand) sdkmath.LegacyDec {
	return sdkmath.LegacyNewDecWithPrec(int64(r.Intn(99)), 2)
}

// GenInflationMax randomized InflationMax
func GenInflationMax() sdkmath.LegacyDec {
	return sdkmath.LegacyNewDecWithPrec(20, 2)
}

// GenInflationMin randomized InflationMin
func GenInflationMin() sdkmath.LegacyDec {
	return sdkmath.LegacyNewDecWithPrec(7, 2)
}

// GenGoalBonded randomized GoalBonded
func GenGoalBonded() sdkmath.LegacyDec {
	return sdkmath.LegacyNewDecWithPrec(67, 2)
}

// GenDistributionProportions randomized DistributionProportions
func GenDistributionProportions(r *rand.Rand) types.DistributionProportions {
	staking := r.Int63n(99)
	left := int64(100) - staking
	funded := r.Int63n(left)
	communityPool := left - funded

	return types.DistributionProportions{
		Staking:         sdkmath.LegacyNewDecWithPrec(staking, 2),
		FundedAddresses: sdkmath.LegacyNewDecWithPrec(funded, 2),
		CommunityPool:   sdkmath.LegacyNewDecWithPrec(communityPool, 2),
	}
}

func GenFundedAddresses(r *rand.Rand) []types.WeightedAddress {
	var (
		addrs         = make([]types.WeightedAddress, 0)
		numAddrs      = r.Intn(51)
		remainWeight  = sdkmath.LegacyNewDec(1)
		maxRandWeight = sdkmath.LegacyNewDecWithPrec(15, 3)
		minRandWeight = sdkmath.LegacyNewDecWithPrec(5, 3)
	)
	for i := 0; i < numAddrs; i++ {
		// each address except the last can have a max of 2% weight and a min of 0.5%
		weight := simtypes.RandomDecAmount(r, maxRandWeight).Add(minRandWeight)
		if i == numAddrs-1 {
			// use residual weight if last address
			weight = remainWeight
		} else {
			remainWeight = remainWeight.Sub(weight)
		}
		wa := types.WeightedAddress{
			Address: sample.AccAddress(),
			Weight:  weight,
		}
		addrs = append(addrs, wa)
	}
	return addrs
}

// RandomizedGenState generates a random GenesisState for mint
func RandomizedGenState(simState *module.SimulationState) {
	// minter
	var inflation sdkmath.LegacyDec
	simState.AppParams.GetOrGenerate(Inflation, &inflation, simState.Rand,
		func(r *rand.Rand) { inflation = GenInflation(r) },
	)

	// params
	var inflationRateChange sdkmath.LegacyDec
	simState.AppParams.GetOrGenerate(InflationRateChange, &inflationRateChange, simState.Rand,
		func(r *rand.Rand) { inflationRateChange = GenInflationRateChange(r) },
	)

	var inflationMax sdkmath.LegacyDec
	simState.AppParams.GetOrGenerate(InflationMax, &inflationMax, simState.Rand,
		func(r *rand.Rand) { inflationMax = GenInflationMax() },
	)

	var inflationMin sdkmath.LegacyDec
	simState.AppParams.GetOrGenerate(InflationMin, &inflationMin, simState.Rand,
		func(r *rand.Rand) { inflationMin = GenInflationMin() },
	)

	var goalBonded sdkmath.LegacyDec
	simState.AppParams.GetOrGenerate(GoalBonded, &goalBonded, simState.Rand,
		func(r *rand.Rand) { goalBonded = GenGoalBonded() },
	)

	var distributionProportions types.DistributionProportions
	simState.AppParams.GetOrGenerate(DistributionProportions, &distributionProportions, simState.Rand,
		func(r *rand.Rand) { distributionProportions = GenDistributionProportions(r) },
	)

	var developmentFundRecipients []types.WeightedAddress
	simState.AppParams.GetOrGenerate(FundedAddresses, &developmentFundRecipients, simState.Rand,
		func(r *rand.Rand) { developmentFundRecipients = GenFundedAddresses(r) },
	)

	mintDenom := sdk.DefaultBondDenom
	blocksPerYear := uint64(60 * 60 * 8766 / 5)
	params := types.NewParams(mintDenom, inflationRateChange, inflationMax, inflationMin, goalBonded, blocksPerYear, distributionProportions, developmentFundRecipients)

	mintGenesis := types.GenesisState{
		Minter: types.InitialMinter(inflation),
		Params: params,
	}

	bz, err := json.MarshalIndent(&mintGenesis, "", " ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Selected randomly generated minting parameters:\n%s\n", bz)
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&mintGenesis)
}
