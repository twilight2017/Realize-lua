package state

//1.首先判断值是否为字符串，若是，则返回字符串长度
//2.看值是否有_len方法，如果有，以值为参数调用元方法。将元方法返回值作为结果
//3.若找不到对应元方法，但值是表，结果就是表的长度
func (self *luaState) Len(idx int) {
	val := self.stack.get(idx) //首先根据索引获得值

	if str, ok := val.(string); ok {
		self.stack.push(int64(len(str)))
	}
}
