package api

type LuaState interface {
	PushGoClosure(f GoFunction, n int)
}
