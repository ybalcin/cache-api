package adapters

import "github.com/ybalcin/cache-api/pkg/inmemorystore"

type (

	// inMemoryCacheAdapter implements out.Cache
	inMemoryCacheAdapter struct {
		client inmemorystore.Client
	}
)

// NewInMemoryCacheAdapter initializes new in memory cache adapter
func NewInMemoryCacheAdapter(client inmemorystore.Client) *inMemoryCacheAdapter {
	return &inMemoryCacheAdapter{client}
}

// Set sets a key-value pair in memory store
func (a *inMemoryCacheAdapter) Set(key string, value string) error {
	return a.client.AddToMemory(key, value)
}

// Get gets a value by key from in memory store
func (a *inMemoryCacheAdapter) Get(key string) (string, error) {
	return a.client.GetFromMemory(key)
}

// FlushCache clears all values in memory store
func (a *inMemoryCacheAdapter) FlushCache() {
	a.client.ClearAllMemory()
}
