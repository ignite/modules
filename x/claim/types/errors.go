package types

// DONTCOVER

import (
	"github.com/ignite/modules/pkg/errors"
)

// x/claim module sentinel errors
var (
	ErrMissionNotFound        = errors.Register(ModuleName, 2, "mission not found")
	ErrClaimRecordNotFound    = errors.Register(ModuleName, 3, "claim record not found")
	ErrMissionCompleted       = errors.Register(ModuleName, 4, "mission already completed")
	ErrAirdropSupplyNotFound  = errors.Register(ModuleName, 5, "airdrop supply not found")
	ErrInitialClaimNotFound   = errors.Register(ModuleName, 6, "initial claim information not found")
	ErrInitialClaimNotEnabled = errors.Register(ModuleName, 7, "initial claim not enabled")
	ErrMissionCompleteFailure = errors.Register(ModuleName, 8, "mission failed to complete")
	ErrNoClaimable            = errors.Register(ModuleName, 9, "no amount to be claimed")
	ErrMissionNotCompleted    = errors.Register(ModuleName, 10, "mission not completed yet")
	ErrMissionAlreadyClaimed  = errors.Register(ModuleName, 11, "mission already claimed")
	ErrAirdropStartNotReached = errors.Register(ModuleName, 12, "Airdrop start has not been reached yet")
)
