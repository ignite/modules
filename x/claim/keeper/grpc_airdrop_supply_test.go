package keeper_test

import (
	"testing"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	testkeeper "github.com/ignite/modules/testutil/keeper"
	"github.com/ignite/modules/testutil/nullify"
	"github.com/ignite/modules/x/claim/types"
)

func TestAirdropSupplyQuery(t *testing.T) {
	var (
		ctx, tk, _   = testkeeper.NewTestSetup(t)
		wctx         = sdk.WrapSDKContext(ctx)
		sampleSupply = sdk.NewCoin("foo", sdkmath.NewInt(1000))
	)
	tk.ClaimKeeper.SetAirdropSupply(ctx, sampleSupply)

	for _, tc := range []struct {
		name     string
		request  *types.QueryGetAirdropSupplyRequest
		response *types.QueryGetAirdropSupplyResponse
		err      error
	}{
		{
			name:     "should allow valid query",
			request:  &types.QueryGetAirdropSupplyRequest{},
			response: &types.QueryGetAirdropSupplyResponse{AirdropSupply: sampleSupply},
		},
		{
			name: "should return InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			response, err := tk.ClaimKeeper.AirdropSupply(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}
