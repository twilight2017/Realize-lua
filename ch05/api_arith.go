package state

import "math"
import "luago/api"
import "luago/number"

var(
	iadd = func(a, b int64) int64{
		return a+b
	}
	fadd = func(a, b float64) float64{
		return a+b
	}
	isub = func(a, b int64) int64{
		return a-b
	}
	fsub = func(a, b float64) float64{
		return a-b
	}
	imul = func(a, b int64) int64{
		return a*b
	}
	fmul = func(a, b float64) float64{
		return a * b
	}
	imod = iMod
	fmod = fMod
	pow = Pow
    div = func(a, b float64) float64{
		return a/b
	}
	iidiv = IFloorDiv
	fidiv = FFloorDiv
	band = func(a, b int64) int64{
		return a & b
	}
	bor = func(a, b int64) int64{
		return a|b
	}
	bxor = func(a, b int64) int64{
		return a ^ b
	}
	shl=ShiftLeft
	shr=ShiftRight
	iunm = func(a, _ int64) int64{
		return -a
	}
	funm = func(a, _ float64) float64{
		return -a
	}
	bnot = func(a, _ int64) int64 
	{
		return ^a
	}
)

type operator struct{
	integerFunc func(int64, int64) int64
	floatFunc func(float64, float64) float64
}