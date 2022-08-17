package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"

	testkeeper "github.com/ignite/modules/testutil/keeper"
	"github.com/ignite/modules/x/claim/keeper"
	"github.com/ignite/modules/x/claim/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)

	return keeper.NewMsgServerImpl(*tk.ClaimKeeper), sdk.WrapSDKContext(ctx)
}
