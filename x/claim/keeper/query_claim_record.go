package keeper

import (
	"context"
	"errors"

	"cosmossdk.io/collections"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/ignite/modules/x/claim/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (q queryServer) ListClaimRecord(ctx context.Context, req *types.QueryAllClaimRecordRequest) (*types.QueryAllClaimRecordResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	claimRecords, pageRes, err := query.CollectionPaginate(
		ctx,
		q.k.ClaimRecord,
		req.Pagination,
		func(_ string, value types.ClaimRecord) (types.ClaimRecord, error) {
			return value, nil
		},
	)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllClaimRecordResponse{ClaimRecord: claimRecords, Pagination: pageRes}, nil
}

func (q queryServer) GetClaimRecord(ctx context.Context, req *types.QueryGetClaimRecordRequest) (*types.QueryGetClaimRecordResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	val, err := q.k.ClaimRecord.Get(ctx, req.Address)
	if err != nil {
		if errors.Is(err, collections.ErrNotFound) {
			return nil, status.Error(codes.NotFound, "not found")
		}

		return nil, status.Error(codes.Internal, "internal error")
	}

	return &types.QueryGetClaimRecordResponse{ClaimRecord: val}, nil
}
