package types

import (
	"fmt"

	sdktypes "github.com/cosmos/cosmos-sdk/codec/types"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		TaskList:       []Task{},
		FutureTaskList: []FutureTask{},
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

	// Check for duplicated index in futureTask
	futureTaskIndexMap := make(map[string]struct{})

	for _, elem := range gs.FutureTaskList {
		index := string(FutureTaskKey(elem.Status, elem.ScheduledOn, elem.TaskId, elem.GasPrice))
		if _, ok := futureTaskIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for futureTask")
		}
		futureTaskIndexMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}

var _ sdktypes.UnpackInterfacesMessage = GenesisState{}

// UnpackInterfaces implements UnpackInterfacesMessage.UnpackInterfaces
func (data GenesisState) UnpackInterfaces(unpacker sdktypes.AnyUnpacker) error {
	for _, p := range data.TaskList {
		err := p.UnpackInterfaces(unpacker)
		if err != nil {
			return err
		}
	}
	return nil
}
