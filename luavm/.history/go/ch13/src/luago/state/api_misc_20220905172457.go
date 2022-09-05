package state

func (self *luaState) Error() int {
	err := self.stack.pop()
	panic(err)
}
