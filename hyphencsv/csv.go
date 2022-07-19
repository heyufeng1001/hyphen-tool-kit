// Package hyphencsv
// Author: hyphen
// Copyright 2022 hyphen. All rights reserved.
// Create-time: 2022/7/18
package hyphencsv

import (
	"encoding/csv"
	"fmt"
	"os"
)

func ParseCSV(path string) (*csv.Reader, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("[parseCSV]bytetreereader file failed: %w", err)
	}
	reader := csv.NewReader(file)
	return reader, nil
}
