//返回的整数说明了返回值的个数
package api

type LuaState interface {
	PushGoFunction(f GoFunction)
}
type GoFunction func(LuaState) int
