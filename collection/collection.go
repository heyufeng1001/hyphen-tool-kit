// Package collection
// Author: hyphen
// Copyright 2022 hyphen. All rights reserved.
// Create-time: 2022/8/5
package collection

type Collection[T any] interface {
	Insert(value T)
}
