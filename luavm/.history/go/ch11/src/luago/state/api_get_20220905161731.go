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

//看指定索引处是否有元表，如果有，则把元表推入栈顶并返回true
//否则栈的状态不变，返回false
func (self *luaState) GetMetatable(idx int) bool {
	val := self.stack.get(idx)

	if mt := GetMetatable(val, self); mt != nil {
		self.stack.push(mt)
		return true
	} else {
		return false
	}
}
