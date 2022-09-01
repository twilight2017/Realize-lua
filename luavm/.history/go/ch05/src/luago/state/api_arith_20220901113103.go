package state

import "math"
import . "luago/api"
import "luago/number"

//将匿名函数赋值给变量
var (
	iadd = func(a, b int) int { return a + b }
)
