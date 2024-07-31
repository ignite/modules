package types

import (
	"fmt"
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
		time.Unix(0, 0),
	)
}

// Validate validates the set of params.
func (p Params) Validate() error {
	if err := validateDecayInformation(p.DecayInformation); err != nil {
		return err
	}
	if err := validateAirdropStart(p.AirdropStart); err != nil {
		return err
	}

	return nil
}

// validateDecayInformation validates the DecayInformation param
func validateDecayInformation(v interface{}) error {
	decayInfo, ok := v.(DecayInformation)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", v)
	}

	return decayInfo.Validate()
}

func validateAirdropStart(i interface{}) error {
	if _, ok := i.(time.Time); !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}
	return nil
}
