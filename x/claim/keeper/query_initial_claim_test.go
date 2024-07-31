package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "github.com/ignite/modules/testutil/keeper"
	"github.com/ignite/modules/testutil/nullify"
	"github.com/ignite/modules/x/claim/keeper"
	"github.com/ignite/modules/x/claim/types"
)

func TestInitialClaimQuery(t *testing.T) {
	k, ctx, _ := keepertest.ClaimKeeper(t)
	qs := keeper.NewQueryServerImpl(k)
	item := types.InitialClaim{}
	err := k.InitialClaim.Set(ctx, item)
	require.NoError(t, err)

	tests := []struct {
		desc     string
		request  *types.QueryGetInitialClaimRequest
		response *types.QueryGetInitialClaimResponse
		err      error
	}{
		{
			desc:     "First",
			request:  &types.QueryGetInitialClaimRequest{},
			response: &types.QueryGetInitialClaimResponse{InitialClaim: item},
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := qs.GetInitialClaim(ctx, tc.request)
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
