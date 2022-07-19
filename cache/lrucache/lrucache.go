// Package lrucache
// Author: hyphen
// Copyright 2022 hyphen. All rights reserved.
// Create-time: 2022/7/8
package lrucache

import (
	"github.com/heyufeng1001/hyphen-tool-kit/list/linkedlist"
	"github.com/heyufeng1001/hyphen-tool-kit/map/concurrentmap"
)

type LRUCache[K comparable, V any] struct {
	maxCacheSize, size int
	valueMap           *concurrentmap.ConcurrentMap[K, V]
	keyList            *linkedlist.LinkedList[K]
	remoteGet          func(K, ...any) (V, error)
	autoUpdate         bool
}

func NewLRUCache[K comparable, V any](maxCacheSize int, remoteGet func(K, ...any) (V, error), autoUpdate bool) *LRUCache[K, V] {
	return &LRUCache[K, V]{
		maxCacheSize: maxCacheSize,
		size:         0,
		valueMap:     concurrentmap.NewConcurrentMap[K, V](),
		keyList:      linkedlist.NewLinkedList[K](),
		remoteGet:    remoteGet,
		autoUpdate:   autoUpdate,
	}
}

func (l *LRUCache[K, V]) Get(key K, extra ...any) (V, error) {
	var err error
	value, ok := l.valueMap.Get(key)
	if !ok {
		value, err = l.remoteGet(key, extra...)
	}
	if err != nil {
		go l.update(key, value, ok)
	}
	return value, err
}

func (l *LRUCache[K, V]) Set(key K, value V) {
	l.basicUpdate(key, value)
}

func (l *LRUCache[K, V]) autoUpdateKey(key K) {
	remoteValue, err := l.remoteGet(key)
	if err != nil {
		return
	}
	l.valueMap.Set(key, remoteValue)
}

func (l *LRUCache[K, V]) update(key K, value V, isExist bool) {
	if isExist && l.autoUpdate {
		remoteValue, err := l.remoteGet(key)
		if err != nil {
			return
		}
		value = remoteValue
	}
	l.basicUpdate(key, value)
}

func (l *LRUCache[K, V]) basicUpdate(key K, value V) {
	if l.size == l.maxCacheSize {
		bottomKey, _ := l.keyList.PopBottom()
		l.valueMap.Delete(bottomKey)
		l.size--
	}
	l.valueMap.Set(key, value)
	l.keyList.InsertHead(key)
	l.size++
}
