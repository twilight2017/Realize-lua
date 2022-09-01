package vm

import "luago/api"

func _binaryArith(i Instruction, vm LuaVM, op ArithOp) {
	a, b, c := i.ABC()
	a += 1

	vm.GetRK(b)
	vm.GETRK(c)
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
