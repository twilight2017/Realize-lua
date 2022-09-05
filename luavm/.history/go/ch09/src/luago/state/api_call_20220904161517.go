func (self *luaState) Call(nArgs, nResults int) {
	val := self.stack.get(-(nArgs + 1))
}