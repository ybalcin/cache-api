package adapters

import "github.com/ybalcin/cache-api/pkg/inmemorystore"

type (
	inMemoryCacheAdapter struct {
	}
)

// NewInMemoryCacheAdapter initializes new in memory cache adapter
func NewInMemoryCacheAdapter() *inMemoryCacheAdapter {
	return &inMemoryCacheAdapter{}
}

// Set sets a key-value pair in memory store
func (a *inMemoryCacheAdapter) Set(key string, value string) error {
	return inmemorystore.Set(key, value)
}

// Get gets a value by key from in memory store
func (a *inMemoryCacheAdapter) Get(key string) (string, error) {
	return inmemorystore.Get(key)
}

// FlushCache clears all values in memory store
func (a *inMemoryCacheAdapter) FlushCache() {
	inmemorystore.Flush()
}
