func (self *luaState) PushGoFunction(f GoFunction) {
	self.stack.push(newGoClosure(f))
}