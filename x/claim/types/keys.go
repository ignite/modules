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

func KeyPrefix(p string) []byte {
	return []byte(p)
}

var (
	// ParamsKey is the prefix to retrieve all Params
	ParamsKey = collections.NewPrefix("p_claim")

	// MissionKey is the prefix to retrieve all Mission
	MissionKey = collections.NewPrefix("mission/value/")
	// MissionCountKey is the prefix to retrieve all MissionCount
	MissionCountKey = collections.NewPrefix("mission/count/")

	// ClaimRecordKey is the prefix to retrieve all ClaimRecord
	ClaimRecordKey = collections.NewPrefix("ClaimRecord/value/")

	// InitialClaimKey is the prefix to retrieve all InitialClaim
	InitialClaimKey = collections.NewPrefix("initialClaim/value/")

	// AirdropSupplyKey is the prefix to retrieve all AirdropSupply
	AirdropSupplyKey = collections.NewPrefix("airdropSupply/value/")
)
