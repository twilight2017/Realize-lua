package state

//该函数实现对表的访问
func (self *luaState) getTable(t, k luaValue, raw bool) LuaType {
	if tb1, ok := t.(luaTable); ok {
		v := tb1.get(k)
		if raw || v != nil || tb1.hasMetafield("__index") {
			self.stack.push(v)
			return typeOf(v)
		}
	}
	if !raw {
		if mf := getMetafield(t, "__index", self); mf != nil {
			switch x := mf.(type) {
			case *luaTable:
				return self.getTable(x, k, false)
			case *closure:
				self.stack.push(mf)
				self.stack.push(t)
				self.stack.push(k)
				self.Call(2, 1)
				v := self.stack.get(-1)
				return typeOf(v)
			}
		}
	}
	panic("index error")
}
