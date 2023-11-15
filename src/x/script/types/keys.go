package types

const (
	// ModuleName defines the module name
	ModuleName = "script"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_script"
)

var (
	ParamsKey = []byte("p_script")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
