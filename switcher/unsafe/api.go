// Package unsafe
// Author: hyphen
// Copyright 2023 hyphen. All rights reserved.
// Create-time: 2023/7/18
package unsafe

import (
	"reflect"
	"unsafe"

	"github.com/heyufeng1001/hyphen-tool-kit/internal/checker"
)

// 此包下存在panic以及unsafe操作，慎用

var (
	ck = checker.NewChecker()
	sm = map[[2]reflect.Type]map[bool]unsafe.Pointer{}
)

type Switcher[F, T any] interface {
	Switch(*F) *T
	BatchSwitch([]*F) []*T
	_idt_unsafe_()
}

// NewSwitcher 所有新增 Switcher 请在 var/初始化 中执行，否则会导致运行时panic
func NewSwitcher[F, T any](ignoreUnsigned bool) Switcher[F, T] {
	f, t := new(F), new(T)
	ft, tt := reflect.TypeOf(f), reflect.TypeOf(t)

	p, ok := sm[[2]reflect.Type{ft, tt}]
	if ok {
		ptr, ok := p[ignoreUnsigned]
		if ok {
			return *(*Switcher[F, T])(ptr)
		}
	} else {
		sm[[2]reflect.Type{ft, tt}] = map[bool]unsafe.Pointer{}
	}

	ck.Check(ft, tt, ignoreUnsigned)
	sw := &switcher[F, T]{From: f, To: t}
	sm[[2]reflect.Type{ft, tt}][ignoreUnsigned] = unsafe.Pointer(sw)
	return sw
}
