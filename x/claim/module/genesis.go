package claim

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ignite/modules/x/claim/keeper"
	"github.com/ignite/modules/x/claim/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the claimRecord
	for _, elem := range genState.ClaimRecordList {
		if err := k.ClaimRecord.Set(ctx, elem.Address, elem); err != nil {
			panic(err)
		}
	}
	// Set all the mission
	for _, elem := range genState.MissionList {
		if err := k.Mission.Set(ctx, elem.MissionId, elem); err != nil {
			panic(err)
		}
	}

	// Set mission count
	if err := k.MissionSeq.Set(ctx, genState.MissionCount); err != nil {
		panic(err)
	}
	// Set if defined
	if err := k.InitialClaim.Set(ctx, genState.InitialClaim); err != nil {
		panic(err)
	}
	// Set if defined
	if err := k.AirdropSupply.Set(ctx, genState.AirdropSupply); err != nil {
		panic(err)
	}
	// this line is used by starport scaffolding # genesis/module/init
	if err := k.Params.Set(ctx, genState.Params); err != nil {
		panic(err)
	}
}

// ExportGenesis returns the module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	var err error

	genesis := types.DefaultGenesis()
	genesis.Params, err = k.Params.Get(ctx)
	if err != nil {
		panic(err)
	}

	if err := k.ClaimRecord.Walk(ctx, nil, func(_ string, val types.ClaimRecord) (stop bool, err error) {
		genesis.ClaimRecordList = append(genesis.ClaimRecordList, val)
		return false, nil
	}); err != nil {
		panic(err)
	}

	err = k.Mission.Walk(ctx, nil, func(key uint64, elem types.Mission) (bool, error) {
		genesis.MissionList = append(genesis.MissionList, elem)
		return false, nil
	})
	if err != nil {
		panic(err)
	}

	genesis.MissionCount, err = k.MissionSeq.Peek(ctx)
	if err != nil {
		panic(err)
	}

	// Get all initialClaim
	initialClaim, err := k.InitialClaim.Get(ctx)
	if err == nil {
		genesis.InitialClaim = initialClaim
	}

	// Get all airdropSupply
	airdropSupply, err := k.AirdropSupply.Get(ctx)
	if err == nil {
		genesis.AirdropSupply = airdropSupply
	}

	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
