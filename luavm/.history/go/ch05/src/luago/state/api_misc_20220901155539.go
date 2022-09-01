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
	}else if n >= 2{
		for i := 0;i<n;i++
	}
}
