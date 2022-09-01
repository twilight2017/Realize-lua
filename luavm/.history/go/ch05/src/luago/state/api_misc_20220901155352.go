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
