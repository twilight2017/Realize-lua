package state

func (self *luaState) PC() int {
	return self.pc
}

func (self *luaState) AddPC(n int) {
	self.pc += n
}

//从函数原型的指令表中取出相应的指令
func (self *luaState) Fecth() uint32 {
	i := self.proto.Code[self.pc] //从Code变量中取出对应的指令
	self.pc++
	return i
}

//从函数原型的常量表中取出对应的常量
func (self *luaState) GetConst(idx int) {
	c := self.proto.Constants[idx]
	return c
}
