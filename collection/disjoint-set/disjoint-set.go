// Package disjoint_set
// Author: hyphen
// Copyright 2022 hyphen. All rights reserved.
// Create-time: 2022/8/5
package disjoint_set

import (
	"fmt"

	"github.com/heyufeng1001/hyphen-tool-kit/hyphenutil"
)

type DisjointSet[T comparable] struct {
	parent   map[T]T
	rank     map[T]int
	enablePC bool // enable path compression
}

func NewDisjointSet[T comparable](values []T, enablePC ...bool) *DisjointSet[T] {
	parent, rank := map[T]T{}, map[T]int{}
	for _, value := range values {
		parent[value] = value
		rank[value] = 0
	}
	return &DisjointSet[T]{
		parent, rank, hyphenutil.ParseVariableBools(enablePC...),
	}
}

func (d *DisjointSet[T]) Find(value T) (T, error) {
	if !d.IsExist(value) {
		return nil, fmt.Errorf("[*DisjointSet.Find]value doesn't exist in disjoint set")
	}
	return d.innerFind(value), nil
}

func (d *DisjointSet[T]) innerFind(value T) T {
	return hyphenutil.TernaryForm(d.enablePC, func(v T) T {
		return hyphenutil.TernaryForm(d.parent[v] != v, func() T {
			d.parent[v] = d.innerFind(d.parent[v])
			return v
		}(), d.parent[v])
	}, func(v T) T {
		return hyphenutil.TernaryForm(d.parent[v] == v, v, d.innerFind(d.parent[v]))
	})(value)
}

func (d *DisjointSet[T]) IsExist(value T) bool {
	_, ok := d.parent[value]
	return ok
}

func (d *DisjointSet[T]) Merge(a, b T) (bool, error) {
	if !d.IsExist(a) {
		return false, fmt.Errorf("[*DisjointSet.Merge]a doesn't exist in disjoint set")
	}
	if !d.IsExist(b) {
		return false, fmt.Errorf("[*DisjointSet.Merge]b doesn't exist in disjoint set")
	}
	aRoot, bRoot := d.innerFind(a), d.innerFind(b)
	if aRoot == bRoot {
		return false, nil
	}
	if d.rank[aRoot] > d.rank[bRoot] {
		d.parent[bRoot] = aRoot
	} else if d.rank[aRoot] < d.rank[bRoot] {
		d.parent[aRoot] = bRoot
	} else {
		d.rank[aRoot]++
		d.parent[bRoot] = aRoot
	}
	return true, nil
}

func (d *DisjointSet[T]) AppendNode(value T) error {
	if d.IsExist(value) {
		return fmt.Errorf("[*DisjointSet.AppendNode]value exists in disjoint set")
	}
	d.parent[value] = value
	return nil
}
