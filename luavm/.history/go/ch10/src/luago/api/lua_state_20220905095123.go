package api

//先从栈顶弹出n个Lua值
type LuaState interface {
	PushGoClosure(f GoFunction, n int)
}

func LuaUpvalueIndex(i int) int {
	return LUA_REGISTRYINDEX - i
}
