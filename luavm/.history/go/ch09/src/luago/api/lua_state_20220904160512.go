//返回的整数说明了返回值的个数
package api

type LuaState interface {
	//把Go函数转变为Go闭包并推入Lua栈中
	PushGoFunction(f GoFunction)
}
type GoFunction func(LuaState) int
