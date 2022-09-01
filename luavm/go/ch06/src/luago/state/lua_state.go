package state

import "luago/binchunk"

type luaState struct {
	stack *luaStack
	//new
	proto *binchunk.Prototype
	pc    int //
}

func New(stackSize int, proto *binchunk.Prototype) *luaState {
	return &luaState{
		stack: newLuaStack(stackSize),
		proto: proto,
		pc:    0, //虚拟机一定是从第一条指令开始执行点，所以这里的pc值初始化为0
	}
}
