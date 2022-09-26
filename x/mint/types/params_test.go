package types

import (
	"errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParamsValidate(t *testing.T) {
	invalid := DefaultParams()
	// set inflation min to larger than inflation max
	invalid.InflationMin = invalid.InflationMax.Add(invalid.InflationMax)

	tests := []struct {
		name   string
		params Params
		err    error
	}{
		{
			name:   "should prevent validate params with inflation min larger than inflation max",
			params: invalid,
			err: errors.New("max inflation (0.200000000000000000) must be greater than or equal " +
				"to min inflation (0.400000000000000000)"),
		},
		{
			name:   "should validate valid params",
			params: DefaultParams(),
			err:    nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.params.Validate()
			if tt.err != nil {
				require.Error(t, err, tt.err)
				require.Equal(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestValidateMintDenom(t *testing.T) {
	tests := []struct {
		name  string
		denom interface{}
		err   error
	}{
		{
			name:  "should prevent validate mint denom with invalid interface",
			denom: 10,
			err:   errors.New("invalid parameter type: int"),
		},
		{
			name:  "should prevent validate empty mint denom",
			denom: "",
			err:   errors.New("mint denom cannot be blank"),
		},
		{
			name:  "should prevent validate mint denom with invalid value",
			denom: "invalid&",
			err:   errors.New("invalid denom: invalid&"),
		},
		{
			name:  "should validate valid mint denom",
			denom: DefaultMintDenom,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateMintDenom(tt.denom)
			if tt.err != nil {
				require.Error(t, err, tt.err)
				require.Equal(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestValidateDec(t *testing.T) {
	tests := []struct {
		name  string
		value interface{}
		err   error
	}{
		{
			name:  "should prevent validate dec with invalid interface",
			value: "string",
			err:   errors.New("invalid parameter type: string"),
		},
		{
			name:  "should prevent validate dec with negative value",
			value: sdk.NewDec(-1),
			err:   errors.New("cannot be negative: -1.000000000000000000"),
		}, {
			name:  "should prevent validate dec too large a value",
			value: sdk.NewDec(2),
			err:   errors.New("dec too large: 2.000000000000000000"),
		},
		{
			name:  "should validate valid dec",
			value: DefaultInflationRateChange,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateDec(tt.value)
			if tt.err != nil {
				require.Error(t, err, tt.err)
				require.Equal(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestValidateBlocksPerYear(t *testing.T) {
	tests := []struct {
		name  string
		value interface{}
		err   error
	}{
		{
			name:  "should prevent validate blocks per year with invalid interface",
			value: "string",
			err:   errors.New("invalid parameter type: string"),
		},
		{
			name:  "should prevent validate blocks per year with zero value",
			value: uint64(0),
			err:   errors.New("blocks per year must be positive: 0"),
		},
		{
			name:  "should validate valid blocks per year",
			value: DefaultBlocksPerYear,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateBlocksPerYear(tt.value)
			if tt.err != nil {
				require.Error(t, err, tt.err)
				require.Equal(t, err, tt.err)
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
		err              error
	}{
		{
			name:             "should prevent validate distribution proportions with invalid interface",
			distrProportions: "string",
			err:              errors.New("invalid parameter type: string"),
		},
		{
			name:             "should prevent validate distribution proportions with negative staking",
			distrProportions: "string",
			err:              errors.New("invalid parameter type: string"),
		},
		{
			name:             "should validate valid distribution proportions",
			distrProportions: DefaultDistributionProportions,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateDistributionProportions(tt.distrProportions)
			if tt.err != nil {
				require.Error(t, err, tt.err)
				require.Equal(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestValidateWeightedAddresses(t *testing.T) {
	tests := []struct {
		name              string
		weightedAddresses interface{}
		err               error
	}{
		{
			name:              "should prevent validate weighed addresses with invalid interface",
			weightedAddresses: "string",
			err:               errors.New("invalid parameter type: string"),
		},
		{
			name:              "should validate valid weighted addresses",
			weightedAddresses: DefaultFundedAddresses,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateWeightedAddresses(tt.weightedAddresses)
			if tt.err != nil {
				require.Error(t, err, tt.err)
				require.Equal(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
