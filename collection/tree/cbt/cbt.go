// Package cbt
// Author: hyphen
// Copyright 2022 hyphen. All rights reserved.
// Create-time: 2022/9/30
package cbt

type CompleteBinaryTree[T any] struct {
	cbt []T
}

func NewCompleteBinaryTree[T any]() *CompleteBinaryTree[T] {
	return &CompleteBinaryTree[T]{cbt: []T{}}
}

func (c *CompleteBinaryTree[T]) Insert(value T) {
	c.cbt = append(c.cbt, value)
}

func (c *CompleteBinaryTree[T]) GetLeftChild(index int) T {
	panic("need implement")
}
