// Package sample provides methods to initialize sample object of various types for test purposes
package sample

import (
	"math/rand"
	"strconv"

	sdkmath "cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	"github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"
	crypto "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"

	claim "github.com/ignite/modules/x/claim/types"
)

// Codec returns a codec with preregistered interfaces
func Codec() codec.Codec {
	interfaceRegistry := codectypes.NewInterfaceRegistry()

	cryptocodec.RegisterInterfaces(interfaceRegistry)
	authtypes.RegisterInterfaces(interfaceRegistry)
	stakingtypes.RegisterInterfaces(interfaceRegistry)
	banktypes.RegisterInterfaces(interfaceRegistry)
	claim.RegisterInterfaces(interfaceRegistry)

	return codec.NewProtoCodec(interfaceRegistry)
}

// Bool returns randomly true or false
func Bool(r *rand.Rand) bool {
	b := r.Intn(100)
	return b < 50
}

// Uint64 returns a random uint64
func Uint64(r *rand.Rand) uint64 {
	return uint64(r.Intn(10000))
}

// String returns a random string of length n
func String(r *rand.Rand, n int) string {
	letter := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	randomString := make([]rune, n)
	for i := range randomString {
		randomString[i] = letter[r.Intn(len(letter))]
	}
	return string(randomString)
}

// AlphaString returns a random string with lowercase alpha char of length n
func AlphaString(r *rand.Rand, n int) string {
	letter := []rune("abcdefghijklmnopqrstuvwxyz")

	randomString := make([]rune, n)
	for i := range randomString {
		randomString[i] = letter[r.Intn(len(letter))]
	}
	return string(randomString)
}

// PubKey returns a sample account PubKey
func PubKey(r *rand.Rand) crypto.PubKey {
	seed := []byte(strconv.Itoa(r.Int()))
	return ed25519.GenPrivKeyFromSecret(seed).PubKey()
}

// AccAddress returns a sample account address
func AccAddress(r *rand.Rand) sdk.AccAddress {
	addr := PubKey(r).Address()
	return sdk.AccAddress(addr)
}

// Address returns a sample string account address
func Address(r *rand.Rand) string {
	return AccAddress(r).String()
}

// Coin returns a sample coin structure
func Coin(r *rand.Rand) sdk.Coin {
	return sdk.NewCoin(AlphaString(r, 5), sdkmath.NewInt(r.Int63n(10000)+1))
}

// Int returns a sample sdkmath.Int
func Int(r *rand.Rand) sdkmath.Int {
	return sdkmath.NewInt(r.Int63())
}
