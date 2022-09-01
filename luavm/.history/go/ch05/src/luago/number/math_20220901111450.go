package number

import "math"

func IFloorDiv(a, b int64) int64 {
	if a > 0 && b > 0 || a < 0 && b < 0 || a%b == 0 {
		return a / b
	} else {
		return a/b - 1
	}
}

func FFloorDiv(a, b float64) float64 {
	return math.Floor(a / b)
}

func IMod(a, b int64) int64 {
	return a - IFloorDiv(a, b)*b
}

func FMOd(a, b float64) float64 {
	return a - math.Floor(a/b)*b
}

//左移运算
func ShiftLeft(a, n int64) int64 {
	if n >= 0 {
		return a << uint64(n)
	} else {
		return ShiftRight(a, -n) //因为Go语言中位移运算符右边的数只能是无符号整数，所以这里做个类型转换
	}
}

//右移运算
func ShiftRight(a, n int64) int64 {
	if n >= 0 {
		return int64(a >> uint64(n))
	} else {
		return ShiftLeft(a, -n)
	}
}

//浮点数转换为整数
func FloatToInteger(f float64) (int64, bool) {
	i := int64(f)
	return i, float64(i) == f
}
