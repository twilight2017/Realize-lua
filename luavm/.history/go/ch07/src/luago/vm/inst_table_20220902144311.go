package vm

import . "luago/api"
import "luago/number"

func newTable(i Instruction, vm LuaVm){
	a, b, c := i.ABC()
	a += 1

	vm.CreateTable(Fb2int(b), Fb2int(c))
}
