package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ignite/modules/pkg/errors"
	"github.com/ignite/modules/x/claim/types"
)

func (k msgServer) Claim(goCtx context.Context, msg *types.MsgClaim) (*types.MsgClaimResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := k.ClaimMission(ctx, msg.MissionID, msg.Claimer); err != nil {
		return nil, errors.Wrap(types.ErrMissionCompleteFailure, err.Error())
	}
	return &types.MsgClaimResponse{}, nil
}
