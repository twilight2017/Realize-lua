package state

import "luago/binchunk"

type closure struct {
	proto  *binchunk.Prototype
	goFunc GoFunction
	upvals []*upvalue
}

type upvalue struct {
	val *luaState
}
