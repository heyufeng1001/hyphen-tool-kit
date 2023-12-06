// Package unsafe
// Author: hyphen
// Copyright 2023 hyphen. All rights reserved.
// Create-time: 2023/7/18
package unsafe_test

import (
	"fmt"
	"testing"

	"code.byted.org/lark-devops/deploy_cd/utils/switcher/unsafe"
)

type Kun struct {
	Year  int
	Month string
}

func (k *Kun) Sing() {
	fmt.Println("jinitaimei")
}

func (k *Kun) Dance() {
	fmt.Println("xiangchilaofan")
}

type Chicken struct {
	Year  uint
	Month string
}

func (c *Chicken) Rap() {
	fmt.Println("xiangjinjianyu")
}

var sw = unsafe.NewSwitcher[Kun, Chicken](true)

func TestSwitcher(t *testing.T) {
	cxk := &Kun{
		Year:  2,
		Month: "half",
	}
	cxk.Sing()
	ikun := sw.Switch(cxk)
	ikun.Rap()
	t.Log(ikun)
}
