package state

func (self *luaStack) isValid(idx int) bool {
	if idx < LUA_REGISTRYINDEX {
		/* upvalues */
		uvIdx := LUA_REGISTRYINDEX - idx - 1
		c := self.closure
		return c != nil && uvIdx < len(c.upvals)
	}
}

func (self *luaStack) get(idx int) luaValue {
	if idx < LUA_REGISTRYINDEX {
		/*upvalues*/
		uvIdx := LUA_REGISTRYINDEX - idx - 1
		c := self.closure
		if c == nil || uvIdx >= len(c.upvals) {
			return nil
		}
		return *(c.upvals[uvIdx].val)
	}
}

func (self *luaStack) set(idx int, val luaValue) {
	if idx < LUA_REGISTRYINDEX {
		/* upvalues*/
	}
}
