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
	// ParamsKey is the prefix to retrieve all Params
	ParamsKey = collections.NewPrefix("p_mint")

	// MinterKey is the prefix to retrieve all Minter
	MinterKey = collections.NewPrefix("minter/value/")
)
