package keeper_test

import (
	"context"
	"strconv"
	"testing"

	sdkmath "cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "github.com/ignite/modules/testutil/keeper"
	"github.com/ignite/modules/testutil/nullify"
	"github.com/ignite/modules/x/claim/keeper"
	"github.com/ignite/modules/x/claim/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNClaimRecord(keeper keeper.Keeper, ctx context.Context, n int) []types.ClaimRecord {
	items := make([]types.ClaimRecord, n)
	for i := range items {
		items[i].Address = strconv.Itoa(i)
		items[i].Claimable = sdkmath.ZeroInt()
		_ = keeper.ClaimRecord.Set(ctx, items[i].Address, items[i])
	}
	return items
}

func TestClaimRecordQuerySingle(t *testing.T) {
	k, ctx, _ := keepertest.ClaimKeeper(t)
	qs := keeper.NewQueryServerImpl(k)
	msgs := createNClaimRecord(k, ctx, 2)
	tests := []struct {
		desc     string
		request  *types.QueryGetClaimRecordRequest
		response *types.QueryGetClaimRecordResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetClaimRecordRequest{
				Address: msgs[0].Address,
			},
			response: &types.QueryGetClaimRecordResponse{ClaimRecord: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetClaimRecordRequest{
				Address: msgs[1].Address,
			},
			response: &types.QueryGetClaimRecordResponse{ClaimRecord: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetClaimRecordRequest{
				Address: strconv.Itoa(100000),
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := qs.GetClaimRecord(ctx, tc.request)
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
	k, ctx, _ := keepertest.ClaimKeeper(t)
	qs := keeper.NewQueryServerImpl(k)
	msgs := createNClaimRecord(k, ctx, 5)

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
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(msgs); i += step {
			resp, err := qs.ListClaimRecord(ctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.ClaimRecord), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.ClaimRecord),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := qs.ListClaimRecord(ctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.ClaimRecord), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.ClaimRecord),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := qs.ListClaimRecord(ctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.ClaimRecord),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := qs.ListClaimRecord(ctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
