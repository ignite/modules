package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/ignite/modules/testutil/nullify"
	"github.com/ignite/modules/x/claim/types"
)

func TestInitialClaimQuery(t *testing.T) {
	ctx, tk := createClaimKeeper(t)
	item := createTestInitialClaim(tk, ctx)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetInitialClaimRequest
		response *types.QueryGetInitialClaimResponse
		err      error
	}{
		{
			desc:     "should allow valid query",
			request:  &types.QueryGetInitialClaimRequest{},
			response: &types.QueryGetInitialClaimResponse{InitialClaim: item},
		},
		{
			desc: "should return InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := tk.InitialClaim(ctx, tc.request)
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
