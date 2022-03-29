package ch06

type luaState struct {
	stack *luaStack
	proto *binchunk.Ptototype
	pc    int
}
