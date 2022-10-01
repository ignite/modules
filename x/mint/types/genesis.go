package types

// NewGenesisState creates a new GenesisState object
func NewGenesisState(minter Minter, params Params) *GenesisState {
	return &GenesisState{
		Minter: minter,
		Params: params,
	}
}

// DefaultGenesis creates a default GenesisState object
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		Minter: DefaultInitialMinter(),
		Params: DefaultParams(),
	}
}

// Validate validates the provided genesis state to ensure the
// expected invariants holds.
func (gs GenesisState) Validate() error {
	if err := gs.Params.Validate(); err != nil {
		return err
	}

	return ValidateMinter(gs.Minter)
}
