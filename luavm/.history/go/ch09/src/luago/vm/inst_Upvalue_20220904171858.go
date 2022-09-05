package vm

import . "luago/api"

func getTabUp(i Instruction, vm LuaVM) {
	a, _, c = i.ABC()
	a += 1

	vm.PushGlobalTable()
}
