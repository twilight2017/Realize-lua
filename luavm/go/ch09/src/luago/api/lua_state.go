//返回的整数说明了返回值的个数
package api

type luaState interface {
	//把Go函数转变为Go闭包并推入Lua栈中
	PushGoFunction(f GoFunction)
	IsGoFunction(idx int) bool
	ToGoFunction(idx int) GoFunction
	registry *luaTable //注册表，用户可以在里面存放任何lua值
	PushGlobalTable()
	GetGloabal(name string)
	SetGlobal(name string)
	Register(name string, f GoFunction)
    stack *luaStack
}
func New() *luaState{
	registry := newLuaTable(0, 0)
	registry.put(LUA_RIDX_GLOBALS, newLuaTable(0, 0)) //全局环境, 其也是一个普通的lua表

	ls := &luaState{registry: registry}  //用表来创建luaState结构体实例
	ls.pushLuaStack(newLuaStack(LUA_MINSTACK, ls)) //随后将其推入一个新的Lua栈(调用帧)
}
type GoFunction func(LuaState) int
