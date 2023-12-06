// Package checker
// Author: hyphen
// Copyright 2023 hyphen. All rights reserved.
// Create-time: 2023/7/18
// nolint: byted_s_panic_detect
package checker

import (
	"fmt"
	"reflect"
)

type Checker struct {
	Checked map[[2]reflect.Type]bool
}

func NewChecker() *Checker {
	return &Checker{
		Checked: map[[2]reflect.Type]bool{},
	}
}

// nolint:cyclo_complexity
func (c *Checker) Check(source, target reflect.Type, iu bool) {
	thisCheck := [2]reflect.Type{source, target}
	if c.Checked[thisCheck] {
		return
	}
	c.Checked[thisCheck] = true

	if source.Kind() != target.Kind() && !(iu && c.ignoreUnsigned(source, target)) {
		panic(fmt.Errorf("expect %s and %s have same Kind, but %v != %v", source, target, source.Kind(), target.Kind()))
	}

	switch source.Kind() {
	case reflect.Bool:
	case reflect.Int:
	case reflect.Int8:
	case reflect.Int16:
	case reflect.Int32:
	case reflect.Int64:
	case reflect.Uint:
	case reflect.Uint8:
	case reflect.Uint16:
	case reflect.Uint32:
	case reflect.Uint64:
	case reflect.Uintptr:
	case reflect.Float32:
	case reflect.Float64:
	case reflect.Complex64:
	case reflect.Complex128:
	case reflect.String:
	case reflect.UnsafePointer:
	// 以上为基础类型，不需要进一步比较

	case reflect.Array:
		if source.Len() != target.Len() {
			panic(fmt.Errorf("expect %s and %s have same Len, but %v != %v", source, target, source.Len(), target.Len()))
		}
		c.Check(source.Elem(), target.Elem(), iu)
	case reflect.Chan:
		c.Check(source.Elem(), target.Elem(), iu)
	case reflect.Func:
		c.checkFunction(source, target, iu)
	case reflect.Map:
		c.Check(source.Key(), target.Key(), iu)
		c.Check(source.Elem(), target.Elem(), iu)
	case reflect.Ptr:
		c.Check(source.Elem(), target.Elem(), iu)
	case reflect.Slice:
		c.Check(source.Elem(), target.Elem(), iu)
	case reflect.Struct:
		c.checkStruct(source, target, iu)
	case reflect.Interface:
		// todo support interface
		panic(fmt.Errorf("expect %s.%s and %s.%s have same Type, but Interface found", source.PkgPath(), source.Name(), target.PkgPath(), target.Name()))
	}
}

func (c *Checker) checkFunction(source, target reflect.Type, iu bool) {
	if source.NumIn() != target.NumIn() {
		panic(fmt.Errorf("expect %s.%s and %s.%s have same NumIn, but %v != %v", source.PkgPath(), source.Name(), target.PkgPath(), target.Name(), source.NumIn(), target.NumIn()))
	}
	if source.NumOut() != target.NumOut() {
		panic(fmt.Errorf("expect %s.%s and %s.%s have same NumOut, but %v != %v", source.PkgPath(), source.Name(), target.PkgPath(), target.Name(), source.NumOut(), target.NumOut()))
	}
	for i := 0; i < source.NumIn(); i++ {
		sourceField, targetField := source.In(i), target.In(i)
		c.Check(sourceField, targetField, iu)
	}
	for i := 0; i < source.NumOut(); i++ {
		sourceField, targetField := source.Out(i), target.Out(i)
		c.Check(sourceField, targetField, iu)
	}
}

func (c *Checker) checkStruct(source, target reflect.Type, iu bool) {
	if source.NumField() != target.NumField() {
		panic(fmt.Errorf("expect %s.%s and %s.%s have same NumField, but %d != %d", source.PkgPath(), source.Name(), target.PkgPath(), target.Name(), source.NumField(), target.NumField()))
	}

	for i := 0; i < source.NumField(); i++ {
		sourceField, targetField := source.Field(i), target.Field(i)
		c.Check(sourceField.Type, targetField.Type, iu)
	}
}

var igm = map[reflect.Kind]reflect.Kind{
	reflect.Int:   reflect.Uint,
	reflect.Int8:  reflect.Uint8,
	reflect.Int16: reflect.Uint16,
	reflect.Int32: reflect.Uint32,
	reflect.Int64: reflect.Uint64,

	reflect.Uint:   reflect.Int,
	reflect.Uint8:  reflect.Int8,
	reflect.Uint16: reflect.Int16,
	reflect.Uint32: reflect.Int32,
	reflect.Uint64: reflect.Int64,
}

func (c *Checker) ignoreUnsigned(source, target reflect.Type) bool {
	k, ok := igm[source.Kind()]
	if !ok {
		return false
	}
	return k == target.Kind()
}
