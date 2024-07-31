package simulation

import (
	"fmt"
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client"
	simtestutil "github.com/cosmos/cosmos-sdk/testutil/sims"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	sdksimulation "github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/ignite/modules/testutil/simulation"
	"github.com/ignite/modules/x/claim/keeper"
	"github.com/ignite/modules/x/claim/types"
)

var TypeMsgClaim = sdk.MsgTypeURL(&types.MsgClaim{})

func SimulateMsgClaim(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
	txGen client.TxConfig,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		msg := &types.MsgClaim{}

		// find an account
		simAccount, _ := simtypes.RandomAcc(r, accs)

		// check the account has a claim record and initial claim has not been completed
		cr, err := k.ClaimRecord.Get(ctx, simAccount.Address.String())
		if err != nil {
			return simtypes.NoOpMsg(
				types.ModuleName,
				TypeMsgClaim,
				fmt.Sprintf("account has no claim record: %s", err.Error()),
			), nil, nil
		}

		var (
			mission    types.Mission
			hasMission = false
		)
		missions, _, err := query.CollectionPaginate(
			ctx,
			k.Mission,
			nil,
			func(_ uint64, value types.Mission) (types.Mission, error) {
				return value, nil
			},
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
				TypeMsgClaim,
				fmt.Sprintf("%s don't have mission to claim", simAccount.Address.String()),
			), nil, nil
		}

		// verify that there is claimable amount
		airdropSupply, err := k.AirdropSupply.Get(ctx)
		if err != nil {
			return simtypes.NoOpMsg(
				types.ModuleName,
				TypeMsgClaim,
				"don't have airdrop supply",
			), nil, nil
		}
		claimableAmount := cr.ClaimableFromMission(mission)
		claimable := sdk.NewCoins(sdk.NewCoin(airdropSupply.Supply.Denom, claimableAmount))
		// calculate claimable after decay factor

		params, err := k.Params.Get(ctx)
		if err != nil {
			return simtypes.NoOpMsg(
				types.ModuleName,
				TypeMsgClaim,
				"don't have params",
			), nil, nil
		}
		decayInfo := params.DecayInformation
		claimable = decayInfo.ApplyDecayFactor(claimable, ctx.BlockTime())

		// check final claimable non-zero
		if claimable.Empty() {
			return simtypes.NoOpMsg(types.ModuleName, TypeMsgClaim, types.ErrNoClaimable.Error()), nil, nil
		}

		// initialize basic message
		msg = &types.MsgClaim{
			Claimer: simAccount.Address.String(),
		}

		txCtx := sdksimulation.OperationInput{
			R:               r,
			App:             app,
			TxGen:           txGen,
			Cdc:             nil,
			Msg:             msg,
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
