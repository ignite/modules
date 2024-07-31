package constructor

import (
	sdkmath "cosmossdk.io/math"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

// Coin returns a sdk.Coin from a string
func Coin(t testing.TB, str string) sdk.Coin {
	coin, err := sdk.ParseCoinNormalized(str)
	require.NoError(t, err)
	return coin
}

// Coins returns a sdk.Coins from a string
func Coins(t testing.TB, str string) sdk.Coins {
	coins, err := sdk.ParseCoinsNormalized(str)
	require.NoError(t, err)
	return coins
}

// Dec returns a sdk.Dec from a string
func Dec(t testing.TB, str string) sdkmath.LegacyDec {
	dec, err := sdkmath.LegacyNewDecFromStr(str)
	require.NoError(t, err)
	return dec
}
