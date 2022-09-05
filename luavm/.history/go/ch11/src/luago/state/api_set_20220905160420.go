package state

func (self *luaState) setTable(t, k, v luaValue, raw bool) {
	if tb1, ok := t.(*luaTable); ok {
		if raw || tb1.get(k) != nil || !tb1.hasMetafield("__newindex") {
			tb1.put(k, v)
			return
		}
	}
	if !raw {

	}
}
