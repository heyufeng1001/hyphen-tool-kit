// Package tag
// Author: hyphen
// Copyright 2023 hyphen. All rights reserved.
// Create-time: 2023/7/19
package tag

import (
	"context"
)

const (
	DefaultTagName = "switch"
)

type internalSwitcherOptioner interface {
	setTagName(string)
	setErrable(bool)
	setUseMemberName(bool)
	setIgnoreUnsigned(bool)
	setIgnorePtr(bool)
	setSkipTypeMismatch(bool)
}

type Switcher[F, T any] interface {
	Switch(context.Context, *F) (*T, error)
	Unswitch(context.Context, *T) (*F, error)
	internalSwitcherOptioner
}

func NewSwitcher[F, T any](opts ...Option) Switcher[F, T] {
	sw := newDefaultSwitcher[F, T]()
	for _, opt := range opts {
		opt(sw)
	}
	return sw
}

type Option func(internalSwitcherOptioner)

func WithTagName[F, T any](tag string) Option {
	return func(s internalSwitcherOptioner) {
		s.setTagName(tag)
	}
}

func WithIgnorePtr() Option {
	return func(s internalSwitcherOptioner) {
		s.setIgnorePtr(true)
	}
}

func WithIgnoreUnsigned() Option {
	return func(s internalSwitcherOptioner) {
		s.setIgnoreUnsigned(true)
	}
}

func WithUseMemberName() Option {
	return func(s internalSwitcherOptioner) {
		s.setUseMemberName(true)
	}
}

func WithErrable() Option {
	return func(s internalSwitcherOptioner) {
		s.setErrable(true)
	}
}

func WithSkipTypeMismatch() Option {
	return func(s internalSwitcherOptioner) {
		s.setSkipTypeMismatch(true)
	}
}
