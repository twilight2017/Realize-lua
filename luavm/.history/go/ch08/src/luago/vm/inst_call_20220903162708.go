package vm

import . "luago/api"

//closure把当前Lua函数的子函数原型实例化为闭包，放入由操作数A指定的寄存器中
func closure(i Instruction, vm LuaVM) {
	a, bx := i.ABx()
	a += 1

	vm.LoadProto(bx)
	vm.Replace(a)
}

func call(i Instruction, vm LuaVM) {
	a, b, c := i.ABC()
	a += 1

	nArgs := _pushFuncAndArgs(a, b, vm)
	vm.Call(nArgs, c-1)
	_popResults(a, c, vm)
}

func _pushFuncAndArgs(a, b int, vm LuaVM) (nArgs int) {
	if b >= 1 {
		vm.CheckStack(b)
		for i := a; i < a+b; i++ {
			vm.PushValue(i)
		}
		return b - 1
	}
}

func _popResults(a, c int, vm LuaVM) {
	if c == 1 {

	} else if c > 1 {
		for i := a + c - 2; i >= a; i-- {
			vm.Replace(i)
		}
	} else {
		vm.CheckStack(1)
		vm.PushInteger(int64(a))
	}
}
