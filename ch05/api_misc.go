package ch05

func (self *luaState) Len(idx int) {
	val := self.stack.get(idx)
	if s, ok := val.(string); ok {
		self.stack.push(int64(len(s)))
	} else {
		panic("length error!")
	}
}

//从栈顶弹出n个值，进行拼接后压入栈中
func (self *luaState) Contact(n int) {
	if n == 0 {
		self.stack.push("") //n为0，则直接推入一个空字符串
	} else if n >= 2 {
		for i := 1; i < n; i++ {
			if self.IsString(-1) && self.IsString(-2) {
				s2 := self.ToString(-1)
				s1 := self.ToString(-2) //取栈顶的两个值
				self.stack.pop()        //弹出，空出物理位置
				self.stack.pop()
				self.stack.push(s1 + s2)
				continue
			}
			panic("concatenation error!")
		}
	}
}
