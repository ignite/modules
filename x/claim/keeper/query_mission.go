package keeper

import (
	"context"
	"errors"

	"cosmossdk.io/collections"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/ignite/modules/x/claim/types"
)

func (q queryServer) ListMission(ctx context.Context, req *types.QueryAllMissionRequest) (*types.QueryAllMissionResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	missions, pageRes, err := query.CollectionPaginate(
		ctx,
		q.k.Mission,
		req.Pagination,
		func(_ uint64, value types.Mission) (types.Mission, error) {
			return value, nil
		},
	)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllMissionResponse{Mission: missions, Pagination: pageRes}, nil
}

func (q queryServer) GetMission(ctx context.Context, req *types.QueryGetMissionRequest) (*types.QueryGetMissionResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	mission, err := q.k.Mission.Get(ctx, req.MissionID)
	if err != nil {
		if errors.Is(err, collections.ErrNotFound) {
			return nil, sdkerrors.ErrKeyNotFound
		}

		return nil, status.Error(codes.Internal, "internal error")
	}

	return &types.QueryGetMissionResponse{Mission: mission}, nil
}
