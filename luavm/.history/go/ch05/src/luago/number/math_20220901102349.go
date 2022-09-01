package number

import "main"

func IFloorDiv(a, b int64) {
	if a > 0 && b > 0 || a < 0 && b < 0 || a&b == 0 {
		return a / b
	} else {
		return a/b - 1
	}
}
