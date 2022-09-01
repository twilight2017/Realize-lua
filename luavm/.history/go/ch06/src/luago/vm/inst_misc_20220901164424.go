package vm

import "luago/api"

func move(i Instruction, vm LuaVM) {
	a, b, _ = i.ABC() //获取两个操作数
	a += 1
	b += 1        //转换成栈索引
	vm.Copy(b, a) //该方法可以拷贝栈值
}

func jump(i Instruction, vm LuaVM) {
	a, sBx := i.AsBx()
	vm.AddPC(sBx)
	if a != 0 {
		panic("todo!")
	}
}

//把值加载到寄存器里
//LOADNIL，给连续n个寄存器存放nil值
func loadNil(i instruction, vm LuaVM){
	a, b, _ = i.ABC()
	a += 1
	i:=a;i<=a+b;i++{
		vm.Copy(-1,i)
	}
	vm.Pop(1)
}
