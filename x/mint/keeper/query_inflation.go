package keeper

import (
	"context"

	"github.com/ignite/modules/x/mint/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (q queryServer) Inflation(ctx context.Context, req *types.QueryInflationRequest) (*types.QueryInflationResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	minter, err := q.k.Minter.Get(ctx)
	if err != nil {
		return nil, err
	}

	return &types.QueryInflationResponse{Inflation: minter.Inflation}, nil
}
