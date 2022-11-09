package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simtestutil "github.com/cosmos/cosmos-sdk/testutil/sims"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module/testutil"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	sdksimulation "github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/ignite/modules/testutil/simulation"
	"github.com/ignite/modules/x/claim/keeper"
	"github.com/ignite/modules/x/claim/types"
)

func SimulateMsgClaimInitial(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		msg := &types.MsgClaimInitial{}

		// find an account
		simAccount, _ := simtypes.RandomAcc(r, accs)

		// check the account has a claim record and initial claim has not been completed
		cr, found := k.GetClaimRecord(ctx, simAccount.Address.String())
		if !found {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "account has no claim record"), nil, nil
		}

		// verify if initial claim mission is completed
		if cr.IsMissionCompleted(0) {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "account already completed initial claim"), nil, nil
		}

		// verify that there is claimable amount
		m, _ := k.GetMission(ctx, 0)
		airdropSupply, _ := k.GetAirdropSupply(ctx)
		claimableAmount := cr.ClaimableFromMission(m)
		claimable := sdk.NewCoins(sdk.NewCoin(airdropSupply.Denom, claimableAmount))
		// calculate claimable after decay factor
		decayInfo := k.GetParams(ctx).DecayInformation
		claimable = decayInfo.ApplyDecayFactor(claimable, ctx.BlockTime())

		// check final claimable non-zero
		if claimable.Empty() {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), types.ErrNoClaimable.Error()), nil, nil
		}

		// initialize basic message
		msg = &types.MsgClaimInitial{
			Claimer: simAccount.Address.String(),
		}

		txCtx := sdksimulation.OperationInput{
			R:               r,
			App:             app,
			TxGen:           testutil.MakeTestEncodingConfig().TxConfig,
			Cdc:             nil,
			Msg:             msg,
			MsgType:         msg.Type(),
			Context:         ctx,
			SimAccount:      simAccount,
			AccountKeeper:   ak,
			Bankkeeper:      bk,
			ModuleName:      types.ModuleName,
			CoinsSpentInMsg: sdk.NewCoins(),
		}

		return simulation.GenAndDeliverTxWithRandFees(txCtx, simtestutil.DefaultGenTxGas)
	}
}
