package luago

//给值关联元表
func setMetatable(val luaValue, mt *luaTable, ls *luaState) {
	if t, ok := val.(*luaTable); ok {
		t.metatable = mt
		return
		//先判断值是否是表，如果是，直接修改其元表字段
	}
	key := fmt.Sprintf("_MT%d", typeOf(val))
	//否则根据变量类型把元表存储在注册表里
	ls.registry.put(key, mt)
}

func getMetatable(val luaValue, ls *luaState) *luaTable {
	if t, ok := val.(*luaTable); ok {
		return t.metatable
	}
	key := fmt.Sprintf("_MT%d", typeOf(val))
	if mt := ls.registry.get(key); mt != nil {
		return mt.(*luaTable)
	}
}
