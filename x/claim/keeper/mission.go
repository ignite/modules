package keeper

import (
	"context"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ignite/modules/pkg/errors"
	"github.com/ignite/modules/x/claim/types"
)

// CompleteMission saves the completion of the mission. The claim will
// be called automatically if the airdrop start has already been reached.
// If not, it will only save the mission as completed.
func (k Keeper) CompleteMission(
	ctx context.Context,
	missionID uint64,
	address string,
) (claimed math.Int, err error) {
	// retrieve mission
	if _, err := k.Mission.Get(ctx, missionID); err != nil {
		return claimed, errors.Wrapf(types.ErrMissionNotFound, "mission %d not found: %s", missionID, err.Error())
	}

	// retrieve claim record of the user
	claimRecord, err := k.ClaimRecord.Get(ctx, address)
	if err != nil {
		return claimed, errors.Wrapf(types.ErrClaimRecordNotFound, "claim record not found for address %s: %s", address, err.Error())
	}

	// check if the mission is already completed for the claim record
	if claimRecord.IsMissionCompleted(missionID) {
		return claimed, errors.Wrapf(
			types.ErrMissionCompleted,
			"mission %d completed for address %s",
			missionID,
			address,
		)
	}
	claimRecord.CompletedMissions = append(claimRecord.CompletedMissions, missionID)

	if err := k.ClaimRecord.Set(ctx, claimRecord.Address, claimRecord); err != nil {
		return claimed, err
	}

	// try to claim the mission if airdrop start is reached
	blockTime := sdk.UnwrapSDKContext(ctx).BlockTime()
	params, err := k.Params.Get(ctx)
	if err != nil {
		return claimed, err
	}
	airdropStart := params.AirdropStart
	if blockTime.After(airdropStart) {
		return k.ClaimMission(ctx, claimRecord, missionID)
	}

	return claimed, nil
}

// ClaimMission distributes the claimable portion of the airdrop to the user
// the method fails if the mission has already been claimed or not completed
func (k Keeper) ClaimMission(
	ctx context.Context,
	claimRecord types.ClaimRecord,
	missionID uint64,
) (claimed math.Int, err error) {
	airdropSupply, err := k.AirdropSupply.Get(ctx)
	if err != nil {
		return claimed, errors.Wrapf(types.ErrAirdropSupplyNotFound, "airdrop supply is not defined: %s", err.Error())
	}

	// retrieve mission
	mission, err := k.Mission.Get(ctx, missionID)
	if err != nil {
		return claimed, errors.Wrapf(types.ErrMissionNotFound, "mission %d not found: %s", missionID, err.Error())
	}

	// check if the mission is not completed for the claim record
	if !claimRecord.IsMissionCompleted(missionID) {
		return claimed, errors.Wrapf(
			types.ErrMissionNotCompleted,
			"mission %d is not completed for address %s",
			missionID,
			claimRecord.Address,
		)
	}
	if claimRecord.IsMissionClaimed(missionID) {
		return claimed, errors.Wrapf(
			types.ErrMissionAlreadyClaimed,
			"mission %d is already claimed for address %s",
			missionID,
			claimRecord.Address,
		)
	}
	claimRecord.ClaimedMissions = append(claimRecord.ClaimedMissions, missionID)

	// calculate claimable from mission weight and claim
	claimableAmount := claimRecord.ClaimableFromMission(mission)
	claimable := sdk.NewCoins(sdk.NewCoin(airdropSupply.Supply.Denom, claimableAmount))

	// calculate claimable after decay factor
	params, err := k.Params.Get(ctx)
	if err != nil {
		return claimed, err
	}
	decayInfo := params.DecayInformation
	blockTime := sdk.UnwrapSDKContext(ctx).BlockTime()
	claimable = decayInfo.ApplyDecayFactor(claimable, blockTime)

	// check final claimable non-zero
	if claimable.Empty() {
		return claimed, types.ErrNoClaimable
	}

	// decrease airdrop supply
	claimed = claimable.AmountOf(airdropSupply.Supply.Denom)
	airdropSupply.Supply.Amount = airdropSupply.Supply.Amount.Sub(claimed)
	if airdropSupply.Supply.Amount.IsNegative() {
		return claimed, errors.Critical("airdrop supply is lower than total claimable")
	}

	// send claimable to the user
	claimer, err := sdk.AccAddressFromBech32(claimRecord.Address)
	if err != nil {
		return claimed, errors.Criticalf("invalid claimer address %s", err.Error())
	}
	if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, claimer, claimable); err != nil {
		return claimed, errors.Criticalf("can't send claimable coins %s", err.Error())
	}

	// update store
	if err := k.AirdropSupply.Set(ctx, airdropSupply); err != nil {
		return claimed, err
	}
	if err := k.ClaimRecord.Set(ctx, claimRecord.Address, claimRecord); err != nil {
		return claimed, err
	}

	return claimed, nil
}
