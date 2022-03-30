package ch06

//LOADNIL 给连续n个寄存器存放nil值
func loadNil(i Instruction, vm LuaVM) {
	a, b, _ = i.ABC()
	a += 1

	vm.PushNil()
	for i := a; i < a+b; i++ {
		vm.Copy(-1, i)
	}
	vm.Pop(1)
}

//LOADBOOL 给单个寄存器设置bool值
//寄存器索引：由操作数A指定
//bool值：由寄存器B指定（0->false not 0->true）
func loadBool(i Instruction, vm LuaVM) {
	a, b, c = i.ABC() //获取三个操作数
	a += 1
	vm.PushBoolean(b != 0)
	vm.Replace(a)
	if c != 0 {
		vm.AddPC(1)
	}
}
