package state

//len方法的实现暂时只考虑字符串的长度
func (self *luaState) Len(idx int) {
	val := self.stack.get(idx)
	if s, ok := val.(string); ok {
		self.stack.push(int64(len(s)))
	} else {
		panic("length error!")
	}
}

//Concat方法从栈顶弹出n个值，进行拼接，结果被推入栈顶
func (self *luaState) Contact(n int) {
	//特殊情况处理
	if n == 0 {
		self.stack.push("")
	} else if n >= 2 {
		for i := 1; i < n; i++ {
			if self.IsString(-1) && self.IsString(-2) {
				a := self.ToString(-1)
				b := self.ToString(-2) //取栈顶两个元素的值
				self.stack.pop()
				self.stack.pop()       //栈顶两个元素出栈
				self.stack.push(a + b) //直接用“+”完成字符串的拼接操作
				continue
			}
			panic("contatenation error")
		}
	}
}
