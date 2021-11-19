// Package application provides access to the application core
package application

import (
	"github.com/ybalcin/cache-api/internal/application/services"
	"github.com/ybalcin/cache-api/internal/infrastructure/adapters"
	"github.com/ybalcin/cache-api/pkg/inmemorystore"
)

// Application struct provides access to the application core
type Application struct {
	CacheService services.CacheService
}

// New initializes new application
func New() *Application {
	inMemoryClient := inmemorystore.NewClient(60, "")
	inMemAdapter := adapters.NewInMemoryCacheAdapter(inMemoryClient)
	cacheService := services.NewCacheService(inMemAdapter)

	return &Application{cacheService}
}
