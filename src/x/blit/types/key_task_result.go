package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// TaskResultKeyPrefix is the prefix to retrieve all TaskResult
	TaskResultKeyPrefix = "TaskResult/value/"
)

// TaskResultKey returns the store key to retrieve a TaskResult from the index fields
func TaskResultKey(
	id uint64,
) []byte {
	var key []byte

	indexBytes := []byte(string(id))
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
