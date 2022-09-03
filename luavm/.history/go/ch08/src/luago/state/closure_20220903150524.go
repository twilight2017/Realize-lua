package state

import "luaogo/binchunk"

type closure struct {
	proto *binchunk.Prototype //存放函数原型
}

//用于创建lua闭包的函数
func newLuaClosure(proto *binchunk.Prototype) *closure {
	return &closure{proto: proto}
}
