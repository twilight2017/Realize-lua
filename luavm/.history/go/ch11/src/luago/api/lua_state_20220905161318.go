package api

type LuaState interface {
	GetMetatable(idx int) bool
	SetMetatable(idx int)
	RawLen(idx int) uint
}
