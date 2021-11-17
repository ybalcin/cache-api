// Package inmemorystore is a simple library for storing key-value pair in memory
package inmemorystore

import (
	"fmt"
	"sync"
)

type (
	inMemStore map[string]string
)

var (
	lock  = sync.Mutex{}
	cache = inMemStore{}

	errEmptyKey    = newPackageError("key is empty!")
	errEmptyValue  = newPackageError("value is empty!")
	errNotFoundKey = newPackageError("key not found!")
)

func newPackageError(message string) error {
	return fmt.Errorf("inmemorystore: %s", message)
}

// Set sets a cacheVal for the cacheKey.
func Set(key string, value string) error {
	if key == "" {
		return errEmptyKey
	}
	if value == "" {
		return errEmptyValue
	}

	lock.Lock()
	defer lock.Unlock()

	if cache == nil {
		cache = make(map[string]string)
	}

	cache[key] = value
	return nil
}

// Get retrieves a cacheVal by cacheKey.
func Get(key string) (string, error) {
	if key == "" {
		return "", errEmptyKey
	}

	lock.Lock()
	defer lock.Unlock()

	val, ok := cache[key]
	if !ok {
		return "", errNotFoundKey
	}

	return val, nil
}

// Flush clears all values that keeping in memory.
func Flush() {
	lock.Lock()
	defer lock.Unlock()

	cache = make(map[string]string)
}
