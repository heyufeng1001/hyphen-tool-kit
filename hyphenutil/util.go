// Package hyphenutil
// Author: hyphen
// Copyright 2022 hyphen. All rights reserved.
// Create-time: 2022/7/8
package hyphenutil

func TernaryForm[T any](condition bool, trueVal, falseVal T) T {
	if condition {
		return trueVal
	}
	return falseVal
}

func TernaryFormInterface(condition bool, trueVal, falseVal interface{}) interface{} {
	if condition {
		return trueVal
	}
	return falseVal
}

func ParseVariableBools(flag ...bool) (ret bool) {
	if len(flag) > 0 && flag[0] {
		ret = true
	}
	return
}
