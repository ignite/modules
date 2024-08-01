package keeper

import (
	"context"

	"cosmossdk.io/collections"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ignite/modules/pkg/errors"
	"github.com/ignite/modules/x/claim/types"
)

func (k msgServer) Claim(ctx context.Context, msg *types.MsgClaim) (*types.MsgClaimResponse, error) {
	// retrieve claim record of the user
	claimRecord, err := k.ClaimRecord.Get(ctx, msg.Claimer)
	if err != nil {
		return &types.MsgClaimResponse{}, errors.Wrapf(
			types.ErrClaimRecordNotFound,
			"claim record not found for address %s: %s",
			msg.Claimer,
			err.Error(),
		)
	}

	// check if the claim is an initial claim
	initialClaim, err := k.InitialClaim.Get(ctx)
	if err != nil && errors.IsOf(err, collections.ErrNotFound) {
		return nil, err
	} else if err == nil {
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
	params, err := k.Params.Get(ctx)
	if err != nil {
		return nil, err
	}

	airdropStart := params.AirdropStart
	blockTime := sdk.UnwrapSDKContext(ctx).BlockTime()
	if blockTime.Before(airdropStart) {
		return &types.MsgClaimResponse{}, errors.Wrapf(
			types.ErrAirdropStartNotReached,
			"airdrop start not reached: %s",
			airdropStart.String(),
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
