package keeper

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/ignite/modules/x/mint/types"
)

func (q queryServer) AnnualProvisions(ctx context.Context, req *types.QueryAnnualProvisionsRequest) (*types.QueryAnnualProvisionsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	minter, err := q.k.Minter.Get(ctx)
	if err != nil {
		return nil, err
	}

	return &types.QueryAnnualProvisionsResponse{AnnualProvisions: minter.AnnualProvisions}, nil
}
