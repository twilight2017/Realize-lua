package state

import "luaogo/binchunk"

type closure struct {
	proto *binchunk.Prototype //存放函数原型
}
