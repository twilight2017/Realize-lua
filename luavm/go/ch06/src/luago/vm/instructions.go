package vm

import "luago/api"

func (self Instruction) Execute(vm api.LuaVM) {
	action := opcodes[self.Opcode()].action
	if action != nil {
		action(self, vm)
	} else {
		panic(self.OpName())
	}
}
