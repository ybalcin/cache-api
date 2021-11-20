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
		t.Fatalf("expected: %#v, but got %#v", actual, expected)
	}
}

func TestInMemoryCacheAdapter_Set(t *testing.T) {
	testCases := []struct {
		key      string
		value    string
		expected error
	}{
		{dummyKey, dummyValue, nil},
		{"", dummyValue, inmemorystore.ErrEmptyKey},
		{dummyKey, "", inmemorystore.ErrEmptyValue},
	}

	mockInMemClient := new(adapters.MockInMemoryStore)
	adapter := adapters.NewInMemoryCacheAdapter(mockInMemClient)

	for _, c := range testCases {
		mockInMemClient.AddToMemoryFn = func(key string, value string) error {
			return c.expected
		}

		actual := adapter.Set(c.key, c.value)
		mustEqual(t, actual, c.expected)
	}
}
