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

func newLuaClosure(proto *binchunk.Prototype) *closure {
	c := &closure{proto: proto}
	if nUpvals := len(proto.Upvalues); nUpvals > 0 {
		c.upvals = make([]*upvalue, nUpvals) //Lua闭包捕获的Upvalue数量已经由编译器计算好了，在这里先预分配空间
	}
}
