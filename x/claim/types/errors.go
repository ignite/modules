package types

// DONTCOVER

import (
	sdkerrors "cosmossdk.io/errors"
)

// x/claim module sentinel errors
var (
	ErrInvalidSigner          = sdkerrors.Register(ModuleName, 1101, "expected gov account as only signer for proposal message")
	ErrMissionNotFound        = sdkerrors.Register(ModuleName, 1102, "mission not found")
	ErrClaimRecordNotFound    = sdkerrors.Register(ModuleName, 1103, "claim record not found")
	ErrMissionCompleted       = sdkerrors.Register(ModuleName, 1104, "mission already completed")
	ErrAirdropSupplyNotFound  = sdkerrors.Register(ModuleName, 1105, "airdrop supply not found")
	ErrInitialClaimNotFound   = sdkerrors.Register(ModuleName, 1106, "initial claim information not found")
	ErrInitialClaimNotEnabled = sdkerrors.Register(ModuleName, 1107, "initial claim not enabled")
	ErrMissionCompleteFailure = sdkerrors.Register(ModuleName, 1108, "mission failed to complete")
	ErrNoClaimable            = sdkerrors.Register(ModuleName, 1109, "no amount to be claimed")
	ErrMissionNotCompleted    = sdkerrors.Register(ModuleName, 1110, "mission not completed yet")
	ErrMissionAlreadyClaimed  = sdkerrors.Register(ModuleName, 1111, "mission already claimed")
	ErrAirdropStartNotReached = sdkerrors.Register(ModuleName, 1112, "airdrop start has not been reached yet")
)
