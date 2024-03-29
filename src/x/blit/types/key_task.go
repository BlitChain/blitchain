package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// TaskKeyPrefix is the prefix to retrieve all Task
	TaskKeyPrefix = "Task/value/"
)

// TaskKey returns the store key to retrieve a Task from the index fields
func TaskKey(
	id uint64,
) []byte {
	var key []byte

	indexBytes := []byte(string(id))
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
