package util

import (
	"strconv"
)

// Float64 转换float64
func Float64(any interface{}) float64 {
	if any == nil {
		return 0
	}
	switch value := any.(type) {
	case float64:
		return value
	case float32:
		return float64(value)
	case uint:
		return float64(value)
	case uint8:
		return float64(value)
	case uint32:
		return float64(value)
	case uint16:
		return float64(value)
	case int:
		return float64(value)
	case int8:
		return float64(value)
	case int16:
		return float64(value)
	case int32:
		return float64(value)
	case int64:
		return float64(value)
	case string:
		v, _ := strconv.ParseFloat(value, 64)
		return v
	default:
		return 0
	}
}
