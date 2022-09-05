package state

type operator struct {
	metamethod  string
	integerFunc func(int64, int64) int64
	floatFunc   func(float64, float64) float64
}

var operators = []operator{
	operator{"_add", iadd, fadd},
	operator{"_sub", isub, fsub},
	operator{"_mul", imul, fmul},
	operator{"_mod", imod, fmod},
	operator{"_pow", nil, pow},
	operator{"_div", nil, div},
	operator{"_idiv", iidiv, fidiv},
	operator{"_band", band, nil},
	operator{"_bor", bor, nil},
	operator{"_bxor", bxor, nil},
	operator{"_shl", shl, nil},
	operator{"_shr", shr, nil},
	operator{"_unm", iunm, funm},
	operator{"_bnot", bnot, nil},
}

func (self *luaState) Arith(op ArithOp) {
	if result := _arith(a, b, operator); result != nil {
		self.stack.push(result)
		return
	}
	mm := operator.metamethod
	if result, ok := callMetamethod(a, b, mm, self); ok {
		self.stack.push(result)
		return
	}
	panic("arithmetic error!")
}
