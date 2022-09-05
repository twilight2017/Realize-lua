package state

import "luago/api"

type luaStack struct {
	state *luaState
}

func newLuaStack(size int, state *luaState) *luaStack{
	return &luaStack{
		slots: make([]luaValue, size),
		top : 0,
		state: state
	}
}

func (self *luaStack) absIndex(idx int) int{
	if idx <= LUA_REGISTRYINDEX{
		return idx
	}
}
