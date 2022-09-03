package state

//键和值从栈里弹出，表位于指定索引处
func (self *luaState) SetTable(idx int) {
	t := self.stack.get(idx) //获得表
	v := self.stack.pop()
	k := self.stack.pop()
	self.setTable(t, k, v)
}

func (self *luaState) setTable(t, k, v luaValue)
