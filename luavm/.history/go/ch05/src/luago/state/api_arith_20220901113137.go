package state

import "math"
import . "luago/api"
import "luago/number"

//将匿名函数赋值给变量
var (
	iadd = func(a, b int64) int64 { return a + b }
	fadd = func(a, b float64) float64 { return a + b }
)
