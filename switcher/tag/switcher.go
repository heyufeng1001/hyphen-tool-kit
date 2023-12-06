// Package tag
// Author: hyphen
// Copyright 2023 hyphen. All rights reserved.
// Create-time: 2023/7/20
package tag

import (
	"context"
	"fmt"
	"log"
	"reflect"
)

type switcher[F, T any] struct {
	tagName          string
	errable          bool
	ignoreUnsigned   bool
	ignorePtr        bool
	useMemberName    bool
	skipTypeMismatch bool
}

func (s *switcher[F, T]) setTagName(tag string) {
	s.tagName = tag
}

func (s *switcher[F, T]) setErrable(b bool) {
	s.errable = b
}

func (s *switcher[F, T]) setUseMemberName(b bool) {
	s.useMemberName = b
}

func (s *switcher[F, T]) setIgnoreUnsigned(b bool) {
	s.ignoreUnsigned = b
}

func (s *switcher[F, T]) setIgnorePtr(b bool) {
	s.ignorePtr = b
}

func (s *switcher[F, T]) setSkipTypeMismatch(b bool) {
	s.skipTypeMismatch = b
}

var (
	ErrOnlyStructAllowed = fmt.Errorf("only struct can be switched")
	ErrMemberTypeDiff    = fmt.Errorf("member's type must be the same")
)

func newDefaultSwitcher[F, T any]() *switcher[F, T] {
	return &switcher[F, T]{
		tagName:          DefaultTagName,
		errable:          false,
		ignoreUnsigned:   false,
		ignorePtr:        false,
		useMemberName:    false,
		skipTypeMismatch: false,
	}
}

func (s *switcher[F, T]) Switch(ctx context.Context, f *F) (*T, error) {
	t := new(T)

	fv := reflect.ValueOf(f).Elem()
	tv := reflect.ValueOf(t).Elem()
	if fv.Kind() != reflect.Struct || tv.Kind() != reflect.Struct {
		return t, s.throw(ctx, ErrOnlyStructAllowed)
	}

	nm, tm := map[string]int{}, map[string]int{}
	for i := 0; i < tv.NumField(); i++ {
		field := tv.Type().Field(i)
		nm[field.Name] = i
		aim := field.Tag.Get(s.tagName)
		if aim != "" {
			tm[aim] = i
		}
	}

	for i := 0; i < fv.NumField(); i++ {
		aim := fv.Type().Field(i).Tag.Get(s.tagName)
		j, ok := tm[aim]
		if !ok || aim == "" {
			if !s.useMemberName {
				continue
			}
			aim = fv.Type().Field(i).Name
			j, ok = nm[aim]
			if !ok {
				continue
			}
		}

		ff, tf := fv.Field(i), tv.Field(j)

		iu := s.ignoreUnsigned && ignoreUnsigned(ff.Type(), tf.Type())
		ip := s.ignorePtr && ignorePtr(ff.Type(), tf.Type())
		if ff.Kind() != tf.Kind() && !iu && !ip {
			log.Printf("[Switch]fv: %s; tv: %s", ff.Kind(), tf.Kind())
			if s.skipTypeMismatch {
				tf.Set(reflect.New(tf.Type()))
				continue
			}
			return t, s.throw(ctx, ErrMemberTypeDiff)
		}
		if iu {
			tf.Set(ff.Convert(tv.Field(i).Type()))
			continue
		} else if ip {
			if ff.Kind() == reflect.Pointer {
				tf.Set(ff.Elem())
			} else {
				nf := reflect.New(ff.Type())
				nf.Elem().Set(ff)
				tf.Set(nf)
			}
			continue
		}
		tf.Set(ff)
	}
	return t, nil
}

func (s *switcher[F, T]) Unswitch(ctx context.Context, t *T) (*F, error) {
	f := new(F)

	fv := reflect.ValueOf(f).Elem()
	tv := reflect.ValueOf(t).Elem()
	if fv.Kind() != reflect.Struct || tv.Kind() != reflect.Struct {
		return f, s.throw(ctx, ErrOnlyStructAllowed)
	}

	nm, tm := map[string]int{}, map[string]int{}
	for i := 0; i < tv.NumField(); i++ {
		field := tv.Type().Field(i)
		nm[field.Name] = i
		aim := field.Tag.Get(s.tagName)
		if aim != "" {
			tm[aim] = i
		}
	}

	for i := 0; i < fv.NumField(); i++ {
		aim := fv.Type().Field(i).Tag.Get(s.tagName)
		j, ok := tm[aim]
		if !ok || aim == "" {
			if !s.useMemberName {
				continue
			}
			aim = fv.Type().Field(i).Name
			j, ok = nm[aim]
			if !ok {
				continue
			}
		}

		ff, tf := fv.Field(i), tv.Field(j)

		iu := s.ignoreUnsigned && ignoreUnsigned(ff.Type(), tf.Type())
		ip := s.ignorePtr && ignorePtr(ff.Type(), tf.Type())
		if ff.Kind() != tf.Kind() && !iu && !ip {
			log.Printf("[Switch]fv: %s; tv: %s", ff.Kind(), tf.Kind())
			if s.skipTypeMismatch {
				continue
			}
			return f, s.throw(ctx, ErrMemberTypeDiff)
		}
		if iu {
			ff.Set(tf.Convert(ff.Type()))
			continue
		} else if ip {
			if tf.Kind() == reflect.Pointer {
				ff.Set(tf.Elem())
			} else {
				nt := reflect.New(tf.Type())
				nt.Elem().Set(tf)
				tf.Set(nt)
			}
			continue
		}
		ff.Set(tf)
	}
	return f, nil
}

func (s *switcher[F, T]) throw(ctx context.Context, err error) error {
	log.Printf("[tag.Switcher]err was throw: %s", err)
	if s.errable {
		return err
	}
	return nil
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

func ignoreUnsigned(source, target reflect.Type) bool {
	k, ok := igm[source.Kind()]
	if !ok {
		return false
	}
	return k == target.Kind()
}

func ignorePtr(source, target reflect.Type) bool {
	return (source.Kind() == reflect.Pointer && source.Elem().Kind() == target.Kind()) ||
		(target.Kind() == reflect.Pointer && target.Elem().Kind() == source.Kind())
}
