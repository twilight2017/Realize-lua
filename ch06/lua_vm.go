package ch06

type LuaVM interface {
	luaState
	PC() int          //返回当前pc
	AddPC(n int)      //修改PC(用于实现跳转指令)
	Fetch() uint32    //取出当前指令，将PC指向下一条指令
	GetConst(idx int) //将指定常量推入栈顶
	GetRK(rk int)     //将指定常量或栈值推入栈顶
}

func (self *luaState) PC() int{
	return self.pc
}

func AddPC(self *luaState) AddPC(n int) int{
	return self.pc += n
}

//fetch从函数原型的指令表中取出当前指令，然后把PC+1
func(self *luaState) Fetch() uint32{
	i := self.proto.Code[self.pc]
	self.pc++
	return i
}

//GetConst()从函数原型的常量表中取出一个常量值，然后把它推入栈顶
func(self *luaState) GetConst(idx int){
	c := self.proto.Constants[idx]
	self.stack.push(c)
}

//GetRK把某个常量或某个索引处的值推入栈顶
func (self *luaState) GetRK(rk int){
	if rk > 0xFF{
		//constant
		self.GetConst(rk & 0xFF)
	}else{
		//register
		self.PushValue(rk + 1)//因为寄存器索引从0开始，而Lua API里的索引是从1开始的
		}
}
