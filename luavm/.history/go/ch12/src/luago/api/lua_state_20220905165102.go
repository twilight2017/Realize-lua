package api

type LuaState interface {
	Next(idx int) bool
}
