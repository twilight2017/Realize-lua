package state

//该函数实现对表的访问
func (self *luaState) getTable(t, k luaValue, raw bool) LuaType
