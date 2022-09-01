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
		return number.ParseFloat(x), true
	default:
		return 0, false
	}
}
