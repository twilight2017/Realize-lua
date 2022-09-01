package state

import "luago/binchunk"

type luaState struct {
	stack *luaStack
	//new
	proto *binchunk.Prototype
	pc    int
}
