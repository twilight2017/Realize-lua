package state

//1.首先判断值是否为字符串，若是，则返回字符串长度
//2.看值是否有_len方法，如果有，以值为参数调用元方法。将元方法返回值作为结果
//3.若找不到对应元方法，但值是表，结果就是表的长度
func (self *luaState) Len(idx int) {
	val := self.stack.get(idx) //首先根据索引获得值

	if str, ok := val.(string); ok {
		self.stack.push(int64(len(str)))
	} else if result, ok := callMetamethod(val, val, "_len", self); ok {
		self.stack.push(result)
	} else if t, ok := val.(*luaTable); ok {
		self.stack.push(int64(t.len()))
	} else {
		panic("length error!")
	}
}

//拼接方法
//1.如果两个操作数都是字符串，则进行字符串拼接
//2.否则，尝试调用_concat方法，查找和调用规则同二元算术元方法
func (self *luaState) Concat(n int) {
	if n < 1 {
		panic("paramter error")
	} else if n >= 2 {
		for i := 1; i < n; i++ {
			b := self.stack.pop()
			a := self.stack.pop()
			if result, ok := callMetamethod(a, b, "__concat", self); ok {
				self.stack.push(result)
				continue
			}
			panic("contatenation error!")
		}
	}
}
