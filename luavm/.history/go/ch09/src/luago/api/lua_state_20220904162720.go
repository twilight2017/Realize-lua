//返回的整数说明了返回值的个数
package api

type luaState interface {
	//把Go函数转变为Go闭包并推入Lua栈中
	PushGoFunction(f GoFunction)
	IsGoFunction(idx int) bool
	ToGoFunction(idx int) GoFunction
	registry *luaTable //注册表，用户可以在里面存放任何lua值
    stack *luaStack
}
func New() *luaState{
	registry := newLuaTable(0, 0)
	registry.put(LUA_RIDX_GLOBALS, newLuaTable(0, 0)) //全局环境

	ls := &luaState{registry: registry}
	ls.pushLuaStack(newLuaStack(LUA_MINSTACK, ls))
}
type GoFunction func(LuaState) int
