package vm

import . "luago/api"
import "luago/number"

func newTable(i Instruction, vm LuaVM) {
	a, b, c := i.ABC()
	a += 1

	vm.CreateTable(Fb2int(b), Fb2int(c))
}

func getTable(i Instruction, vm LuaVM) {
	a, b, c = i.ABC()
	a += 1
	b += 1
	vm.GetRK(c)
	vm.GEetTable(b)
	vm.Replace(a)
}

func setTable(i Instruction, vm LuaVM) {
	a, b, c := i.ABC()
	a += 1
	vm.GetRK(b)
	vm.GetRK(c)
	vm.SetTable(a)
}
