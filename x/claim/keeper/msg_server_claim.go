package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ignite/modules/pkg/errors"
	"github.com/ignite/modules/x/claim/types"
)

func (k msgServer) Claim(goCtx context.Context, msg *types.MsgClaim) (*types.MsgClaimResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// retrieve claim record of the user
	claimRecord, found := k.GetClaimRecord(ctx, msg.Claimer)
	if !found {
		return &types.MsgClaimResponse{}, errors.Wrapf(
			types.ErrClaimRecordNotFound,
			"claim record not found for address %s",
			msg.Claimer,
		)
	}

	// check if airdrop start time already reached
	airdropStart := k.AirdropStart(ctx)
	if ctx.BlockTime().Before(airdropStart) {
		return &types.MsgClaimResponse{}, errors.Wrapf(
			types.ErrAirdropStartNotReached,
			"airdrop start not reached: %s",
			airdropStart.String(),
		)
	}
	if err := k.ClaimMission(ctx, claimRecord, msg.MissionID); err != nil {
		return nil, err
	}

	return &types.MsgClaimResponse{}, nil
}
