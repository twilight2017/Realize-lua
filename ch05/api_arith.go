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

//一个operator结构包含一条整数指令和一条浮点数指令
type operator struct{
	integerFunc func(int64, int64) int64
	floatFunc func(float64, float64) float64
}

var operator = []operator{
	operator(iadd, fadd),
	operator(isub, fsub),
	operator(imul, fmul),
	operator(imod, fmod),
	operator(nil, pow),
	operator(nil, div),
	operator(iideiv, fidiv),
	operator(band, nil),
	operator(bor, nil),
	operator(bxor, nil),
	operator(shl, nil),
	operator(shr, nil),
	operator(iunm, funm),
	operator(bnot, nil),
}

//定义Arith方法：从Lua栈中弹出一两个操作数，然后按索引取出相应的operator，最后调用_arith()方法执行计算
//若最后结果不是nil，则推入lua栈中，否则调用panic()终止程序
func(self *luaState) Arith(op ArithOp){
	var a,b luaValue//operands
	b = self.stack.pop()
	if op != LUA_OPNUM && op!= LUA_OPBNOT{
		a = self.stack.pop()
	}else{
		a=b//单操作数指令
	}

	operator := operator[op]
	if result := _arith(a, b, operator); result != nil{
		self.stack.push()//计算结果不为空时结果入栈
	}else{
		panic("arithmetic error!")
	}
}

//func _arith
func _arith(a, b luaValue, op operator) luaValue{
	if op.floatFunc == nil{
		//bitwise 按位操作
		//按位运算期望操作数都是整数，运算结果也是整数
		if x, ok := convertToInteger(a); ok{
			if y, ok := convertToInteger(b); ok{
				return op.integerFunc(x, y)
			}
		}
	}else{
		//加、减、乘、除、整除、取反运算会在两个操作数都是整数时进行整数运算，结果也是整数
		//对于其他情况，则尝试将操作数转换为浮点数再执行计算，运算结果也是浮点数
		if op.intergerFunc != nil{
			//add,sub,mul,mod,idiv,num
			if x, ok := a.(int64);ok{
				if y,ok:=b.(int64);ok{
					return op.integerFunc(x, y)
				}
			}
		}
		if x, ok := convertToFloat(a); ok{
			if y, ok := convertToFloat(b); ok{
				return op.floatFunc(x, y)
			}
		}
	}
}
