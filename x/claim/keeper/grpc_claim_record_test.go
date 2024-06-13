package keeper_test

import (
	"strconv"
	"testing"

	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/ignite/modules/testutil/nullify"
	"github.com/ignite/modules/x/claim/types"
)

func TestClaimRecordQuerySingle(t *testing.T) {
	ctx, tk := createClaimKeeper(t)
	msgs := createNClaimRecord(tk, ctx, 2)

	for _, tc := range []struct {
		name     string
		request  *types.QueryGetClaimRecordRequest
		response *types.QueryGetClaimRecordResponse
		err      error
	}{
		{
			name: "should allow valid query 1",
			request: &types.QueryGetClaimRecordRequest{
				Address: msgs[0].Address,
			},
			response: &types.QueryGetClaimRecordResponse{ClaimRecord: msgs[0]},
		},
		{
			name: "should allow valid query 2",
			request: &types.QueryGetClaimRecordRequest{
				Address: msgs[1].Address,
			},
			response: &types.QueryGetClaimRecordResponse{ClaimRecord: msgs[1]},
		},
		{
			name: "should return KeyNotFound",
			request: &types.QueryGetClaimRecordRequest{
				Address: strconv.Itoa(100000),
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			name: "should return InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			response, err := tk.ClaimRecord(ctx, tc.request)
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

func TestClaimRecordQueryPaginated(t *testing.T) {
	ctx, tk := createClaimKeeper(t)
	msgs := createNClaimRecord(tk, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllClaimRecordRequest {
		return &types.QueryAllClaimRecordRequest{
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
		}
	}
	t.Run("should paginate by offset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(msgs); i += step {
			resp, err := tk.ClaimRecordAll(ctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.ClaimRecord), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.ClaimRecord),
			)
		}
	})
	t.Run("should paginate by key", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := tk.ClaimRecordAll(ctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.ClaimRecord), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.ClaimRecord),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("should paginate all", func(t *testing.T) {
		resp, err := tk.ClaimRecordAll(ctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.ClaimRecord),
		)
	})
	t.Run("should return InvalidRequest", func(t *testing.T) {
		_, err := tk.ClaimRecordAll(ctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
