package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ignite/modules/pkg/errors"
	"github.com/ignite/modules/x/claim/types"
)

// Claim claims the Airdrop by the mission id if available and reach the airdrop start time
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

	// check if the claim is an initial claim
	initialClaim, found := k.GetInitialClaim(ctx)
	if found {
		if initialClaim.MissionID == msg.MissionID {
			if !initialClaim.Enabled {
				return nil, types.ErrInitialClaimNotEnabled
			}
			// if is an initial claim, automatically add to completed missions
			// the `ClaimMission` will update the claim record later
			claimRecord.CompletedMissions = append(claimRecord.CompletedMissions, msg.MissionID)
		}
	}

	// check if airdrop start time already reached
	params := k.GetParams(ctx)
	if ctx.BlockTime().Before(params.AirdropStart) {
		return &types.MsgClaimResponse{}, errors.Wrapf(
			types.ErrAirdropStartNotReached,
			"airdrop start not reached: %s",
			params.AirdropStart.String(),
		)
	}
	claimed, err := k.ClaimMission(ctx, claimRecord, msg.MissionID)
	if err != nil {
		return nil, err
	}

	return &types.MsgClaimResponse{
		Claimed: claimed,
	}, nil
}
