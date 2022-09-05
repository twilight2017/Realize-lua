package state

func (self *luaState) PushGoFunction(f GoFunction) {
	self.stack.push(newGoClosure(f))
}

//判断索引处的值是否能转换为Go函数
//该方法以栈索引为参数，返回布尔值，不改变栈的状态
func (self *luaState) PushGlobalTable() {
	global := self.registry.get(LUA_RIDX_GLOBALS)
	self.stack.push(global)
}
