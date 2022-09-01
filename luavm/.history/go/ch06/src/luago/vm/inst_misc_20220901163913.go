package vm

import "luago/api"

func move(i Instruction, vm LuaVM) {
	a, b, _ = i.ABC() //获取两个操作数
	a += 1
	b += 1
	vm.Copy(b, a)
}
