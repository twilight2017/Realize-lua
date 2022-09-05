package api

type LuaVM interface {
	CloseUpvalues(a int)
}

func (self *luaState) CloseUpvalues(a int)
