package state

func _eq(a, b luaValue, ls *luaState) bool {
	switch x := a.(type) {
	case *luaTable:
		if y, ok := b.(*luaTable); ok && x != y && ls != nil {
			if result, ok := callMetamethod(x, y, "__eq", ls); ok {
				return convertToBoolean(result)
			}
		}
		return a == b
	default:
		return a == b
	}
}

func _lt(a, b luaValue, ls *luaState) bool{
	switch x := a.(type)
	if result, ok:= callMetamethod(a, b, "_lt", ls);ok{
		return convertToBoolean(result)
	}else{
		panic("comparison error!")
	}
}

func _le(a, b luaValue, ls *luaState) bool{
	switch x :=a.(type)
	if result, ok := callMetamethod(a, b, "_le", ls);ok{
		return convertToBoolean(result)
	}else if result, ok := callMetamethod(b, a, "__lt", ls);ok{
		return !convertToBoolean(result)
	}else{
		panic("comparison error!")
	}
}

func (self *luaState) Compare(idx1, idx2 int, op CompareOp) bool{
	switch op{
	case LUA_OPEQ: return _eq(a, b, self)
	case LUA_OPLT: return _lt(a, b, self)
	case LUA_OPLE: return _le(a, b, self)
	default: panic("invalid compare op!")
	}
}