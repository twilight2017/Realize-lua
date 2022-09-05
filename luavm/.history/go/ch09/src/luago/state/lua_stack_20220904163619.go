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

func (self *luaStack) isValid(idx int) bool{
	if idx == LUA_REGISTRYINDEX{
		return true
	}
}

func (self *luaStack) get(idx int) luaValue{
	if idx == LUA_REGISTRY{
		return self.stack.registry
	}
}
