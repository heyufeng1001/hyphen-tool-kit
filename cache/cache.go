// Package cache
// Author: hyphen
// Copyright 2022 hyphen. All rights reserved.
// Create-time: 2022/7/8
package cache

type Cache[K comparable, V any] interface {
	Get(key K) V
	Set(key K, value V)
}
