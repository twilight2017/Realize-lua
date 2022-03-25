//从索引处取出一个数字，若不是数字，则需要进行类型转化
func (self *luaState) ToNumberX(idx int) (float64, bool) {
	val := self.stack.get(idx)
	return convertToFloat(val)
}

//从索引出取出一个整数，若不是整数，则需要进行类型转换
func (self *luaState) ToIntegerX(idx int) (int64, bool) {
	val := self.stack.get(idx)
	return convertToInteger(val)
}