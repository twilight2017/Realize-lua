package state

func (self *luaState) Len(idx int) {
	val := self.stack.get(idx)
}
