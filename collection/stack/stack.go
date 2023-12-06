// Package stack
// Author: hyphen
// Copyright 2022 hyphen. All rights reserved.
// Create-time: 2022/10/28
package stack

type Stack interface {
	Len() int
	Pop() (comparable, bool)
	Peek() (comparable, bool)
	Push(v comparable)
}
