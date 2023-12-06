// Package heap
// Author: hyphen
// Copyright 2022 hyphen. All rights reserved.
// Create-time: 2022/9/30
package heap

type Heap[T comparable] struct {
	values []T
	less   func(s, b T) bool
}

// NewHeap
// Variable less returns the result of compare between p and c,
// p is the parent while c is the child.
// In detail, if you want to construct a min-heap, please make sure less returns
// true when s<b, as max-heap is on the opposite.
func NewHeap[T comparable](less func(s, b T) bool) *Heap[T] {
	return &Heap[T]{[]T{}, less}
}

// Heapify Construct a heap by a slice
func Heapify[T comparable](values []T, less func(s, b T) bool) *Heap[T] {

	panic("need implement")
}

func (h *Heap[T]) Size() int {
	return len(h.values)
}

func (h *Heap[T]) IsEmpty() bool {
	return h.Size() == 0
}

func (h *Heap[T]) Insert(value T) {
	h.values = append(h.values, value)
	h.liftBottom()
	return
}

func (h *Heap[T]) liftBottom() {
	index := len(h.values) - 1
	for index > 0 {
		pIndex := (index+index%2)/2 - 1
		if h.less(h.values[index], h.values[pIndex]) {
			break
		}
		h.values[index], h.values[pIndex] = h.values[pIndex], h.values[index]
		index = pIndex
	}
	return
}

func (h *Heap[T]) Top() (ret T) {
	if h.IsEmpty() {
		return
	}
	return h.values[0]
}

func (h *Heap[T]) Pop() (ret T) {
	ret = h.Top()
	if h.IsEmpty() {
		return
	}
	h.siftTop()
	return
}

func (h *Heap[T]) siftTop() {
	newTop := h.values[h.Size()-1]
	h.values = h.values[:h.Size()-1]
	if h.IsEmpty() {
		return
	}
	h.values[0] = newTop
	index := 0
	for index*2+1 < h.Size() {
		cIndex := index*2 + 1
		child := h.values[cIndex]
		rIndex := index*2 + 2
		if rIndex < h.Size() && !h.less(child, h.values[rIndex]) {
			child = h.values[rIndex]
			cIndex = rIndex
		}
		if h.less(h.values[index], child) {
			break
		}
		h.values[index], h.values[cIndex] = h.values[cIndex], h.values[index]
		index = cIndex
	}
	return
}
