package state

func (self *luaState) setTable(t, k, v luaValue, raw bool) {
	if tb1, ok := t.(*luaTable); ok {
		if raw || tb1.get(k) != nil || !tb1.hasMetafield("__newindex") {
			tb1.put(k, v)
			return
		}
	}
	if !raw {
		if mf := getMetafield(t, "__newindex", self); mf != nil {
			switch x := mf.(type) {
			case *luaTable:
				self.setTable(x, k, v, false)
				return
			case *closure:
				self.stack.push(mf)
				self.stack.push(t)
				self.stack.push(k)
				self.stack.push(v)
				self.Call(3, 0)
				return
			}
		}
	}
	panic("index error")
}

//从栈顶弹出一个表，然后把索引处值的元表设置成该表
func (self *luaState) SetMetatable(idx int) {
	val := self.stack.get(idx)
	mtVal := self.stack.pop()

	if mtVal == nil {
		SetMetatable(val, nil, self)
	} else if mt, ok := mtVal.(*luaTable); ok {
		SetMetatable(val, mt, self)
	} else {
		panic("table expected!")
	}
}
