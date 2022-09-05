package state

func _eq(a, b luaValue, ls *luaState) bool {
	switch x := a.(type) {
	case *luaTable:
		if y, ok := b.(*luaTable); ok && x != y && ls != nil {
			if result, ok := callMetamehod(x, y, "__eq", ls); ok {
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
}
