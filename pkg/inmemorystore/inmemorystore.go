// Package inmemorystore is a simple library for storing key-value pair in memory
package inmemorystore

import (
	"fmt"
	"log"
	"strings"
	"sync"
)

type (
	storage map[string]string

	// Client struct inmemory store client
	Client struct {
		interval int
	}

	cacheItem struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	}
	cacheItems struct {
		Data []cacheItem `json:"data"`
	}
)

var (
	lock  = sync.Mutex{}
	cache = storage{}

	errEmptyKey    = newPackageError("key is empty")
	errEmptyValue  = newPackageError("value is empty")
	errNotFoundKey = newPackageError("key not found")
)

func newPackageError(message string) error {
	return fmt.Errorf("inmemorystore: %s", message)
}

// NewClient initializes new inmemorystore client.
// param interval: cache file saving interval time as minutes
func NewClient(interval int) *Client {
	cli := &Client{
		interval: interval,
	}

	return cli
}

// Set sets a key-value pair.
func (c *Client) Set(key string, value string) error {
	if key == "" {
		return errEmptyKey
	}
	if value == "" {
		return errEmptyValue
	}

	lock.Lock()
	defer lock.Unlock()

	if cache == nil {
		cache = storage{}
	}

	key = strings.ReplaceAll(key, " ", "")
	cache[key] = value
	return nil
}

// Get retrieves a value by key.
func (c *Client) Get(key string) (string, error) {
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
func (c *Client) Flush() {
	lock.Lock()
	defer lock.Unlock()

	cache = make(map[string]string)
}

// Load loads cache from last saved file
func (c *Client) Load() {
	lock.Lock()
	defer lock.Unlock()

	fileCache := getCacheFromFile()
	if len(fileCache) > 0 {
		cache = fileCache
		log.Println("File loaded.")
	}
}
