package state

func (self *luaStack) isValid(idx int) bool {
	if idx < LUA_REGISTRYINDEX {
		/* upvalues */
		uvIdx := LUA_REGISTRYINDEX - idx - 1
		c := self.closure
		return c != nil && uvIdx < len(c.upvals)
	}
}
