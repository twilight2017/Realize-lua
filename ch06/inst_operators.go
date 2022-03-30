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
