package common

import (
	"strconv"
	"strings"
)

func SplitInt64(s string) []int64 {
	if s == "" {
		return nil
	}
	parts := strings.Split(s, ",")
	res := make([]int64, 0, len(parts))
	for _, part := range parts {
		if val, err := strconv.ParseInt(part, 10, 64); err == nil {
			res = append(res, val)
		}
	}
	return res
}
