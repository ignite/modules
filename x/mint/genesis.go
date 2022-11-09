package mint

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ignite/modules/x/mint/keeper"
	"github.com/ignite/modules/x/mint/types"
)

// InitGenesis new mint genesis
func InitGenesis(ctx sdk.Context, k keeper.Keeper, ak types.AccountKeeper, genState *types.GenesisState) {
	k.SetMinter(ctx, genState.Minter)
	if err := k.SetParams(ctx, genState.Params); err != nil {
		panic(err)
	}

	ak.GetModuleAccount(ctx, types.ModuleName)
}

// ExportGenesis returns a GenesisState for a given context and keeper.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()

	genesis.Minter = k.GetMinter(ctx)
	genesis.Params = k.GetParams(ctx)

	return genesis
}
