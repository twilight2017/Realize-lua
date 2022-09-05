package state

//只有Go闭包能转换为Go函数
func (self *luaState) IsGoFunction(idx int) bool {
	val := self.stack.get(idx)
	if c, ok := val.(*closure); ok {
		return c.goFunc != nil
	}
	return false
}

//把索引处的值转换为Go函数并返回
