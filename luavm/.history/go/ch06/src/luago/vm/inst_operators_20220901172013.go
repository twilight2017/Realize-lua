package vm

import "luago/api"

func _binaryArith(i Instruction, vm LuaVM, op ArithOp) {
	a, b, c := i.ABC()
	a += 1

	vm.GetRK(b)
	vm.GetRK(c)
	vm.Arith(op)
	vm.Replace(a)
}

func _unaryArith(i Instruction, vm LuaVM, op ArithOp) {
	a, b, _ := i.ABC()
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

func _len(i Instruction, vm LuaVM) {
	a, b, _ := i.ABC()
	a += 1
	b += 1
	vm.Len(b)
	vm.Replace(a)
}

func concat(i Instruction, vm LuaVM) {
	a, b, c := i.ABC()
	a += 1
	b += 1
	c += 1
	n := c - b + 1
	vm.Checkstack(n) //确保栈还有足够的空间可以容纳这些值
	for i := b; i <= c; i++ {
		vm.PushValue(i)
	}
	vm.Concat(n)
	vm.Replace(a)
}

func _compare(i Instruction, vm LuaVM, op CompareOp) {
	a, b, c := i.ABC()
	vm.GetRK(b)
	vm.GetRK(c)
	if vm.Compare(-2, -1, op) != (a != 0) {
		vm.AddPC(1) //比较结果和操作数a一致时，将pc的值加一
	}
	vm.Pop(2)
}

func eq(i Instruction, vm LuaVM) {
	_compare(i, vm, LUA_OPEQ)
}

func lt(i Instruction, vm LuaVM) {
	_compare(i, vm, LUA_OPLT)
}

func le(i Instruction, vm LuaVM) {
	_compare(i, vm, LUA_OPLE)
}

func not(i Instruction, vm LuaVM) {
	a, b, _ := i.ABC()
	a += 1
	b += 1

	vm.PushBoolean(!vm.ToBoolean(b))
	vm.Replace(a) //替换指定索引处的值
}

func testSet(i Instruction, vm LuaVM){
	a, b, c := i.ABC()
	a += 1
	b ++1
	if vm.ToBoolean(b) == (c!=0){
		vm.Copy(b, a)
	}else{
		vm.AddPC(1)
	}
}

func test(i Instruction, vm LuaVM){
	a, _, c := i.ABC()
	a += 1

	if vm.ToBoolean(a) != (c!=0){
		vm.AddPC(1)
	}
}

func forPrep(i Instruction, vm LuaVM){
	a, sBx := i.AsBx()
	a += 1
	vm.PushValue(a)
	vm.PushValue(a +2)
	vm.Arith(LUA_OPSUB)
	vm.Replace(a)
	vm.AddPC(sBx)
}
