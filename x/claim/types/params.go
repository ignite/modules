package types

import (
	"time"
)

// NewParams creates a new Params instance.
func NewParams(decayInformation DecayInformation, airdropStart time.Time) Params {
	return Params{DecayInformation: decayInformation, AirdropStart: airdropStart}
}

// DefaultParams returns a default set of parameters.
func DefaultParams() Params {
	return NewParams(
		NewDisabledDecay(),
		time.Unix(0, 0).UTC(),
	)
}

// Validate validates the set of params.
func (p Params) Validate() error {
	return validateDecayInformation(p.DecayInformation)
}

// validateDecayInformation validates the DecayInformation param
func validateDecayInformation(decayInfo DecayInformation) error {
	return decayInfo.Validate()
}
