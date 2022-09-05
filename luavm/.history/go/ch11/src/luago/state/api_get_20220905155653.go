package state

//该函数实现对表的访问
func (self *luaState) getTable(t, k luaValue, raw bool) LuaType {
	if tb1, ok := t.(luaTable); ok {
		v := tb1.get(k)
		if raw || v != nil || tb1.hasMetafield("__index") {
			self.stack.push(v)
		}
	}
}
