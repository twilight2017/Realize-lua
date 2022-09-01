package vm

import "luago/api"

type opcode struct {
	action func(i Instruction, vm api.LuaVm)
}

var opcodes = []opcode{
	opcode{0, 1, OpArgR, OpArgN, IABC, "MOVE    ", move},
	opcode{0, 1, OpArgR, OpArgN, IABx, "LOADK   ", loadK},
	opcode{0. 1, OpArgR, OpArgN, IABx, "LOADKX  ", loadKx},
	opcode{0, 1, OpArgU, OpArgU, IABC, "LOADBOOL", loadBool},
	opcode{0, 1, OpArgU, OpArgU, IABC, "LOADNIL ", loadnil},
}