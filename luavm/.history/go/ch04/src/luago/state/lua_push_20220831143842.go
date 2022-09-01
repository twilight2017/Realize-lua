package state

func (self *luaState) PushNil() {
	self.stack.push(nil)
}

func (self *luaState) PushBoolean(b bool) {
	self.stack.push(b)
}
