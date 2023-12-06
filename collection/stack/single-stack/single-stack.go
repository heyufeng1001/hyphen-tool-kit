// Package single_stack
// Author: hyphen
// Copyright 2022 hyphen. All rights reserved.
// Create-time: 2022/10/28
package single_stack

type SingleStack[T comparable] struct {
	val []T
	top int
	// pre代表早入栈的元素，cmp返回true是将会pop pre
	cmp func(pre, back T) bool
}

func (s *SingleStack[T]) Pop() (T, bool) {
	var ret T
	if s.Len() != 0 {
		ret = s.val[s.top-1]
		s.top--
		return ret, true
	}
	return ret, false
}

func (s *SingleStack[T]) Peek() (T, bool) {
	var ret T
	if s.Len() != 0 {
		ret = s.val[s.top-1]
		return ret, true
	}
	return ret, false
}

func (s *SingleStack[T]) Push(v T) {
	for s.Len() > 0 && s.cmp(s.MustPeek(), v) {
		s.MustPop()
	}
	s.val = append(s.val, v)
	s.top++
}

func (s *SingleStack[T]) Len() int {
	return s.top
}

func (s *SingleStack[T]) MustPop() T {
	var ret T
	if s.Len() != 0 {
		ret = s.val[s.Len()-1]
		s.top--
	}
	return ret
}

func (s *SingleStack[T]) MustPeek() T {
	var ret T
	if s.Len() != 0 {
		ret = s.val[s.top-1]
		return ret
	}
	return ret
}
