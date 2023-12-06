// Package unsafe
// Author: hyphen
// Copyright 2023 hyphen. All rights reserved.
// Create-time: 2023/7/18
package unsafe

import (
	"reflect"
	"unsafe"
)

type switcher[F, T any] struct {
	From *F
	To   *T
}

// Switch 使用此方法必须保证F和T的内存布局严格一致，否则会导致unsafe操作带来的不可预期情形，慎用！
// 标准的使用方式请查看 api_test.go ，不要瞎搞，更不要随便改
func (s *switcher[F, T]) Switch(from *F) (to *T) {
	tt := reflect.TypeOf(s.To)
	fp := unsafe.Pointer(from)
	to = reflect.NewAt(tt.Elem(), fp).Interface().(*T)
	return
}

func (s *switcher[F, T]) _idt_unsafe_() {

}

func (s *switcher[F, T]) BatchSwitch(fs []*F) []*T {
	ret := []*T{}
	for _, f := range fs {
		ret = append(ret, s.Switch(f))
	}
	return nil
}
