package main

import (
	"strconv"
	"strings"
)

func ParsePercentage(s string) (float64, error) {
	s = strings.ReplaceAll(s, "%", "")

	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return -1, err
	}

	return f, nil
}
