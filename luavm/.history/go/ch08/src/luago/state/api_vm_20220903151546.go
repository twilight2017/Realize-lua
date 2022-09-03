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
}