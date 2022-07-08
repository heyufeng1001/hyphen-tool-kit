// Package concurrentmap
// Author: hyphen
// Copyright 2022 hyphen. All rights reserved.
// Create-time: 2022/7/8
package concurrentmap

import "sync"

type ConcurrentMap[K comparable, V any] struct {
	valueMap map[K]V
	rwLock   sync.RWMutex
}

func NewConcurrentMap[K comparable, V any]() *ConcurrentMap[K, V] {
	return &ConcurrentMap[K, V]{
		valueMap: map[K]V{},
		rwLock:   sync.RWMutex{},
	}
}

func (c *ConcurrentMap[K, V]) Get(key K) (V, bool) {
	c.rwLock.RLock()
	value, ok := c.valueMap[key]
	c.rwLock.RUnlock()
	return value, ok
}

func (c *ConcurrentMap[K, V]) Set(key K, value V) {
	c.rwLock.Lock()
	defer c.rwLock.Unlock()
	c.valueMap[key] = value
}

func (c *ConcurrentMap[K, V]) Delete(key K) {
	c.rwLock.Lock()
	defer c.rwLock.Unlock()
	delete(c.valueMap, key)
}
