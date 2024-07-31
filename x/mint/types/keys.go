package types

import "cosmossdk.io/collections"

const (
	// ModuleName defines the module name
	ModuleName = "mint"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_mint"
)

var (
	ParamsKey = collections.NewPrefix("p_mint")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

var (
	MinterKey = collections.NewPrefix("minter/value/")
)
