// Package safe_chanel
// Author: hyphen
// Copyright 2023 hyphen. All rights reserved.
// Create-time: 2023/2/13
package safe_chanel

import (
	"sync"

	"github.com/heyufeng1001/hyphen-tool-kit/hyphenutil"
)

type ch[T any] chan T

type SafeChan[T any] struct {
	ch[T]
	once sync.Once
}

func NewSafeChan[T any](size ...int) *SafeChan[T] {
	return &SafeChan[T]{hyphenutil.TernaryForm(len(size) == 0, make(chan T), make(chan T, size[0])), sync.Once{}}
}

func (s *SafeChan[T]) Close() {
	s.once.Do(func() {
		close(s.ch)
	})
}
