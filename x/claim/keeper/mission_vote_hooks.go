package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
)

type MissionVoteHooks struct {
	k         Keeper
	missionID uint64
}

// NewMissionVoteHooks returns a GovHooks that triggers mission completion on voting for a proposal
func (k Keeper) NewMissionVoteHooks(missionID uint64) MissionVoteHooks {
	return MissionVoteHooks{k, missionID}
}

var _ govtypes.GovHooks = MissionVoteHooks{}

// AfterProposalVote completes mission when a vote is cast
func (h MissionVoteHooks) AfterProposalVote(ctx context.Context, _ uint64, voterAddr sdk.AccAddress) error {
	_, err := h.k.CompleteMission(ctx, h.missionID, voterAddr.String())
	if err != nil {
		return err
	}

	return nil
}

// AfterProposalSubmission implements GovHooks
func (h MissionVoteHooks) AfterProposalSubmission(_ context.Context, _ uint64) error {
	return nil
}

// AfterProposalDeposit implements GovHooks
func (h MissionVoteHooks) AfterProposalDeposit(_ context.Context, _ uint64, _ sdk.AccAddress) error {
	return nil
}

// AfterProposalFailedMinDeposit implements GovHooks
func (h MissionVoteHooks) AfterProposalFailedMinDeposit(_ context.Context, _ uint64) error {
	return nil
}

// AfterProposalVotingPeriodEnded implements GovHooks
func (h MissionVoteHooks) AfterProposalVotingPeriodEnded(_ context.Context, _ uint64) error {
	return nil
}
