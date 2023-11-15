package types

import (
	"fmt"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		StorageList: []Storage{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated index in storage
	storageIndexMap := make(map[string]struct{})

	for _, elem := range gs.StorageList {
		index := string(StorageKey(elem.Address, elem.Index))
		if _, ok := storageIndexMap[index]; ok {
			return fmt.Errorf("duplicated address + index for storage")
		}
		storageIndexMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
