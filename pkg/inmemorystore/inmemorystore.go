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

	client struct {
		interval int
	}

	// Client is interface that wraps the in memory cache operations
	Client interface {
		// AddToMemory sets a key-value pair.
		AddToMemory(key string, value string) error

		// GetFromMemory retrieves a value by key.
		GetFromMemory(key string) (string, error)

		// ClearAllMemory clears all values that keeping in memory.
		ClearAllMemory()

		// LoadToMemoryFromFile loads cache from last saved file
		LoadToMemoryFromFile()

		// StartSaveToFileFromMemoryTask starts a task that saves cache to file in a specified interval time of minutes.
		StartSaveToFileFromMemoryTask()
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

	ErrEmptyKey    = newPackageError("key is empty")
	ErrEmptyValue  = newPackageError("value is empty")
	ErrNotFoundKey = newPackageError("key not found")
)

func newPackageError(message string) error {
	return fmt.Errorf("inmemorystore: %s", message)
}

// NewClient initializes new inmemorystore client.
// param interval: cache file saving interval time as minutes
func NewClient(interval int) Client {
	cli := &client{
		interval: interval,
	}

	return cli
}

// AddToMemory sets a key-value pair.
func (c *client) AddToMemory(key string, value string) error {
	if key == "" {
		return ErrEmptyKey
	}
	if value == "" {
		return ErrEmptyValue
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

// GetFromMemory retrieves a value by key.
func (c *client) GetFromMemory(key string) (string, error) {
	if key == "" {
		return "", ErrEmptyKey
	}

	lock.Lock()
	defer lock.Unlock()

	val, ok := cache[key]
	if !ok {
		return "", ErrNotFoundKey
	}

	return val, nil
}

// ClearAllMemory clears all values that keeping in memory.
func (c *client) ClearAllMemory() {
	lock.Lock()
	defer lock.Unlock()

	cache = make(map[string]string)
}

// LoadToMemoryFromFile loads cache from last saved file
func (c *client) LoadToMemoryFromFile() {
	lock.Lock()
	defer lock.Unlock()

	fileCache := getCacheFromFile()
	if len(fileCache) > 0 {
		cache = fileCache
		log.Println("File loaded.")
	}
}

// StartSaveToFileFromMemoryTask starts a task that saves cache to file in a specified interval time of minutes.
func (c *client) StartSaveToFileFromMemoryTask() {
	if c.interval <= 0 {
		return
	}

	startSaveTask(c.interval)
}
