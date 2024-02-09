package types

const (
	// ModuleName defines the module name
	ModuleName = "services"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_services"

    
)

var (
	ParamsKey = []byte("p_services")
)



func KeyPrefix(p string) []byte {
    return []byte(p)
}
