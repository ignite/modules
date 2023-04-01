package simulation

// DONTCOVER

import (
	"fmt"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/ignite/modules/x/mint/types"
	"math/rand"
	"strings"
)

const (
	keyInflationRateChange     = "InflationRateChange"
	keyInflationMax            = "InflationMax"
	keyInflationMin            = "InflationMin"
	keyGoalBonded              = "GoalBonded"
	keyDistributionProportions = "DistributionProportions"
	keyFundedAddresses         = "FundedAddresses"
)

// ParamChanges defines the parameters that can be modified by param change proposals
// on the simulation
func ParamChanges(r *rand.Rand) []simtypes.LegacyParamChange {
	return []simtypes.LegacyParamChange{
		simulation.NewSimLegacyParamChange(types.ModuleName, keyInflationRateChange,
			func(r *rand.Rand) string {
				return fmt.Sprintf("\"%s\"", GenInflationRateChange(r))
			},
		),
		simulation.NewSimLegacyParamChange(types.ModuleName, keyInflationMax,
			func(r *rand.Rand) string {
				return fmt.Sprintf("\"%s\"", GenInflationMax(r))
			},
		),
		simulation.NewSimLegacyParamChange(types.ModuleName, keyInflationMin,
			func(r *rand.Rand) string {
				return fmt.Sprintf("\"%s\"", GenInflationMin(r))
			},
		),
		simulation.NewSimLegacyParamChange(types.ModuleName, keyGoalBonded,
			func(r *rand.Rand) string {
				return fmt.Sprintf("\"%s\"", GenGoalBonded(r))
			},
		),
		simulation.NewSimLegacyParamChange(types.ModuleName, keyDistributionProportions,
			func(r *rand.Rand) string {
				proportions := GenDistributionProportions(r)
				return fmt.Sprintf(
					`{"staking":"%s","funded_addresses":"%s","community_pool":"%s"}`,
					proportions.Staking.String(),
					proportions.FundedAddresses.String(),
					proportions.CommunityPool.String(),
				)
			},
		),
		simulation.NewSimLegacyParamChange(types.ModuleName, keyFundedAddresses,
			func(r *rand.Rand) string {
				weightedAddrs := GenFundedAddresses(r)
				weightedAddrsStr := make([]string, 0)
				for _, wa := range weightedAddrs {
					s := fmt.Sprintf(
						`{"address":"%s","weight":"%s"}`,
						wa.Address,
						wa.Weight.String(),
					)
					weightedAddrsStr = append(weightedAddrsStr, s)
				}
				return fmt.Sprintf("[%s]", strings.Join(weightedAddrsStr, ","))
			},
		),
	}
}
