package types

import (
	types "cosmossdk.io/math"
)

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		PortId: PortID,
		// this line is used by starport scaffolding # genesis/types/default
		Params:           DefaultParams(),
		NextCollectionId: types.NewUint(1),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.

// IMPORTANT: We assume badges are well-formed and validated here
func (gs GenesisState) Validate() error {

	return gs.Params.Validate()
}
