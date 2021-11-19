package adapters

import "github.com/ybalcin/cache-api/pkg/inmemorystore"

type (
	inMemoryCacheAdapter struct {
		client *inmemorystore.Client
	}
)

// NewInMemoryCacheAdapter initializes new in memory cache adapter
func NewInMemoryCacheAdapter(client *inmemorystore.Client) *inMemoryCacheAdapter {
	return &inMemoryCacheAdapter{client}
}

// Set sets a key-value pair in memory store
func (a *inMemoryCacheAdapter) Set(key string, value string) error {
	return a.client.Set(key, value)
}

// Get gets a value by key from in memory store
func (a *inMemoryCacheAdapter) Get(key string) (string, error) {
	return a.client.Get(key)
}

// FlushCache clears all values in memory store
func (a *inMemoryCacheAdapter) FlushCache() {
	a.client.Flush()
}
