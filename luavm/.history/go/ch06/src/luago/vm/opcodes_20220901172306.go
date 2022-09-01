package vm

import "luago/api"

type opcode struct {
	action func(i Instruction, vm api.LuaVm)
}
