package types

import (
	"math/rand"
	"testing"

	sdkmath "cosmossdk.io/math"
	"github.com/stretchr/testify/require"

	"github.com/ignite/modules/testutil/sample"
)

func TestParamsValidate(t *testing.T) {
	tests := []struct {
		name    string
		params  Params
		isValid bool
	}{
		{
			name:    "should validate valid params",
			params:  DefaultParams(),
			isValid: true,
		},
		{
			name: "should prevent validate params with inflation max less than inflation min",
			params: Params{
				MintDenom:               DefaultMintDenom,
				InflationRateChange:     DefaultInflationRateChange,
				InflationMax:            DefaultInflationMax,
				InflationMin:            DefaultInflationMin.Add(DefaultInflationMax),
				GoalBonded:              DefaultGoalBonded,
				BlocksPerYear:           DefaultBlocksPerYear,
				DistributionProportions: DefaultDistributionProportions,
				FundedAddresses:         DefaultFundedAddresses,
			},
			isValid: false,
		},
		{
			name: "should prevent validate params with negative inflation min",
			params: Params{
				MintDenom:               DefaultMintDenom,
				InflationRateChange:     DefaultInflationRateChange,
				InflationMax:            DefaultInflationMax,
				InflationMin:            sdkmath.LegacyNewDec(-1),
				GoalBonded:              DefaultGoalBonded,
				BlocksPerYear:           DefaultBlocksPerYear,
				DistributionProportions: DefaultDistributionProportions,
				FundedAddresses:         DefaultFundedAddresses,
			},
			isValid: false,
		},
		{
			name: "should prevent validate params with negative inflation max",
			params: Params{
				MintDenom:               DefaultMintDenom,
				InflationRateChange:     DefaultInflationRateChange,
				InflationMax:            sdkmath.LegacyNewDec(-1),
				InflationMin:            DefaultInflationMin,
				GoalBonded:              DefaultGoalBonded,
				BlocksPerYear:           DefaultBlocksPerYear,
				DistributionProportions: DefaultDistributionProportions,
				FundedAddresses:         DefaultFundedAddresses,
			},
			isValid: false,
		},
		{
			name: "should prevent validate params with negative goal bonded",
			params: Params{
				MintDenom:               DefaultMintDenom,
				InflationRateChange:     DefaultInflationRateChange,
				InflationMax:            DefaultInflationMax,
				InflationMin:            DefaultInflationMin,
				GoalBonded:              sdkmath.LegacyNewDec(-1),
				BlocksPerYear:           DefaultBlocksPerYear,
				DistributionProportions: DefaultDistributionProportions,
				FundedAddresses:         DefaultFundedAddresses,
			},
			isValid: false,
		},
		{
			name: "should prevent invalid mint denom",
			params: Params{
				MintDenom:               "",
				InflationRateChange:     DefaultInflationRateChange,
				InflationMax:            DefaultInflationMax,
				InflationMin:            DefaultInflationMin,
				GoalBonded:              DefaultGoalBonded,
				BlocksPerYear:           DefaultBlocksPerYear,
				DistributionProportions: DefaultDistributionProportions,
				FundedAddresses:         DefaultFundedAddresses,
			},
			isValid: false,
		},
		{
			name: "should prevent invalid blocks per year",
			params: Params{
				MintDenom:               DefaultMintDenom,
				InflationRateChange:     DefaultInflationRateChange,
				InflationMax:            DefaultInflationMax,
				InflationMin:            DefaultInflationMin,
				GoalBonded:              DefaultGoalBonded,
				BlocksPerYear:           0,
				DistributionProportions: DefaultDistributionProportions,
				FundedAddresses:         DefaultFundedAddresses,
			},
			isValid: false,
		},
		{
			name: "should prevent invalid distribution proportions",
			params: Params{
				MintDenom:           DefaultMintDenom,
				InflationRateChange: DefaultInflationRateChange,
				InflationMax:        DefaultInflationMax,
				InflationMin:        DefaultInflationMin,
				GoalBonded:          DefaultGoalBonded,
				BlocksPerYear:       DefaultBlocksPerYear,
				DistributionProportions: DistributionProportions{
					Staking:         sdkmath.LegacyNewDecWithPrec(3, 1),  // 0.3
					FundedAddresses: sdkmath.LegacyNewDecWithPrec(-4, 1), // -0.4
					CommunityPool:   sdkmath.LegacyNewDecWithPrec(3, 1),  // 0.3
				},
				FundedAddresses: DefaultFundedAddresses,
			},
			isValid: false,
		},
		{
			name: "should prevent invalid weighted addresses",
			params: Params{
				MintDenom:               DefaultMintDenom,
				InflationRateChange:     DefaultInflationRateChange,
				InflationMax:            DefaultInflationMax,
				InflationMin:            DefaultInflationMin,
				GoalBonded:              DefaultGoalBonded,
				BlocksPerYear:           DefaultBlocksPerYear,
				DistributionProportions: DefaultDistributionProportions,
				FundedAddresses: []WeightedAddress{
					{
						Address: "invalid",
						Weight:  sdkmath.LegacyOneDec(),
					},
				},
			},
			isValid: false,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.params.Validate()
			if !tc.isValid {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestValidateMintDenom(t *testing.T) {
	tests := []struct {
		name    string
		denom   string
		isValid bool
	}{
		{
			name:    "should validate valid mint denom",
			denom:   DefaultMintDenom,
			isValid: true,
		},
		{
			name:    "should prevent validate empty mint denom",
			denom:   "",
			isValid: false,
		},
		{
			name:    "should prevent validate mint denom with invalid value",
			denom:   "invalid&",
			isValid: false,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := validateMintDenom(tc.denom)
			if !tc.isValid {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestValidateDec(t *testing.T) {
	tests := []struct {
		name    string
		value   sdkmath.LegacyDec
		isValid bool
	}{
		{
			name:    "should validate valid dec",
			value:   DefaultInflationRateChange,
			isValid: true,
		},
		{
			name:    "should prevent validate dec with negative value",
			value:   sdkmath.LegacyNewDec(-1),
			isValid: false,
		},
		{
			name:    "should prevent validate dec too large a value",
			value:   sdkmath.LegacyNewDec(2),
			isValid: false,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := validateDec(tc.value)
			if !tc.isValid {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestValidateBlocksPerYear(t *testing.T) {
	tests := []struct {
		name    string
		value   uint64
		isValid bool
	}{
		{
			name:    "should validate valid blocks per year",
			value:   DefaultBlocksPerYear,
			isValid: true,
		},
		{
			name:    "should prevent validate blocks per year with zero value",
			value:   uint64(0),
			isValid: false,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := validateBlocksPerYear(tc.value)
			if !tc.isValid {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestValidateDistributionProportions(t *testing.T) {
	tests := []struct {
		name             string
		distrProportions DistributionProportions
		isValid          bool
	}{
		{
			name:             "should validate valid distribution proportions",
			distrProportions: DefaultDistributionProportions,
			isValid:          true,
		},
		{
			name: "should prevent validate distribution proportions with negative staking ratio",
			distrProportions: DistributionProportions{
				Staking:         sdkmath.LegacyNewDecWithPrec(-3, 1), // -0.3
				FundedAddresses: sdkmath.LegacyNewDecWithPrec(4, 1),  // 0.4
				CommunityPool:   sdkmath.LegacyNewDecWithPrec(3, 1),  // 0.3
			},
			isValid: false,
		},
		{
			name: "should prevent validate distribution proportions with negative funded addresses ratio",
			distrProportions: DistributionProportions{
				Staking:         sdkmath.LegacyNewDecWithPrec(3, 1),  // 0.3
				FundedAddresses: sdkmath.LegacyNewDecWithPrec(-4, 1), // -0.4
				CommunityPool:   sdkmath.LegacyNewDecWithPrec(3, 1),  // 0.3
			},
			isValid: false,
		},
		{
			name: "should prevent validate distribution proportions with negative community pool ratio",
			distrProportions: DistributionProportions{
				Staking:         sdkmath.LegacyNewDecWithPrec(3, 1),  // 0.3
				FundedAddresses: sdkmath.LegacyNewDecWithPrec(4, 1),  // 0.4
				CommunityPool:   sdkmath.LegacyNewDecWithPrec(-3, 1), // -0.3
			},
			isValid: false,
		},
		{
			name: "should prevent validate distribution proportions total ratio not equal to 1",
			distrProportions: DistributionProportions{
				Staking:         sdkmath.LegacyNewDecWithPrec(3, 1),  // 0.3
				FundedAddresses: sdkmath.LegacyNewDecWithPrec(4, 1),  // 0.4
				CommunityPool:   sdkmath.LegacyNewDecWithPrec(31, 2), // 0.31
			},
			isValid: false,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := validateDistributionProportions(tc.distrProportions)
			if !tc.isValid {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestValidateWeightedAddresses(t *testing.T) {
	s := rand.NewSource(1)
	r := rand.New(s)

	tests := []struct {
		name              string
		weightedAddresses []WeightedAddress
		isValid           bool
	}{
		{
			name: "should validate valid  weighted addresses",
			weightedAddresses: []WeightedAddress{
				{
					Address: sample.Address(r),
					Weight:  sdkmath.LegacyNewDecWithPrec(5, 1),
				},
				{
					Address: sample.Address(r),
					Weight:  sdkmath.LegacyNewDecWithPrec(5, 1),
				},
			},
			isValid: true,
		},
		{
			name:              "should validate valid empty weighted addresses",
			weightedAddresses: DefaultFundedAddresses,
			isValid:           true,
		},
		{
			name: "should prevent validate weighed addresses with invalid SDK address",
			weightedAddresses: []WeightedAddress{
				{
					Address: "invalid",
					Weight:  sdkmath.LegacyOneDec(),
				},
			},
			isValid: false,
		},
		{
			name: "should prevent validate weighed addresses with negative value",
			weightedAddresses: []WeightedAddress{
				{
					Address: sample.Address(r),
					Weight:  sdkmath.LegacyNewDec(-1),
				},
			},
			isValid: false,
		},
		{
			name: "should prevent validate weighed addresses with weight greater than 1",
			weightedAddresses: []WeightedAddress{
				{
					Address: sample.Address(r),
					Weight:  sdkmath.LegacyNewDec(2),
				},
			},
			isValid: false,
		},
		{
			name: "should prevent validate weighed addresses with sum greater than 1",
			weightedAddresses: []WeightedAddress{
				{
					Address: sample.Address(r),
					Weight:  sdkmath.LegacyNewDecWithPrec(6, 1),
				},
				{
					Address: sample.Address(r),
					Weight:  sdkmath.LegacyNewDecWithPrec(5, 1),
				},
			},
			isValid: false,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := validateWeightedAddresses(tc.weightedAddresses)
			if !tc.isValid {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
		})
	}
}
