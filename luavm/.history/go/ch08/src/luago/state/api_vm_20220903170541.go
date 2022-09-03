package state

func (self *luaState) PC() int{
	return self.stack.pc
}

func (self *luaState) AddPC (n int){
	return self.stack.pc += n
}

func (self *luaState) Fetch() uint32{
	i := self.stack.closure.proto.Code[self.stack.pc]
	self.stack.pc++
	return i
}

func (self *luaState) GetConst(idx int){
	c := self.stack.closure.proto.Constants[idx]
	self.stack.push(c)
}

//返回当前Lua函数函数所操作的寄存器数量
func (self *luaState) RegisterCount() int{
	return int(self.stack.closure.proto.MaxStackSize)
}

//把传递给当前Lua函数的变长参数推入栈顶
func (self *luaState) LoadVararg(n int){
	if n < 0{
		n = len(self.stack.varargs)
	}
	self.stack.check(n)
	self.stack.pushN(self.stack.varargs, n)
}