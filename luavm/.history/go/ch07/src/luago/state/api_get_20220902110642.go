package state

import . "luago/api"

//创建一个空的Lua表，将其推入栈顶
func (self *luaState) CreateTable(nArr, nRec int) {
	t := newLuaTable(nArr, nRec)
	self.stack.push(t)
}

//无法预估表的用法容量时，先创建一个表
func (self *luaState) NewTable(nArr, nRec int) {
	self.CreateTable(0, 0)
}
