package services

import (
	"github.com/ybalcin/cache-api/internal/core/application/ports/out"
)

type (
	// CacheService interface for used to caching
	CacheService interface {
		Set(key string, value string) error
		Get(key string) (string, error)
		ClearAll()
	}

	cacheService struct {
		cachePort out.Cache
	}
)

// NewCacheService initializes new cache service
func NewCacheService(cachePort out.Cache) *cacheService {
	return &cacheService{cachePort}
}

// Set sets a key-value pair in cache
func (s *cacheService) Set(key string, value string) error {
	return s.cachePort.Set(key, value)
}

// Get gets a value by key from cache
func (s *cacheService) Get(key string) (string, error) {
	return s.cachePort.Get(key)
}

// ClearAll clears all values in cache
func (s *cacheService) ClearAll() {
	s.cachePort.FlushCache()
}
