// Package application provides access to the application core
package application

import (
	"github.com/ybalcin/cache-api/internal/application/dtos"
	"github.com/ybalcin/cache-api/internal/application/services"
	"github.com/ybalcin/cache-api/internal/infrastructure/adapters"
	"github.com/ybalcin/cache-api/pkg/inmemorystore"
)

type (
	// cacheService is the interface that wraps the cache operations
	cacheService interface {
		// Set sets a key-value pair in cache
		Set(dto *dtos.CacheDto) error

		// Get gets a value by key from cache
		Get(key string) (*dtos.CacheDto, error)

		// ClearAll clears all values in cache
		ClearAll()
	}
)

// Application struct provides access to the application core
type Application struct {
	CacheService cacheService
}

// New initializes new application
func New() *Application {
	inMemoryClient := inmemorystore.NewClient(0)
	inMemAdapter := adapters.NewInMemoryCacheAdapter(inMemoryClient)
	cacheService := services.NewCacheService(inMemAdapter)

	// load file from disk and cache it
	defer inMemoryClient.LoadToMemoryFromFile()
	// start a save background task
	defer inMemoryClient.StartSaveToFileFromMemoryTask()

	return &Application{cacheService}
}
