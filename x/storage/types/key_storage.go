package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// StorageKeyPrefix is the prefix to retrieve all Storage
	StorageKeyPrefix = "Storage/value/"
)

// StorageKey returns the store key to retrieve a Storage from the index fields
func StorageKey(
	address string,
	index string,
) []byte {
	var key []byte

	addressBytes := []byte(address)
	key = append(key, addressBytes...)
	key = append(key, []byte("/")...)

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
func StorageKeyFilterPrefix(
	address string,
	index string,
) []byte {
	var key []byte

	addressBytes := []byte(address)
	key = append(key, addressBytes...)
	key = append(key, []byte("/")...)

	indexBytes := []byte(index)
	key = append(key, indexBytes...)

	return key
}
