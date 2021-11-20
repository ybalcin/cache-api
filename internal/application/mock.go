package application

import (
	"github.com/ybalcin/cache-api/internal/application/dtos"
)

type (
	MockCacheService struct {
		SetFn        func(dto *dtos.CacheDto) error
		SetFnInvoked bool

		GetFn        func(key string) (*dtos.CacheDto, error)
		GetFnInvoked bool

		ClearAllFn      func()
		ClearAllInvoked bool
	}
)

func (s *MockCacheService) Set(dto *dtos.CacheDto) error {
	s.SetFnInvoked = true
	return s.SetFn(dto)
}

func (s *MockCacheService) Get(key string) (*dtos.CacheDto, error) {
	s.GetFnInvoked = true
	return s.GetFn(key)
}

func (s *MockCacheService) ClearAll() {
	s.ClearAllInvoked = true
	s.ClearAllFn()
}
