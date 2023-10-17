package types

import (
	"fmt"
	// this line is used by starport scaffolding # genesis/types/import
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		ScriptList: []Script{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated index in script
	scriptIndexMap := make(map[string]struct{})

	for _, elem := range gs.ScriptList {
		index := string(ScriptKey(elem.Address))
		if _, ok := scriptIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for script")
		}
		scriptIndexMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
