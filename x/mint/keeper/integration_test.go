package keeper_test

import (
	"encoding/json"

	simtestutil "github.com/cosmos/cosmos-sdk/testutil/sims"
	abci "github.com/tendermint/tendermint/abci/types"

	testapp "github.com/ignite/modules/app"
	"github.com/ignite/modules/testutil"
)

func setup(isCheckTx bool) *testapp.App {
	app, genesisState := testutil.GenApp(!isCheckTx)
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
			},
		)
	}

	return app
}
