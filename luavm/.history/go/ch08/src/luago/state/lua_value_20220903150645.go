package state

func typeOf(val luaValue) LuaType {
	switch val.(type) {
	case *closure:
		return LUA_TFUNCTION
	}
}
