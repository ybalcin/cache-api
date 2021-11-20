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

func TestClient_Set(t *testing.T) {
	tests := []struct {
		key      string
		value    string
		expected interface{}
	}{
		{testKey, testVal, nil},
		{testKey, "", ErrEmptyValue},
		{"", testVal, ErrEmptyKey},
	}

	client := NewClient(0)

	for _, c := range tests {
		actual := client.AddToMemory(c.key, c.value)
		mustEqual(t, actual, c.expected)
	}

	cache = nil
	for _, c := range tests {
		actual := client.AddToMemory(c.key, c.value)
		mustEqual(t, actual, c.expected)
	}
}

func TestClient_Get(t *testing.T) {
	cache[testKey] = testVal
	cache[testKey+"1"] = testVal + "1"

	tests := []struct {
		key      string
		expected interface{}
	}{
		{testKey, testVal},
		{testKey + "1", testVal + "1"},
		{"", ErrEmptyKey},
		{fmt.Sprint(time.Now().UnixNano()), ErrNotFoundKey},
	}

	client := NewClient(0)

	for _, c := range tests {
		value, err := client.GetFromMemory(c.key)
		if err != nil {
			mustEqual(t, err, c.expected)
			mustEqual(t, value, "")
		} else {
			mustEqual(t, value, c.expected)
			mustEqual(t, err, nil)
		}
	}

	cache = nil
	val, err := client.GetFromMemory(testVal)
	mustEqual(t, err, ErrNotFoundKey)
	mustEqual(t, val, "")
}

func TestClient_Flush(t *testing.T) {
	cache = storage{}
	cache[testVal] = testVal

	client := NewClient(0)

	client.ClearAllMemory()

	val, err := client.GetFromMemory(testVal)
	mustEqual(t, err, ErrNotFoundKey)
	mustEqual(t, val, "")
}

//func TestClient_Load(t *testing.T) {
//	client := NewClient(1)
//
//	client.LoadToMemoryFromFile()
//
//	if len(cache) <= 0 {
//		t.Fail()
//	}
//}

func TestClient_StartSaveTask(t *testing.T) {
	cache[testKey] = testVal
	cache[testKey+"1"] = testVal + "1"

	//saveCacheToFile(1)
	if err := recover(); err != nil {
		t.Fail()
	}

	//getLatestFile()
}
