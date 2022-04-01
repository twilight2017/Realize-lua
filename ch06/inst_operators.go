package ch06

func _binaryArith(i Instruction, vm LuaVM, op ArithOp) {
	a, b, c = i.ABC()
	a += 1

	vm.GetRK(b) //将b推入堆栈
	vm.GetRK(c)
	vm.Arith(op)  //进行算术运算
	vm.Replace(a) //用值替换指定索引处的值
}

func _unaryArith(i Instruction, vm LuaVM, op ArithOp) {
	a, b, _ = i.ABC()
	a += 1
	b += 1
	vm.PushValue(b)
	vm.Arith(op)
	vm.Replace(a)
}

//实现全部算术指令
func add(i Instruction, vm LuaVM) {
	_binaryArith(i, vm, LUA_OPADD)
} //+

//-
func sub(i Instruction, vm LuaVM) {
	_binaryArith(i, vm, LUA_OPSUB)
}

//*
func mul(i Instruction, vm LuaVM) {
	_binaryArith(i, vm, LUA_OPMUL)
}

//%
func mod(i Instruction, vm LuaVM) {
	_binaryArith(i, vm, LUA_OPMOD)
}

//^
func pow(i Instruction, vm LuaVM) {
	_binaryArith(i, vm, LUA_OPPOW)
}

// /
func div(i Instruction, vm LuaVM) {
	_binaryArith(i, vm, LUA_OPDIV)
}

// //
func idiv(i Instruction, vm LuaVM) {
	_binaryArith(i, vm, LUA_OPIDIV)
}

// &
func band(i Instruction, vm LuaVM) {
	_binaryArith(i, vm, LUA_OPBAND)
}

// |
func bor(i Instruction, vm LuaVM) {
	_binaryArith(i, vm, LUA_OPBOR)
}

// -
func bxor(i Instruction, vm LuaVM) {
	_binaryArith(i, vm, LUA_OPBXOR)
}

// <<
func shl(i Instruction, vm LuaVM) {
	_binaryArith(i, vm, LUA_OPSHL)
}

// >>
func shr(i Instruction, vm LuaVM) {
	_binaryArith(i, vm, LUA_OPSHR)
}

// -
func unm(i Instruction, vm LuaVM) {
	_unaryArith(i, vm, LUA_UNM)
}

//-
func bnot(i Instruction, vm LuaVM) {
	_unaryArith(i, vm, LUA_BNOT)
}

//长度计算函数
func _len(i Instruction, vm LuaVM) {
	a, b, _ = i.ABC()
	a += 1
	b += 1

	vm.Len(b)
	vm.Replace(a)
}

//concat():将连续n个寄存器（起止索引分别由操作数B和操作数C指定）的值进行拼接
//将计算结果放入另一个寄存器（索引由操作数A决定）
func concat(i Instruction, vm LuaVM) {
	a, b, c = i.ABC()
	a += 1
	b += 1
	c += 1
	n := c - b + 1   //获取n的长度
	vm.CheckStack(n) //检查栈中剩余空间是否有n个
	for i := b; i <= c; i++ {
		vm.PushValue(i)
	}
	vm.Contact(n)
	vm.Repalce(a)
}

//比较指令：比较寄存器或常量表里的两个值（索引分别由操作数B和C指定）
//如何比较结果和操作数A（转换为布尔值）匹配，则跳过下一条指令
func _compare(i Instruction, vm LuaVM, op CompareOp) {
	a, b, c = i.ABC()

	vm.GetRK(b)
	vm.GetRK(c) //将值推入栈顶
	if vm.Compare(-2, -1, op) != (a != 0) {
		vm.AddPC(1)
	}
	vm.Pop(2) //清理栈顶
}

//==
func eq(i Instruction, vm LuaVM) {
	_compare(i, vm, LUA_OPEQ)
}

//<
func lt(i Instruction, vm LuaVM) {
	_compare(i, vm, LUA_OPLT)
}

//<=
func le(i, Instruction, vm LuaVM) {
	_compare(i, vm, LUA_OPLE)
}
