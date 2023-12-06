// Package errgroup
// Author: hyphen
// Copyright 2023 hyphen. All rights reserved.
// Create-time: 2023/2/13
package errgroup

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestHyphenErrGroup(t *testing.T) {
	eg, _ := HyphenErrGroupWithCtx(context.Background(), WaitAll)
	eg.Go(func() error {
		time.Sleep(time.Second * 3)
		fmt.Println("123")
		return fmt.Errorf("123")
	})
	eg.Go(func() error {
		time.Sleep(time.Second * 6)
		fmt.Println("456")
		return fmt.Errorf("456")
	})
	if err := eg.Wait(); err != nil {
		fmt.Println("err catch: ", err)
	}
	//time.Sleep(time.Second * 9)
	fmt.Println("789")
}
