package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/simapp/helpers"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	sdksimulation "github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/ignite/modules/testutil/simulation"
	"github.com/ignite/modules/x/claim/keeper"
	"github.com/ignite/modules/x/claim/types"
)

func SimulateMsgClaim(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		msg := &types.MsgClaim{}

		// find an account
		simAccount, _ := simtypes.RandomAcc(r, accs)

		// check the account has a claim record and initial claim has not been completed
		cr, found := k.GetClaimRecord(ctx, simAccount.Address.String())
		if !found {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "account has no claim record"), nil, nil
		}

		// verify if initial claim mission is completed
		if !cr.IsMissionCompleted(1) {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "account already completed claim"), nil, nil
		}

		// verify that there is claimable amount
		m, _ := k.GetMission(ctx, 1)
		airdropSupply, _ := k.GetAirdropSupply(ctx)
		claimableAmount := cr.ClaimableFromMission(m)
		claimable := sdk.NewCoins(sdk.NewCoin(airdropSupply.Denom, claimableAmount))
		// calculate claimable after decay factor
		decayInfo := k.DecayInformation(ctx)
		claimable = decayInfo.ApplyDecayFactor(claimable, ctx.BlockTime())

		// check final claimable non-zero
		if claimable.Empty() {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), types.ErrNoClaimable.Error()), nil, nil
		}

		// initialize basic message
		msg = &types.MsgClaim{
			Claimer: simAccount.Address.String(),
		}

		txCtx := sdksimulation.OperationInput{
			R:               r,
			App:             app,
			TxGen:           simappparams.MakeTestEncodingConfig().TxConfig,
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

		return simulation.GenAndDeliverTxWithRandFees(txCtx, helpers.DefaultGenTxGas)
	}
}
