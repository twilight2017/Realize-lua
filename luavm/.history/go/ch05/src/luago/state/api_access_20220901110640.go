package state

func (self *luaState) ToNumberX(idx int) (float64, bool) {
	val := self.stack.get(idx)
	return converToFloat(val)
}
