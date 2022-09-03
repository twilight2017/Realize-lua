package vm

func setList(i Instruction, vm LuaVM) {
	a, b, c := i.ABC()
	a += 1
	bIsZero := b == 0
	if bIsZero {
		for j := vm.RegisterCount() + 1; j <= vm.GetTop(); j++ {
			idx++
			vm.PushValue(j)
			vm.SetI(a, idx)

		}
		vm.SetTop(vm.RegisterCount(0))
	}
}
