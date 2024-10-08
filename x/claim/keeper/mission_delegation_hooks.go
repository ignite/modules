package keeper

import (
	"context"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

type MissionDelegationHooks struct {
	k         Keeper
	missionID uint64
}

// NewMissionDelegationHooks returns a StakingHooks that triggers mission completion on delegation for an account
func (k Keeper) NewMissionDelegationHooks(missionID uint64) MissionDelegationHooks {
	return MissionDelegationHooks{k, missionID}
}

var _ stakingtypes.StakingHooks = MissionDelegationHooks{}

// BeforeDelegationCreated completes mission when a delegation is performed
func (h MissionDelegationHooks) BeforeDelegationCreated(ctx context.Context, delAddr sdk.AccAddress, _ sdk.ValAddress) error {
	_, err := h.k.CompleteMission(ctx, h.missionID, delAddr.String())
	return err
}

// AfterUnbondingInitiated implements StakingHooks
func (h MissionDelegationHooks) AfterUnbondingInitiated(context.Context, uint64) error {
	return nil
}

// AfterValidatorCreated implements StakingHooks
func (h MissionDelegationHooks) AfterValidatorCreated(_ context.Context, _ sdk.ValAddress) error {
	return nil
}

// AfterValidatorRemoved implements StakingHooks
func (h MissionDelegationHooks) AfterValidatorRemoved(_ context.Context, _ sdk.ConsAddress, _ sdk.ValAddress) error {
	return nil
}

// BeforeDelegationSharesModified implements StakingHooks
func (h MissionDelegationHooks) BeforeDelegationSharesModified(_ context.Context, _ sdk.AccAddress, _ sdk.ValAddress) error {
	return nil
}

// AfterDelegationModified implements StakingHooks
func (h MissionDelegationHooks) AfterDelegationModified(_ context.Context, _ sdk.AccAddress, _ sdk.ValAddress) error {
	return nil
}

// BeforeValidatorSlashed implements StakingHooks
func (h MissionDelegationHooks) BeforeValidatorSlashed(_ context.Context, _ sdk.ValAddress, _ sdkmath.LegacyDec) error {
	return nil
}

// BeforeValidatorModified implements StakingHooks
func (h MissionDelegationHooks) BeforeValidatorModified(_ context.Context, _ sdk.ValAddress) error {
	return nil
}

// AfterValidatorBonded implements StakingHooks
func (h MissionDelegationHooks) AfterValidatorBonded(_ context.Context, _ sdk.ConsAddress, _ sdk.ValAddress) error {
	return nil
}

// AfterValidatorBeginUnbonding implements StakingHooks
func (h MissionDelegationHooks) AfterValidatorBeginUnbonding(_ context.Context, _ sdk.ConsAddress, _ sdk.ValAddress) error {
	return nil
}

// BeforeDelegationRemoved implements StakingHooks
func (h MissionDelegationHooks) BeforeDelegationRemoved(_ context.Context, _ sdk.AccAddress, _ sdk.ValAddress) error {
	return nil
}
