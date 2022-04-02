package ch06

func (self Instruction) Execute(vm api.LuaVM) {
	action := opcodes[self.Opcode()].action //通过指令表，获取指令对应操作
	if action != nil {
		action(self, vm)
	} else {
		panic(self.OpName())
	}
}
