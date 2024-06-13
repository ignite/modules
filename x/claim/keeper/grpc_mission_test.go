package keeper_test

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/ignite/modules/pkg/errors"
	"github.com/ignite/modules/testutil/nullify"
	"github.com/ignite/modules/x/claim/types"
)

func TestMissionQuerySingle(t *testing.T) {
	ctx, tk := createClaimKeeper(t)

	msgs := createNMission(tk, ctx, 2)
	for _, tc := range []struct {
		name     string
		request  *types.QueryGetMissionRequest
		response *types.QueryGetMissionResponse
		err      error
	}{
		{
			name:     "should allow valid query 1",
			request:  &types.QueryGetMissionRequest{MissionID: msgs[0].MissionID},
			response: &types.QueryGetMissionResponse{Mission: msgs[0]},
		},
		{
			name:     "should allow valid query 2",
			request:  &types.QueryGetMissionRequest{MissionID: msgs[1].MissionID},
			response: &types.QueryGetMissionResponse{Mission: msgs[1]},
		},
		{
			name:    "should return KeyNotFound",
			request: &types.QueryGetMissionRequest{MissionID: uint64(len(msgs))},
			err:     errors.ErrKeyNotFound,
		},
		{
			name: "should return InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			response, err := tk.Mission(ctx, tc.request)
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

func TestMissionQueryPaginated(t *testing.T) {
	ctx, tk := createClaimKeeper(t)

	msgs := createNMission(tk, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllMissionRequest {
		return &types.QueryAllMissionRequest{
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
			resp, err := tk.MissionAll(ctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Mission), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Mission),
			)
		}
	})
	t.Run("should paginate by key", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := tk.MissionAll(ctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Mission), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Mission),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("should paginate all", func(t *testing.T) {
		resp, err := tk.MissionAll(ctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.Mission),
		)
	})
	t.Run("should return InvalidRequest", func(t *testing.T) {
		_, err := tk.MissionAll(ctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
