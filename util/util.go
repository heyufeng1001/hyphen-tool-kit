// Package util
// Author: hyphen
// Copyright 2022 hyphen. All rights reserved.
// Create-time: 2022/7/8
package util

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
