package state

func (self *luaState) PushGoFunction(f GoFunction) {
	self.stack.push(new)
}
