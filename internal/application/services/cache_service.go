package services

import (
	"github.com/ybalcin/cache-api/internal/application/dtos"
	"github.com/ybalcin/cache-api/internal/ports/out"
)

type (
	// cacheService implements application.cacheService
	cacheService struct {
		cachePort out.Cache
	}
)

// NewCacheService initializes new cache service
func NewCacheService(cachePort out.Cache) *cacheService {
	return &cacheService{cachePort}
}

// Set sets a key-value pair in cache
func (s *cacheService) Set(dto *dtos.CacheDto) error {
	return s.cachePort.Set(dto.Key, dto.Value)
}

// Get gets a value by key from cache
func (s *cacheService) Get(key string) (*dtos.CacheDto, error) {
	val, err := s.cachePort.Get(key)
	if err != nil {
		return nil, err
	}

	return &dtos.CacheDto{
		Key:   key,
		Value: val,
	}, nil
}

// ClearAll clears all values in cache
func (s *cacheService) ClearAll() {
	s.cachePort.FlushCache()
}
