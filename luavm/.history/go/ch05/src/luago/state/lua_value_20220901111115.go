package state

import (
	"luago/number"
)

type luaValue interface{}

func convertToFloat(val luaValue) (float64, bool) {
	switch x := val.(type) {
	case float64:
		return x, true
	case int64:
		return float64(x), true
	case string:
		return number.ParseFloat(x)
	default:
		return 0, false
	}
}

func convertToInteger(val luaValue) (int64, bool) {
	switch x := val.(type) {
	case int64:
		return x, true
	case float64:
		return int64(x), true
	case string:
		return _strintToInteger(x)
	default:
		return 0, false
	}
}

func _strintToInteger(s string) (int64, bool) {
	if i, ok := number.ParseInterger(s); ok {
		return i, true
	}
}
