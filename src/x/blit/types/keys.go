package types

import "cosmossdk.io/collections"

const (
	// ModuleName defines the module name
	ModuleName = "blit"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_blit"

	TaskIndex = "task_index"
)

var (
	ParamsKey            = []byte("p_blit")
	TasksKeyPrefix       = collections.NewPrefix(0x01) // Prefix for tasks
	TaskResultsKeyPrefix = collections.NewPrefix(0x02) // Prefix for task results
	FutureTasksKeyPrefix = collections.NewPrefix(0x03) // Prefix for pending tasks
	TaskIDKey            = collections.NewPrefix(0x04) // Prefix for task ID sequence
	TaskResultIDKey      = collections.NewPrefix(0x05) // Prefix for task result ID sequence
	TaskAddressPrefix    = collections.NewPrefix(0x06) // Prefix for task by address index

)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
