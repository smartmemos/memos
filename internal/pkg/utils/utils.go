package utils

import "encoding/json"

func JSONStr(v any) string {
	data, _ := json.Marshal(v)
	return string(data)
}

func Unique[T comparable](arr []T) []T {
	if len(arr) == 0 {
		return arr
	}
	m := make(map[T]struct{})
	ret := make([]T, 0, len(arr))
	for _, v := range arr {
		if _, ok := m[v]; !ok {
			ret = append(ret, v)
			m[v] = struct{}{}
		}
	}
	return ret
}

func Ptr[T any](v T) *T {
	return &v
}
