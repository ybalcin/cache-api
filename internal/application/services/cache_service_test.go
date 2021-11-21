package services_test

import (
	"errors"
	"github.com/ybalcin/cache-api/internal/application/dtos"
	"github.com/ybalcin/cache-api/internal/application/services"
	"github.com/ybalcin/cache-api/internal/infrastructure/adapters"
	"reflect"
	"testing"
)

func mustEqual(t *testing.T, actual, expected interface{}) {
	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("expectedValue: %#v, but got %#v", actual, expected)
	}
}

func mustNotNil(t *testing.T, val interface{}) {
	if val == nil {
		t.Fatalf("%#v must not nil", val)
	}
}

const (
	dummyKey   = "dummyKey"
	dummyValue = "dummyValue"
)

func TestCacheService_Set(t *testing.T) {
	testCases := []struct {
		dto               *dtos.CacheDto
		expected          error
		adapterSetInvoked bool
	}{
		{&dtos.CacheDto{Key: dummyKey, Value: dummyValue}, nil, true},
		{&dtos.CacheDto{Key: "", Value: dummyValue}, errors.New(""), true},
		{&dtos.CacheDto{Key: dummyKey, Value: ""}, errors.New(""), true},
	}

	mockCacheAdapter := new(adapters.MockCacheAdapter)
	service := services.NewCacheService(mockCacheAdapter)

	for _, c := range testCases {
		mockCacheAdapter.SetInvoked = false
		mockCacheAdapter.SetFn = func(key string, value string) error {
			return c.expected
		}

		err := service.Set(c.dto)
		if c.expected != nil {
			mustNotNil(t, err)
		}

		mustEqual(t, mockCacheAdapter.SetInvoked, c.adapterSetInvoked)
	}
}

func TestCacheService_Get(t *testing.T) {
	testCases := []struct {
		key         string
		value       string
		expectedDto *dtos.CacheDto
		expectedErr error
		getInvoked  bool
	}{
		{dummyKey, dummyValue, &dtos.CacheDto{Key: dummyKey, Value: dummyValue}, nil, true},
		{"", "", nil, errors.New(""), true},
	}

	mockCacheAdapter := new(adapters.MockCacheAdapter)
	service := services.NewCacheService(mockCacheAdapter)

	for _, c := range testCases {
		mockCacheAdapter.GetInvoked = false
		mockCacheAdapter.GetFn = func(key string) (string, error) {
			return c.value, c.expectedErr
		}

		actual, err := service.Get(c.key)
		if c.expectedErr != nil {
			mustNotNil(t, err)
		}

		mustEqual(t, actual, c.expectedDto)
		mustEqual(t, mockCacheAdapter.GetInvoked, c.getInvoked)
	}
}

func TestCacheService_ClearAll(t *testing.T) {
	mockCacheAdapter := new(adapters.MockCacheAdapter)
	service := services.NewCacheService(mockCacheAdapter)

	mockCacheAdapter.FlushCacheFn = func() {
	}

	service.ClearAll()
	mustEqual(t, mockCacheAdapter.FlushInvoked, true)
}

func TestNewCacheService(t *testing.T) {
	mockCacheAdapter := new(adapters.MockCacheAdapter)
	service := services.NewCacheService(mockCacheAdapter)
	mustNotNil(t, service)
}
