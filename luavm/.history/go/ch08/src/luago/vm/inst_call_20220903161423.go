package vm

import . "luago/api"

//closure把当前Lua函数的子函数原型实例化为闭包，放入由操作数A指定的寄存器中
func closure(i Instruction, vm LuaVM) {
	a, bx := i.ABx()
	a += 1

	vm.LoadProto(bx)
	vm.Replace(a)
}
