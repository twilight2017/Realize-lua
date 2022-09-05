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
//若值无法转换为Go函数，则返回nil即可
func (self *luaState) ToGoFunction(idx int) GoFunction {
	val := self.stack.get(idx)
	if c, ok := val.(*closure); ok {
		//先根据索引拿到值，看它是否是闭包。
		//如果是闭包，直接返回goFunc字段即可
		return c.goFunc
	}
	return nil
}
