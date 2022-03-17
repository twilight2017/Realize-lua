package ch04

type LuaState struct {
	stack *luaStack
}

//new()用于创建luaState实例
func New() *luaState {
	return &luaState{
		stack: newLuaStack(20),
	}
}
