package types

import (
	"errors"
	"fmt"
	"strings"

	yaml "gopkg.in/yaml.v2"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	DefaultMintDenom               = sdk.DefaultBondDenom
	DefaultInflationRateChange     = sdk.NewDecWithPrec(13, 2)
	DefaultInflationMax            = sdk.NewDecWithPrec(20, 2)
	DefaultInflationMin            = sdk.NewDecWithPrec(7, 2)
	DefaultGoalBonded              = sdk.NewDecWithPrec(67, 2)
	DefaultBlocksPerYear           = uint64(60 * 60 * 8766 / 5) // assuming 5 seconds block times
	DefaultDistributionProportions = DistributionProportions{
		Staking:         sdk.NewDecWithPrec(3, 1), // 0.3
		FundedAddresses: sdk.NewDecWithPrec(4, 1), // 0.4
		CommunityPool:   sdk.NewDecWithPrec(3, 1), // 0.3
	}
	DefaultFundedAddresses []WeightedAddress
)

func NewParams(
	mintDenom string,
	inflationRateChange,
	inflationMax,
	inflationMin,
	goalBonded sdk.Dec,
	blocksPerYear uint64,
	proportions DistributionProportions,
	fundedAddrs []WeightedAddress,
) Params {
	return Params{
		MintDenom:               mintDenom,
		InflationRateChange:     inflationRateChange,
		InflationMax:            inflationMax,
		InflationMin:            inflationMin,
		GoalBonded:              goalBonded,
		BlocksPerYear:           blocksPerYear,
		DistributionProportions: proportions,
		FundedAddresses:         fundedAddrs,
	}
}

// DefaultParams returns default minting module parameters
func DefaultParams() Params {
	return NewParams(
		DefaultMintDenom,
		DefaultInflationRateChange,
		DefaultInflationMax,
		DefaultInflationMin,
		DefaultGoalBonded,
		DefaultBlocksPerYear,
		DefaultDistributionProportions,
		DefaultFundedAddresses,
	)
}

// Validate validates all params
func (p Params) Validate() error {
	if err := validateMintDenom(p.MintDenom); err != nil {
		return err
	}
	if err := validateDec(p.InflationRateChange); err != nil {
		return err
	}
	if err := validateDec(p.InflationMax); err != nil {
		return err
	}
	if err := validateDec(p.InflationMin); err != nil {
		return err
	}
	if err := validateDec(p.GoalBonded); err != nil {
		return err
	}
	if err := validateBlocksPerYear(p.BlocksPerYear); err != nil {
		return err
	}
	if p.InflationMax.LT(p.InflationMin) {
		return fmt.Errorf(
			"max inflation (%s) must be greater than or equal to min inflation (%s)",
			p.InflationMax, p.InflationMin,
		)
	}
	if err := validateDistributionProportions(p.DistributionProportions); err != nil {
		return err
	}
	return validateWeightedAddresses(p.FundedAddresses)
}

// String implements the Stringer interface.
func (p Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}

func validateMintDenom(i interface{}) error {
	v, ok := i.(string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if strings.TrimSpace(v) == "" {
		return errors.New("mint denom cannot be blank")
	}
	if err := sdk.ValidateDenom(v); err != nil {
		return err
	}

	return nil
}

func validateDec(i interface{}) error {
	v, ok := i.(sdk.Dec)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v.IsNegative() {
		return fmt.Errorf("cannot be negative: %s", v)
	}
	if v.GT(sdk.OneDec()) {
		return fmt.Errorf("dec too large: %s", v)
	}

	return nil
}

func validateBlocksPerYear(i interface{}) error {
	v, ok := i.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v == 0 {
		return fmt.Errorf("blocks per year must be positive: %d", v)
	}

	return nil
}

func validateDistributionProportions(i interface{}) error {
	v, ok := i.(DistributionProportions)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v.Staking.IsNegative() {
		return errors.New("staking distribution ratio should not be negative")
	}

	if v.FundedAddresses.IsNegative() {
		return errors.New("funded addresses distribution ratio should not be negative")
	}

	if v.CommunityPool.IsNegative() {
		return errors.New("community pool distribution ratio should not be negative")
	}

	totalProportions := v.Staking.Add(v.FundedAddresses).Add(v.CommunityPool)

	if !totalProportions.Equal(sdk.NewDec(1)) {
		return errors.New("total distributions ratio should be 1")
	}

	return nil
}

func validateWeightedAddresses(i interface{}) error {
	v, ok := i.([]WeightedAddress)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if len(v) == 0 {
		return nil
	}

	weightSum := sdk.NewDec(0)
	for i, w := range v {
		_, err := sdk.AccAddressFromBech32(w.Address)
		if err != nil {
			return fmt.Errorf("invalid address at index %d", i)
		}
		if !w.Weight.IsPositive() {
			return fmt.Errorf("non-positive weight at index %d", i)
		}
		if w.Weight.GT(sdk.NewDec(1)) {
			return fmt.Errorf("more than 1 weight at index %d", i)
		}
		weightSum = weightSum.Add(w.Weight)
	}

	if !weightSum.Equal(sdk.NewDec(1)) {
		return fmt.Errorf("invalid weight sum: %s", weightSum.String())
	}

	return nil
}
