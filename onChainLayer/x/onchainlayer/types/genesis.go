package types

import (
	"fmt"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		RegistrandList: []Registrand{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated ID in registrand
	registrandIdMap := make(map[uint64]bool)
	registrandCount := gs.GetRegistrandCount()
	for _, elem := range gs.RegistrandList {
		if _, ok := registrandIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for registrand")
		}
		if elem.Id >= registrandCount {
			return fmt.Errorf("registrand id should be lower or equal than the last id")
		}
		registrandIdMap[elem.Id] = true
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
