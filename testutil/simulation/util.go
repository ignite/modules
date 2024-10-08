package simulation

import (
	simtestutil "github.com/cosmos/cosmos-sdk/testutil/sims"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	sdksimulation "github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/ignite/modules/testutil/sample"
)

// GenAndDeliverTxWithRandFees generates a transaction with a random fee and delivers it.
func GenAndDeliverTxWithRandFees(txCtx sdksimulation.OperationInput, gas uint64) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
	account := txCtx.AccountKeeper.GetAccount(txCtx.Context, txCtx.SimAccount.Address)
	spendable := txCtx.Bankkeeper.SpendableCoins(txCtx.Context, account.GetAddress())

	var fees sdk.Coins
	var err error

	coins, hasNeg := spendable.SafeSub(txCtx.CoinsSpentInMsg...)
	if hasNeg {
		return simtypes.NoOpMsg(txCtx.ModuleName, sdk.MsgTypeURL(txCtx.Msg), "message doesn't leave room for fees"), nil, err
	}

	fees, err = sample.Fees(txCtx.R, coins)
	if err != nil {
		return simtypes.NoOpMsg(txCtx.ModuleName, sdk.MsgTypeURL(txCtx.Msg), "unable to generate fees"), nil, err
	}
	return GenAndDeliverTx(txCtx, fees, gas)
}

// GenAndDeliverTx generates a transactions and delivers it.
func GenAndDeliverTx(txCtx sdksimulation.OperationInput, fees sdk.Coins, gas uint64) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
	account := txCtx.AccountKeeper.GetAccount(txCtx.Context, txCtx.SimAccount.Address)
	tx, err := simtestutil.GenSignedMockTx(
		txCtx.R,
		txCtx.TxGen,
		[]sdk.Msg{txCtx.Msg},
		fees,
		gas,
		txCtx.Context.ChainID(),
		[]uint64{account.GetAccountNumber()},
		[]uint64{account.GetSequence()},
		txCtx.SimAccount.PrivKey,
	)
	if err != nil {
		return simtypes.NoOpMsg(txCtx.ModuleName, sdk.MsgTypeURL(txCtx.Msg), "unable to generate mock tx"), nil, err
	}

	_, _, err = txCtx.App.SimDeliver(txCtx.TxGen.TxEncoder(), tx)
	if err != nil {
		return simtypes.NoOpMsg(txCtx.ModuleName, sdk.MsgTypeURL(txCtx.Msg), "unable to deliver tx"), nil, err
	}

	return simtypes.NewOperationMsg(txCtx.Msg, true, ""), nil, nil
}
