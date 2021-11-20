package out

type (
	// Cache is interface that wraps the cache operations as an out cache port
	Cache interface {
		// Set sets a key-value pair in memory store
		Set(key string, value string) error

		// Get gets a value by key from in memory store
		Get(key string) (string, error)

		// FlushCache clears all values in memory store
		FlushCache()
	}
)
