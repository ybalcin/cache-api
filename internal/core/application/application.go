package application

import (
	"github.com/ybalcin/cache-api/internal/core/application/services"
	"github.com/ybalcin/cache-api/internal/infrastructure/adapters"
)

// Application struct provides access to the application core
type Application struct {
	CacheService services.CacheService
}

// New initializes new application
func New() *Application {
	inMemAdapter := adapters.NewInMemoryCacheAdapter()
	cacheService := services.NewCacheService(inMemAdapter)

	return &Application{cacheService}
}
