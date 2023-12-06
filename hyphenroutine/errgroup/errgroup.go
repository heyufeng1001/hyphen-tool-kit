// Package errgroup
// Author: hyphen
// Copyright 2023 hyphen. All rights reserved.
// Create-time: 2023/2/13
package errgroup

import (
	"context"

	"golang.org/x/sync/errgroup"
)

type waitOpt bool

const (
	FinishNow waitOpt = false
	WaitAll   waitOpt = true
)

type HyphenErrGroup interface {
	Go(f func() error)
	SetLimit(n int)
	TryGo(f func() error) bool
	Wait() error
}

func HyphenErrGroupWithCtx(ctx context.Context, wait waitOpt) (HyphenErrGroup, context.Context) {
	switch wait {
	case WaitAll:
		return errgroup.WithContext(ctx)
	default:
		return withContext(ctx)
	}
}

type nweGroup struct {
	*errgroup.Group
	context.Context
	errSig chan error
}

func withContext(ctx context.Context) (*nweGroup, context.Context) {
	eg, ctx := errgroup.WithContext(ctx)
	return &nweGroup{eg, ctx, make(chan error)}, ctx
}

func (g *nweGroup) Go(f func() error) {
	g.Group.Go(func() error {
		defer func() {
			if err := recover(); err != nil {
				return
			}
		}()
		err := f()
		if err != nil {
			g.errSig <- err
			return err
		}
		return nil
	})
}

func (g *nweGroup) SetLimit(n int) {
	g.Group.SetLimit(n)
}

func (g *nweGroup) TryGo(f func() error) bool {
	return g.Group.TryGo(f)
}

func (g *nweGroup) Wait() error {
	var err error
	wch := make(chan int)
	go func() {
		defer func() {
			if err := recover(); err != nil {
				return
			}
		}()
		err = g.Group.Wait()
		wch <- 1
	}()
	select {
	case e := <-g.errSig:
		return e
	case <-wch:
		return err
	}
}
