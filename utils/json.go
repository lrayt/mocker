package utils

import "github.com/tidwall/gjson"

func JsonValueWithDefault[T string | uint64 | float64 | bool](data []byte, key string, val T) T {
	result := gjson.GetBytes(data, key)
	if !result.Exists() {
		return val
	}
	switch any(val).(type) {
	case string:
		return any(result.String()).(T)
	case float64:
		return any(result.Float()).(T)
	case bool:
		return any(result.Bool()).(T)
	case uint64:
		return any(result.Uint()).(T)
	default:
		return val
	}
}
