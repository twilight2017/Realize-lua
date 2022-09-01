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
