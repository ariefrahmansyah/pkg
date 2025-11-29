package cache

import (
	"errors"
	"sync"
)

// Cache is a generic cache interface.
type Cache[K comparable, V any] interface {
	// Get retrieves the value of the item from the cache.
	Get(key K) (V, error)

	// Set stores the value of the item in the cache.
	Set(key K, value V) error

	// Delete removes the item from the cache.
	Delete(key K) error
}

// Item represents a generic item in the cache.
type Item[V any] struct {
	val V
}

// SimpleCache is a simple cache implementation using a sync.Map.
type SimpleCache[K comparable, V any] struct {
	items sync.Map
}

// NewSimpleCache creates a new SimpleCache.
func NewSimpleCache[K comparable, V any]() *SimpleCache[K, V] {
	return &SimpleCache[K, V]{
		items: sync.Map{},
	}
}

// Get retrieves the value of the item from the cache.
func (c *SimpleCache[K, V]) Get(key K) (val V, err error) {
	item, ok := c.items.Load(key)
	if !ok {
		return val, errors.New("key not found")
	}

	return item.(Item[V]).val, nil
}

// Set stores the value of the item in the cache.
func (c *SimpleCache[K, V]) Set(key K, value V) error {
	c.items.Store(key, Item[V]{val: value})
	return nil
}

// Delete removes the item from the cache.
func (c *SimpleCache[K, V]) Delete(key K) error {
	c.items.Delete(key)
	return nil
}
