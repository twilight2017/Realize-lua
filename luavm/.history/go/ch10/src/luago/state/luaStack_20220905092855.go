package state

type luaStack struct {
	openuvs map[int]*upvalue
}
