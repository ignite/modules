package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	errors "github.com/ignite/modules/pkg/errors"
	"github.com/ignite/modules/x/claim/types"
)

// SetMission set a specific mission in the store
func (k Keeper) SetMission(ctx sdk.Context, mission types.Mission) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MissionKey))
	b := k.cdc.MustMarshal(&mission)
	store.Set(types.GetMissionIDBytes(mission.MissionID), b)
}

// GetMission returns a mission from its id
func (k Keeper) GetMission(ctx sdk.Context, id uint64) (val types.Mission, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MissionKey))
	b := store.Get(types.GetMissionIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveMission removes a mission from the store
func (k Keeper) RemoveMission(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MissionKey))
	store.Delete(types.GetMissionIDBytes(id))
}

// GetAllMission returns all mission
func (k Keeper) GetAllMission(ctx sdk.Context) (list []types.Mission) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MissionKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Mission
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// CompleteMission save the completion of the mission
func (k Keeper) CompleteMission(ctx sdk.Context, missionID uint64, address string) error {
	// retrieve claim record of the user
	claimRecord, found := k.GetClaimRecord(ctx, address)
	if !found {
		return errors.Wrapf(types.ErrClaimRecordNotFound, "claim record not found for address %s", address)
	}

	// check if the mission is already completed for the claim record
	if claimRecord.IsMissionCompleted(missionID) || claimRecord.IsMissionClaimed(missionID) {
		return errors.Wrapf(
			types.ErrMissionCompleted,
			"mission %d completed for address %s",
			missionID,
			address,
		)
	}
	claimRecord.CompletedMissions = append(claimRecord.CompletedMissions, missionID)

	k.SetClaimRecord(ctx, claimRecord)

	err := ctx.EventManager().EmitTypedEvent(&types.EventMissionCompleted{
		MissionID: missionID,
		Address:   address,
	})
	if err != nil {
		return err
	}

	// try to claim mission
	return k.ClaimMission(ctx, missionID, address)
}

// ClaimMission distribute the claimable portion of airdrop to the user the method
// fails if the mission has already been claimed or not completed
func (k Keeper) ClaimMission(ctx sdk.Context, missionID uint64, address string) error {
	airdropStart := k.AirdropStart(ctx)
	if ctx.BlockTime().Before(airdropStart) {
		return errors.Wrapf(
			types.ErrAirdropStartNotReached,
			"airdrop start not reached: %s",
			airdropStart.String(),
		)
	}

	airdropSupply, found := k.GetAirdropSupply(ctx)
	if !found {
		return errors.Wrap(types.ErrAirdropSupplyNotFound, "airdrop supply is not defined")
	}

	// retrieve mission
	mission, found := k.GetMission(ctx, missionID)
	if !found {
		return errors.Wrapf(types.ErrMissionNotFound, "mission %d not found", missionID)
	}

	// retrieve claim record of the user
	claimRecord, found := k.GetClaimRecord(ctx, address)
	if !found {
		return errors.Wrapf(types.ErrClaimRecordNotFound, "claim record not found for address %s", address)
	}

	// check if the mission is not completed for the claim record
	if !claimRecord.IsMissionCompleted(missionID) {
		return errors.Wrapf(
			types.ErrMissionNotCompleted,
			"mission %d is not completed for address %s",
			missionID,
			address,
		)
	}
	if claimRecord.IsMissionClaimed(missionID) {
		return errors.Wrapf(
			types.ErrMissionNotCompleted,
			"mission %d is already claimed for address %s",
			missionID,
			address,
		)
	}
	claimRecord.ClaimedMissions = append(claimRecord.ClaimedMissions, missionID)

	// calculate claimable from mission weight and claim
	claimableAmount := claimRecord.ClaimableFromMission(mission)
	claimable := sdk.NewCoins(sdk.NewCoin(airdropSupply.Denom, claimableAmount))

	// calculate claimable after decay factor
	decayInfo := k.DecayInformation(ctx)
	claimable = decayInfo.ApplyDecayFactor(claimable, ctx.BlockTime())

	// check final claimable non-zero
	if claimable.Empty() {
		return types.ErrNoClaimable
	}

	// decrease airdrop supply
	airdropSupply.Amount = airdropSupply.Amount.Sub(claimable.AmountOf(airdropSupply.Denom))
	if airdropSupply.Amount.IsNegative() {
		return errors.Critical("airdrop supply is lower than total claimable")
	}

	// send claimable to the user
	claimer, err := sdk.AccAddressFromBech32(address)
	if err != nil {
		return errors.Criticalf("invalid claimer address %s", err.Error())
	}
	if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, claimer, claimable); err != nil {
		return errors.Criticalf("can't send claimable coins %s", err.Error())
	}

	// update store
	k.SetAirdropSupply(ctx, airdropSupply)
	k.SetClaimRecord(ctx, claimRecord)

	return ctx.EventManager().EmitTypedEvent(&types.EventMissionClaimed{
		MissionID: missionID,
		Claimer:   address,
	})
}
