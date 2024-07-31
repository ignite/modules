package types

import "cosmossdk.io/collections"

const (
	// ModuleName defines the module name
	ModuleName = "claim"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_claim"
)

var (
	ParamsKey = collections.NewPrefix("p_claim")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

var (
	MissionKey      = collections.NewPrefix("mission/value/")
	MissionCountKey = collections.NewPrefix("mission/count/")
)

var (
	InitialClaimKey = collections.NewPrefix("initialClaim/value/")
)

var (
	AirdropSupplyKey = collections.NewPrefix("airdropSupply/value/")
)
