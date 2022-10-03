package types

import (
	"github.com/ignite/modules/testutil/sample"
	"math/rand"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func TestParamsValidate(t *testing.T) {
	invalidInflationMin := DefaultParams()
	// set inflation min to larger than inflation max
	invalidInflationMin.InflationMin = invalidInflationMin.InflationMax.Add(invalidInflationMin.InflationMax)

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
			name:    "should prevent validate params with invalidInflationMin",
			params:  invalidInflationMin,
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
		denom   interface{}
		isValid bool
	}{
		{
			name:    "should validate valid mint denom",
			denom:   DefaultMintDenom,
			isValid: true,
		},
		{
			name:    "should prevent validate mint denom with invalid interface",
			denom:   10,
			isValid: false,
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
		value   interface{}
		isValid bool
	}{
		{
			name:    "should prevent validate dec with invalid interface",
			value:   "string",
			isValid: false,
		},
		{
			name:    "should prevent validate dec with negative value",
			value:   sdk.NewDec(-1),
			isValid: false,
		}, {
			name:    "should prevent validate dec too large a value",
			value:   sdk.NewDec(2),
			isValid: false,
		},
		{
			name:  "should validate valid dec",
			value: DefaultInflationRateChange,
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
		value   interface{}
		isValid bool
	}{
		{
			name:    "should prevent validate blocks per year with invalid interface",
			value:   "string",
			isValid: false,
		},
		{
			name:    "should prevent validate blocks per year with zero value",
			value:   uint64(0),
			isValid: false,
		},
		{
			name:  "should validate valid blocks per year",
			value: DefaultBlocksPerYear,
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
		distrProportions interface{}
		isValid          bool
	}{
		{
			name:             "should prevent validate distribution proportions with invalid interface",
			distrProportions: "string",
			isValid:          false,
		},
		{
			name: "should prevent validate distribution proportions with negative staking ratio",
			distrProportions: DistributionProportions{
				Staking:         sdk.NewDecWithPrec(-3, 1), // -0.3
				FundedAddresses: sdk.NewDecWithPrec(4, 1),  // 0.4
				CommunityPool:   sdk.NewDecWithPrec(3, 1),  // 0.3
			},
			isValid: false,
		},
		{
			name: "should prevent validate distribution proportions with negative funded addresses ratio",
			distrProportions: DistributionProportions{
				Staking:         sdk.NewDecWithPrec(3, 1),  // 0.3
				FundedAddresses: sdk.NewDecWithPrec(-4, 1), // -0.4
				CommunityPool:   sdk.NewDecWithPrec(3, 1),  // 0.3
			},
			isValid: false,
		},
		{
			name: "should prevent validate distribution proportions with negative community pool ratio",
			distrProportions: DistributionProportions{
				Staking:         sdk.NewDecWithPrec(3, 1),  // 0.3
				FundedAddresses: sdk.NewDecWithPrec(4, 1),  // 0.4
				CommunityPool:   sdk.NewDecWithPrec(-3, 1), // -0.3
			},
			isValid: false,
		},
		{
			name: "should prevent validate distribution proportions total ratio not equal to 1",
			distrProportions: DistributionProportions{
				Staking:         sdk.NewDecWithPrec(3, 1),  // 0.3
				FundedAddresses: sdk.NewDecWithPrec(4, 1),  // 0.4
				CommunityPool:   sdk.NewDecWithPrec(31, 2), // 0.31
			},
			isValid: false,
		},
		{
			name:             "should validate valid distribution proportions",
			distrProportions: DefaultDistributionProportions,
			isValid:          true,
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
		weightedAddresses interface{}
		isValid           bool
	}{
		{
			name:              "should prevent validate weighed addresses with invalid interface",
			weightedAddresses: "string",
			isValid:           false,
		},
		{
			name: "should prevent validate weighed addresses with invalid SDK address",
			weightedAddresses: []WeightedAddress{
				{
					Address: "invalid",
					Weight:  sdk.OneDec(),
				},
			},
			isValid: false,
		},
		{
			name: "should prevent validate weighed addresses with negative value",
			weightedAddresses: []WeightedAddress{
				{
					Address: sample.Address(r),
					Weight:  sdk.NewDec(-1),
				},
			},
			isValid: false,
		},
		{
			name: "should prevent validate weighed addresses with weight greater than 1",
			weightedAddresses: []WeightedAddress{
				{
					Address: sample.Address(r),
					Weight:  sdk.NewDec(2),
				},
			},
			isValid: false,
		},
		{
			name: "should prevent validate weighed addresses with sum greater than 1",
			weightedAddresses: []WeightedAddress{
				{
					Address: sample.Address(r),
					Weight:  sdk.NewDecWithPrec(6, 1),
				},
				{
					Address: sample.Address(r),
					Weight:  sdk.NewDecWithPrec(5, 1),
				},
			},
			isValid: false,
		},
		{
			name:              "should validate valid empty weighted addresses",
			weightedAddresses: DefaultFundedAddresses,
		},
		{
			name: "should validate valid  weighted addresses",
			weightedAddresses: []WeightedAddress{
				{
					Address: sample.Address(r),
					Weight:  sdk.NewDecWithPrec(5, 1),
				},
				{
					Address: sample.Address(r),
					Weight:  sdk.NewDecWithPrec(5, 1),
				},
			},
			isValid: true,
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
