// Package hyphensort
// Author: hyphen
// Copyright 2022 hyphen. All rights reserved.
// Create-time: 2022/8/17
package hyphensort

func QuickSort[T any](data []T, less func(first, second T) bool) {
	if len(data) == 0 {
		return
	}
	i, j, base := 0, len(data)-1, data[0]
	for i < j {
		for !less(data[i], base) && i < j {
			j--
		}
		if i != j {
			data[i], data[j] = data[j], data[i]
			i++
		}
		for less(data[i], base) && i < j {
			i++
		}
		if i != j {
			data[i], data[j] = data[j], data[i]
			j--
		}
	}
	data[i] = base
	QuickSort(data[:i], less)
	QuickSort(data[i+1:], less)
}
