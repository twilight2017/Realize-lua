package vm

func setList(i Instruction, vm LuaVM) {
	a, b, c := i.ABC()
	a += 1
	if bIsZero {
		b = int(vm.ToInteger(-1)) - a - 1
		vm.Pop(1)
	}
}
