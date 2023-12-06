// Package tag
// Author: hyphen
// Copyright 2023 hyphen. All rights reserved.
// Create-time: 2023/7/21
package tag

import (
	"context"
	"encoding/json"
	"testing"
)

type From struct {
	Name   string  `json:"name" switch:"Name"`
	Age    int64   `json:"age" switch:"Age"`
	Gender bool    `json:"gender" switch:"Gender"`
	Ptr    *string `json:"ptr" switch:"Ptr"`
}

type To struct {
	Name   string `json:"name"`
	Age    uint64
	Gender *bool
	Ptr    string
}

func TestSwitcher(t *testing.T) {
	sw := NewSwitcher[From, To](WithErrable(), WithIgnoreUnsigned(), WithIgnorePtr())
	f := From{
		Name:   "ikun",
		Age:    2,
		Gender: false,
		Ptr:    nil,
	}
	tt, err := sw.Switch(context.Background(), &f)
	f.Name = "kkk"
	f.Gender = true
	t.Log(err)
	t.Log(MustMarshalIndent(context.Background(), tt, "\t"))
}

func MustMarshalIndent(ctx context.Context, src interface{}, indent string) string {
	res, err := json.MarshalIndent(src, "", indent)
	if err != nil {
		return ""
	}
	return string(res)
}
