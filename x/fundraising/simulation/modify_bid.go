package simulation

import (
	"fmt"
	"math/rand"

	"cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/ignite/modules/x/fundraising/keeper"
	"github.com/ignite/modules/x/fundraising/types"
)

func SimulateMsgModifyBid(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
	txGen client.TxConfig,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		msg := &types.MsgModifyBid{}
		auctions, err := k.Auctions(ctx)
		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "failed to get auctions"), nil, nil
		}
		if len(auctions) == 0 {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "no auction to modify a bid"), nil, nil
		}

		// Select a random auction
		auction := auctions[r.Intn(len(auctions))]
		if auction.GetType() != types.AuctionTypeBatch || auction.GetStatus() != types.AuctionStatusStarted {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), fmt.Sprintf("incorrect auction type or status %v", auction)), nil, nil
		}

		bids, err := k.GetBidsByAuctionID(ctx, auction.GetId())
		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "failed to get bids"), nil, nil
		}
		if len(bids) == 0 {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "no bid modify"), nil, nil
		}

		// Select a random bid
		bid := bids[r.Intn(len(bids))]
		simAccount, _ := FindAccount(accs, bid.Bidder)

		account := ak.GetAccount(ctx, simAccount.Address)
		spendable := bk.SpendableCoins(ctx, account.GetAddress())

		msg = types.NewMsgModifyBid(
			bid.AuctionId,
			account.GetAddress().String(),
			bid.BidId,
			bid.Price,
			bid.Coin.AddAmount(math.OneInt()),
		)

		txCtx := simulation.OperationInput{
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
			CoinsSpentInMsg: spendable,
		}

		return simulation.GenAndDeliverTxWithRandFees(txCtx)
	}
}
