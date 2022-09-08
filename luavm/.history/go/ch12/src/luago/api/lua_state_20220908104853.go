package api

type LuaState interface {
	Next(idx int) bool //根据键获取表的下一键值对
}
