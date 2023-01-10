package simulation

import (
	"fmt"
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

		var (
			mission    types.Mission
			missions   = k.GetAllMission(ctx)
			hasMission = false
		)
		for _, m := range missions {
			if cr.IsMissionCompleted(m.MissionID) && !cr.IsMissionClaimed(m.MissionID) {
				hasMission = true
				mission = m
			}
		}
		if !hasMission {
			return simtypes.NoOpMsg(
				types.ModuleName,
				msg.Type(),
				fmt.Sprintf("%s don't have mission to claim", simAccount.Address.String()),
			), nil, nil
		}

		// verify that there is claimable amount
		airdropSupply, _ := k.GetAirdropSupply(ctx)
		claimableAmount := cr.ClaimableFromMission(mission)
		claimable := sdk.NewCoins(sdk.NewCoin(airdropSupply.Denom, claimableAmount))
		// calculate claimable after decay factor
		decayInfo := k.GetParams(ctx).DecayInformation
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
