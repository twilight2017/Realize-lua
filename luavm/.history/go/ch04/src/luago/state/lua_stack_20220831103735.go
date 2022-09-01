package state

type luaStack struct {
	slots []luaValue
	top   int
}

//创建指定容量的栈
func newLuaStack(size int) *luaStack {
	return &luaStack{
		slots: make([]luaValue, size),
		top:   0,
	}
}
