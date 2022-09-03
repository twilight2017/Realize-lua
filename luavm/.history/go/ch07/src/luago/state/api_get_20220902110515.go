package state

import . "luago/api"

//创建一个空的Lua表，将其推入栈顶
func (self *luaState) CreateTable(nArr, nRec int) {
	t := newLuaTable(nArr, nRec)
	self.stack.push(t)
}
