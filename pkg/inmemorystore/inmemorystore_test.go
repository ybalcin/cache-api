package inmemorystore

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

const (
	testKey = "cacheKey"
	testVal = "cacheVal"
)

func mustEqual(t *testing.T, actual, expected interface{}) {
	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("expected: %#v, but got %#v", actual, expected)
	}
}

func TestSet(t *testing.T) {
	tests := []struct {
		key      string
		value    string
		expected interface{}
	}{
		{testKey, testVal, nil},
		{testKey, "", errEmptyValue},
		{"", testVal, errEmptyKey},
	}

	for _, c := range tests {
		actual := Set(c.key, c.value)
		mustEqual(t, actual, c.expected)
	}

	cache = nil
	for _, c := range tests {
		actual := Set(c.key, c.value)
		mustEqual(t, actual, c.expected)
	}
}

func TestGet(t *testing.T) {
	cache[testKey] = testVal
	cache[testKey+"1"] = testVal + "1"

	tests := []struct {
		key      string
		expected interface{}
	}{
		{testKey, testVal},
		{testKey + "1", testVal + "1"},
		{"", errEmptyKey},
		{fmt.Sprint(time.Now().UnixNano()), errNotFoundKey},
	}

	for _, c := range tests {
		value, err := Get(c.key)
		if err != nil {
			mustEqual(t, err, c.expected)
			mustEqual(t, value, "")
		} else {
			mustEqual(t, value, c.expected)
			mustEqual(t, err, nil)
		}
	}

	cache = nil
	val, err := Get(testVal)
	mustEqual(t, err, errNotFoundKey)
	mustEqual(t, val, "")
}

func TestFlush(t *testing.T) {
	cache = inMemStore{}
	cache[testVal] = testVal
	Flush()

	val, err := Get(testVal)
	mustEqual(t, err, errNotFoundKey)
	mustEqual(t, val, "")
}
