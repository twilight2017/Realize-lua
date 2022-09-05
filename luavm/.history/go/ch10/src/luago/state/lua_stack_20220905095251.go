package state

func (self *luaStack) isValid(idx int) bool {
	if idx < LUA_REGISTRYINDEX {
		/* upvalues */
	}
}
