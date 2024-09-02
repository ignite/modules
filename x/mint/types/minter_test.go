package types_test

import (
	"math/rand"
	"testing"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/ignite/modules/x/mint/types"
)

func TestValidateMinter(t *testing.T) {
	invalid := types.DefaultInitialMinter()
	invalid.Inflation = sdkmath.LegacyNewDec(-1)

	tests := []struct {
		name    string
		minter  types.Minter
		isValid bool
	}{
		{
			name:    "should validate valid minter",
			minter:  types.DefaultInitialMinter(),
			isValid: true,
		},
		{
			name:    "should prevent validate for minter with negative inflation",
			minter:  invalid,
			isValid: false,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.minter.Validate()
			if !tc.isValid {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestNextInflation(t *testing.T) {
	minter := types.DefaultInitialMinter()
	params := types.DefaultParams()
	blocksPerYr := sdkmath.LegacyNewDec(int64(params.BlocksPerYear))

	// Governing Mechanism:
	//    inflationRateChangePerYear = (1- BondedRatio/ GoalBonded) * MaxInflationRateChange

	tests := []struct {
		bondedRatio, setInflation, expChange sdkmath.LegacyDec
	}{
		// with 0% bonded coin supply the inflation should increase by InflationRateChange
		{sdkmath.LegacyZeroDec(), sdkmath.LegacyNewDecWithPrec(7, 2), params.InflationRateChange.Quo(blocksPerYr)},

		// 100% bonded, starting at 20% inflation and being reduced
		// (1 - (1/0.67))*(0.13/8667)
		{
			sdkmath.LegacyOneDec(), sdkmath.LegacyNewDecWithPrec(20, 2),
			sdkmath.LegacyOneDec().Sub(sdkmath.LegacyOneDec().Quo(params.GoalBonded)).Mul(params.InflationRateChange).Quo(blocksPerYr),
		},

		// 50% bonded, starting at 10% inflation and being increased
		{
			sdkmath.LegacyNewDecWithPrec(5, 1), sdkmath.LegacyNewDecWithPrec(10, 2),
			sdkmath.LegacyOneDec().Sub(sdkmath.LegacyNewDecWithPrec(5, 1).Quo(params.GoalBonded)).Mul(params.InflationRateChange).Quo(blocksPerYr),
		},

		// test 7% minimum stop (testing with 100% bonded)
		{sdkmath.LegacyOneDec(), sdkmath.LegacyNewDecWithPrec(7, 2), sdkmath.LegacyZeroDec()},
		{sdkmath.LegacyOneDec(), sdkmath.LegacyNewDecWithPrec(700000001, 10), sdkmath.LegacyNewDecWithPrec(-1, 10)},

		// test 20% maximum stop (testing with 0% bonded)
		{sdkmath.LegacyZeroDec(), sdkmath.LegacyNewDecWithPrec(20, 2), sdkmath.LegacyZeroDec()},
		{sdkmath.LegacyZeroDec(), sdkmath.LegacyNewDecWithPrec(1999999999, 10), sdkmath.LegacyNewDecWithPrec(1, 10)},

		// perfect balance shouldn't change inflation
		{sdkmath.LegacyNewDecWithPrec(67, 2), sdkmath.LegacyNewDecWithPrec(15, 2), sdkmath.LegacyZeroDec()},
	}
	for i, tc := range tests {
		minter.Inflation = tc.setInflation

		inflation := minter.NextInflationRate(params, tc.bondedRatio)
		diffInflation := inflation.Sub(tc.setInflation)

		require.True(t, diffInflation.Equal(tc.expChange),
			"Test Index: %v\nDiff:  %v\nExpected: %v\n", i, diffInflation, tc.expChange)
	}
}

func TestBlockProvision(t *testing.T) {
	minter := types.InitialMinter(sdkmath.LegacyNewDecWithPrec(1, 1))
	params := types.DefaultParams()

	secondsPerYear := int64(60 * 60 * 8766)

	tests := []struct {
		annualProvisions int64
		expProvisions    int64
	}{
		{secondsPerYear / 5, 1},
		{secondsPerYear/5 + 1, 1},
		{(secondsPerYear / 5) * 2, 2},
		{(secondsPerYear / 5) / 2, 0},
	}
	for i, tc := range tests {
		minter.AnnualProvisions = sdkmath.LegacyNewDec(tc.annualProvisions)
		provisions := minter.BlockProvision(params)

		expProvisions := sdk.NewCoin(params.MintDenom,
			sdkmath.NewInt(tc.expProvisions))

		require.True(t, expProvisions.IsEqual(provisions),
			"test: %v\n\tExp: %v\n\tGot: %v\n",
			i, tc.expProvisions, provisions)
	}
}

// Benchmarking :)
// previously using sdkmath.Int operations:
// BenchmarkBlockProvision-4 5000000 220 ns/op
//
// using sdkmath.LegacyDec operations: (current implementation)
// BenchmarkBlockProvision-4 3000000 429 ns/op
func BenchmarkBlockProvision(b *testing.B) {
	b.ReportAllocs()
	minter := types.InitialMinter(sdkmath.LegacyNewDecWithPrec(1, 1))
	params := types.DefaultParams()

	s1 := rand.NewSource(100)
	r1 := rand.New(s1)
	minter.AnnualProvisions = sdkmath.LegacyNewDec(r1.Int63n(1000000))

	// run the BlockProvision function b.N times
	for n := 0; n < b.N; n++ {
		minter.BlockProvision(params)
	}
}

// Next inflation benchmarking
// BenchmarkNextInflation-4 1000000 1828 ns/op
func BenchmarkNextInflation(b *testing.B) {
	b.ReportAllocs()
	minter := types.InitialMinter(sdkmath.LegacyNewDecWithPrec(1, 1))
	params := types.DefaultParams()
	bondedRatio := sdkmath.LegacyNewDecWithPrec(1, 1)

	// run the NextInflationRate function b.N times
	for n := 0; n < b.N; n++ {
		minter.NextInflationRate(params, bondedRatio)
	}
}

// Next annual provisions benchmarking
// BenchmarkNextAnnualProvisions-4 5000000 251 ns/op
func BenchmarkNextAnnualProvisions(b *testing.B) {
	b.ReportAllocs()
	minter := types.InitialMinter(sdkmath.LegacyNewDecWithPrec(1, 1))
	params := types.DefaultParams()
	totalSupply := sdkmath.NewInt(100000000000000)

	// run the NextAnnualProvisions function b.N times
	for n := 0; n < b.N; n++ {
		minter.NextAnnualProvisions(params, totalSupply)
	}
}
