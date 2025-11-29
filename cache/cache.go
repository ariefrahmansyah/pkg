package cache

import (
	"errors"
	"sync"
)

type Cache[K comparable, V any] interface {
	Get(key K) (V, error)
	Set(key K, value V) error
	Delete(key K) error
}

type Item[V any] struct {
	val V
}

type SimpleCache[K comparable, V any] struct {
	items sync.Map
}

func NewSimpleCache[K comparable, V any]() *SimpleCache[K, V] {
	return &SimpleCache[K, V]{
		items: sync.Map{},
	}
}

func (c *SimpleCache[K, V]) Get(key K) (val V, err error) {
	item, ok := c.items.Load(key)
	if !ok {
		return val, errors.New("key not found")
	}

	return item.(Item[V]).val, nil
}

func (c *SimpleCache[K, V]) Set(key K, value V) error {
	c.items.Store(key, Item[V]{val: value})
	return nil
}

func (c *SimpleCache[K, V]) Delete(key K) error {
	c.items.Delete(key)
	return nil
}
