package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// FutureTaskKeyPrefix is the prefix to retrieve all FutureTask
	FutureTaskKeyPrefix = "FutureTask/value/"
)

// FutureTaskKey returns the store key to retrieve a FutureTask from the index fields
func FutureTaskKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
