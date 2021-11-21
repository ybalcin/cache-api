package adapters_test

import (
	"github.com/ybalcin/cache-api/internal/infrastructure/adapters"
	"github.com/ybalcin/cache-api/pkg/inmemorystore"
	"reflect"
	"testing"
)

const (
	dummyKey   = "dummyKey"
	dummyValue = "dummyValue"
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

func TestInMemoryCacheAdapter_Set(t *testing.T) {
	testCases := []struct {
		key                string
		value              string
		expected           error
		addToMemoryInvoked bool
	}{
		{dummyKey, dummyValue, nil, true},
		{"", dummyValue, inmemorystore.ErrEmptyKey, true},
		{dummyKey, "", inmemorystore.ErrEmptyValue, true},
	}

	mockInMemClient := new(adapters.MockInMemoryStore)
	adapter := adapters.NewInMemoryCacheAdapter(mockInMemClient)

	for _, c := range testCases {
		mockInMemClient.AddToMemoryInvoked = false
		mockInMemClient.AddToMemoryFn = func(key string, value string) error {
			return c.expected
		}

		actual := adapter.Set(c.key, c.value)
		mustEqual(t, actual, c.expected)
		mustEqual(t, mockInMemClient.AddToMemoryInvoked, c.addToMemoryInvoked)
	}
}

func TestInMemoryCacheAdapter_Get(t *testing.T) {
	testCases := []struct {
		key                  string
		expectedValue        string
		expectedError        error
		getFromMemoryInvoked bool
	}{
		{dummyKey, dummyValue, nil, true},
		{"", "", inmemorystore.ErrNotFoundKey, true},
	}

	mockInMemClient := new(adapters.MockInMemoryStore)
	adapter := adapters.NewInMemoryCacheAdapter(mockInMemClient)

	for _, c := range testCases {
		mockInMemClient.GetFromMemoryInvoked = false
		mockInMemClient.GetFromMemoryFn = func(key string) (string, error) {
			return c.expectedValue, c.expectedError
		}

		val, err := adapter.Get(c.key)
		mustEqual(t, val, c.expectedValue)
		mustEqual(t, err, c.expectedError)
		mustEqual(t, mockInMemClient.GetFromMemoryInvoked, c.getFromMemoryInvoked)
	}
}

func TestInMemoryCacheAdapter_FlushCache(t *testing.T) {
	mockInMemClient := new(adapters.MockInMemoryStore)
	mockInMemClient.ClearAllMemoryFn = func() {
	}

	adapter := adapters.NewInMemoryCacheAdapter(mockInMemClient)

	adapter.FlushCache()
	mustEqual(t, mockInMemClient.ClearAllMemoryInvoked, true)
}

func TestNewInMemoryCacheAdapter(t *testing.T) {
	mockInMemClient := new(adapters.MockInMemoryStore)
	adapter := adapters.NewInMemoryCacheAdapter(mockInMemClient)
	mustNotNil(t, adapter)
}
