package types

const (
	// ModuleName defines the module name
	ModuleName = "blit"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_blit"
)

var (
	ParamsKey = []byte("p_blit")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
