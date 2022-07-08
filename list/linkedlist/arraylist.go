// Package linkedlist
// Author: hyphen
// Copyright 2022 hyphen. All rights reserved.
// Create-time: 2022/7/8
package linkedlist

import (
	"errors"

	"github.com/heyufeng1001/hyphen-tool-kit/util"
)

type element[T comparable] struct {
	value      T
	prev, next *element[T]
}

func linkElement[T comparable](prev, next *element[T]) {
	prev.next = next
	next.prev = prev
}

func newElement[T comparable](value T, prev, next *element[T]) *element[T] {
	return &element[T]{
		value: value,
		prev:  prev,
		next:  next,
	}
}

type LinkedList[T comparable] struct {
	head, bottom *element[T]
	size         int
}

func NewLinkedList[T comparable]() *LinkedList[T] {
	return &LinkedList[T]{
		head:   nil,
		bottom: nil,
		size:   0,
	}
}

func (l *LinkedList[T]) Len() int {
	return l.size
}

func (l *LinkedList[T]) InsertBottom(value T) {
	if l.head == nil {
		l.head = newElement(value, nil, nil)
		l.bottom = l.head
	} else {
		l.bottom.next = newElement(value, l.bottom, nil)
		l.bottom = l.bottom.next
	}
	l.size++
}

func (l *LinkedList[T]) InsertHead(value T) {
	if l.head == nil {
		l.head = newElement(value, nil, nil)
		l.bottom = l.head
	} else {
		l.head.prev = newElement(value, nil, l.head)
		l.head = l.head.prev
	}
	l.size++
}

func (l *LinkedList[T]) MoveToFront(value T) error {
	node := l.head
	for node != nil {
		if node.value == value {
			linkElement(node.prev, node.next)
			linkElement(node, l.head)
			return nil
		}
	}
	return errors.New("[MoveToFront]failed to find value")
}

func (l *LinkedList[T]) PopBottom() (T, error) {
	return util.TernaryForm(l.Len() == 0, func() (T, error) { return nil, errors.New("[PopBottom]list is empty") },
		func() (T, error) {
			ret := l.bottom
			util.TernaryForm(l.Len() == 1, func() { l.head, l.bottom = nil, nil }, func() { l.bottom.prev.next, l.bottom = nil, l.bottom.prev })()
			l.size--
			return ret.value, nil
		},
	)()
}
