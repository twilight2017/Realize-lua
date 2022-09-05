package state

import "luago/binchunk"

type closure struct {
	proto  *binchunk.Prototype
	goFunc GoFunction
	upvals []*upvalue
}

//添加一个间接层，来封装upvalue
type upvalue struct {
	val *luaState
}
