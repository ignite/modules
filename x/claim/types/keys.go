package types

const (
	// ModuleName defines the module name
	ModuleName = "claim"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_claim"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

const (
	MissionKey       = "Mission-value-"
	AirdropSupplyKey = "AirdropSupply-value-"
	InitialClaimKey  = "InitialClaim-value-"
)

var ParamsKey = []byte{0x02}
