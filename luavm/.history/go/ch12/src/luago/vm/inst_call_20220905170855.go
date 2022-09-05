package vm

func tForCall(i Instruction, vm LuaVM) {
	a, _, c := i.ABC()
	a += 1
}
