package types

const (
	// ModuleName defines the module name
	ModuleName = "hello"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_hello"

	// Keep track of the index of posts
	PostKey      = "Post/value/"
	PostCountKey = "Post/count/"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
