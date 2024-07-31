package types

import (
	"context"
	"cosmossdk.io/math"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// AccountKeeper defines the expected account keeper used for simulations (noalias)
type AccountKeeper interface {
	GetModuleAddress(name string) sdk.AccAddress
	GetAccount(ctx context.Context, addr sdk.AccAddress) sdk.AccountI
}

// BankKeeper defines the expected interface needed to retrieve account balances.
type BankKeeper interface {
	SpendableCoins(ctx context.Context, addr sdk.AccAddress) sdk.Coins
	GetBalance(ctx context.Context, addr sdk.AccAddress, denom string) sdk.Coin
	MintCoins(ctx context.Context, moduleName string, amt sdk.Coins) error
	BurnCoins(ctx context.Context, name string, amt sdk.Coins) error
	SendCoinsFromModuleToAccount(ctx context.Context, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins) error
}

// DistrKeeper defines the contract needed to be fulfilled for distribution keeper.
type DistrKeeper interface {
	FundCommunityPool(ctx context.Context, amount sdk.Coins, sender sdk.AccAddress) error
}

// ParamSubspace defines the expected Subspace interface for parameters.
type ParamSubspace interface {
	Get(context.Context, []byte, interface{})
	Set(context.Context, []byte, interface{})
}

// Event Hooks

// StakingHooks event hooks for staking validator object (noalias)
type StakingHooks interface {
	AfterValidatorCreated(ctx context.Context, valAddr sdk.ValAddress) error                           // Must be called when a validator is created
	BeforeValidatorModified(ctx context.Context, valAddr sdk.ValAddress) error                         // Must be called when a validator's state changes
	AfterValidatorRemoved(ctx context.Context, consAddr sdk.ConsAddress, valAddr sdk.ValAddress) error // Must be called when a validator is deleted

	AfterValidatorBonded(ctx context.Context, consAddr sdk.ConsAddress, valAddr sdk.ValAddress) error         // Must be called when a validator is bonded
	AfterValidatorBeginUnbonding(ctx context.Context, consAddr sdk.ConsAddress, valAddr sdk.ValAddress) error // Must be called when a validator begins unbonding

	BeforeDelegationCreated(ctx context.Context, delAddr sdk.AccAddress, valAddr sdk.ValAddress) error        // Must be called when a delegation is created
	BeforeDelegationSharesModified(ctx context.Context, delAddr sdk.AccAddress, valAddr sdk.ValAddress) error // Must be called when a delegation's shares are modified
	BeforeDelegationRemoved(ctx context.Context, delAddr sdk.AccAddress, valAddr sdk.ValAddress) error        // Must be called when a delegation is removed
	AfterDelegationModified(ctx context.Context, delAddr sdk.AccAddress, valAddr sdk.ValAddress) error
	BeforeValidatorSlashed(ctx context.Context, valAddr sdk.ValAddress, fraction math.LegacyDec) error
	AfterUnbondingInitiated(ctx context.Context, id uint64) error
}

// StakingHooksWrapper is a wrapper for modules to inject StakingHooks using depinject.
type StakingHooksWrapper struct{ StakingHooks }

// IsOnePerModuleType implements the depinject.OnePerModuleType interface.
func (StakingHooksWrapper) IsOnePerModuleType() {}

// GovHooks event hooks for governance proposal object (noalias)
type GovHooks interface {
	AfterProposalSubmission(ctx context.Context, proposalID uint64) error                            // Must be called after proposal is submitted
	AfterProposalDeposit(ctx context.Context, proposalID uint64, depositorAddr sdk.AccAddress) error // Must be called after a deposit is made
	AfterProposalVote(ctx context.Context, proposalID uint64, voterAddr sdk.AccAddress) error        // Must be called after a vote on a proposal is cast
	AfterProposalFailedMinDeposit(ctx context.Context, proposalID uint64) error                      // Must be called when proposal fails to reach min deposit
	AfterProposalVotingPeriodEnded(ctx context.Context, proposalID uint64) error                     // Must be called when proposal's finishes it's voting period
}

type GovHooksWrapper struct{ GovHooks }

// IsOnePerModuleType implements the depinject.OnePerModuleType interface.
func (GovHooksWrapper) IsOnePerModuleType() {}
