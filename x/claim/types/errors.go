package types

// DONTCOVER

import (
	"github.com/ignite/modules/pkg/errors"
)

// x/claim module sentinel errors
var (
	ErrInvalidSigner          = errors.Register(ModuleName, 1101, "expected gov account as only signer for proposal message")
	ErrMissionNotFound        = errors.Register(ModuleName, 1102, "mission not found")
	ErrClaimRecordNotFound    = errors.Register(ModuleName, 1103, "claim record not found")
	ErrMissionCompleted       = errors.Register(ModuleName, 1104, "mission already completed")
	ErrAirdropSupplyNotFound  = errors.Register(ModuleName, 1105, "airdrop supply not found")
	ErrInitialClaimNotFound   = errors.Register(ModuleName, 1106, "initial claim information not found")
	ErrInitialClaimNotEnabled = errors.Register(ModuleName, 1107, "initial claim not enabled")
	ErrMissionCompleteFailure = errors.Register(ModuleName, 1108, "mission failed to complete")
	ErrNoClaimable            = errors.Register(ModuleName, 1109, "no amount to be claimed")
	ErrMissionNotCompleted    = errors.Register(ModuleName, 1110, "mission not completed yet")
	ErrMissionAlreadyClaimed  = errors.Register(ModuleName, 1111, "mission already claimed")
	ErrAirdropStartNotReached = errors.Register(ModuleName, 1112, "airdrop start has not been reached yet")
)
