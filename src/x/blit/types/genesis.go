package types

import (
	"fmt"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		TaskList:       []Task{},
		TaskResultList: []TaskResult{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated index in task
	taskIndexMap := make(map[string]struct{})

	for _, elem := range gs.TaskList {
		key := string(TaskKey(elem.Id))
		if _, ok := taskIndexMap[key]; ok {
			return fmt.Errorf("duplicated index for task")
		}
		taskIndexMap[key] = struct{}{}
	}
	// Check for duplicated index in taskResult
	taskResultIndexMap := make(map[string]struct{})

	for _, elem := range gs.TaskResultList {
		index := string(TaskResultKey(elem.Index))
		if _, ok := taskResultIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for taskResult")
		}
		taskResultIndexMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
