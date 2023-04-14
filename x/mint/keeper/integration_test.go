package keeper_test

import (
	"encoding/json"

	abci "github.com/cometbft/cometbft/abci/types"
	simtestutil "github.com/cosmos/cosmos-sdk/testutil/sims"

	testapp "github.com/ignite/modules/app"
	"github.com/ignite/modules/testutil"
)

func setup(isCheckTx bool) *testapp.App {
	chainID := "simapp-chain-id"
	app, genesisState := testutil.GenApp(chainID, !isCheckTx, 5)
	if !isCheckTx {
		// init chain must be called to stop deliverState from being nil
		stateBytes, err := json.MarshalIndent(genesisState, "", " ")
		if err != nil {
			panic(err)
		}

		// Initialize the chain
		app.InitChain(
			abci.RequestInitChain{
				Validators:      []abci.ValidatorUpdate{},
				ConsensusParams: simtestutil.DefaultConsensusParams,
				AppStateBytes:   stateBytes,
				ChainId:         chainID,
			},
		)
	}

	return app
}
