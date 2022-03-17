//api_stack.go
package ch04

//返回栈顶索引
func (self *luaState) GetTop() int {
	return self.stack.top
}

//把索引转换为绝对索引
func (self *luaState) AbsIndex(idx int) int {
	return self.stack.absIndex(idx)
}

//检查栈的剩余空间
/*剩余空间足够或扩容成功返回True，否则返回False*/
func (self *luaState) CheckStack(n int) bool {
	self.stack.check(n)
	return true
}

//pop
func (self *luaState) Pop(n int) {
	for i := 0; i < n; i++ {
		self.stack.pop()
	}
}

//copy把值从一处复制到另一处
func (self *luaState) Copy(fromIdx, toIdx int) {
	val := self.stack.get(fromIdx)
	self.stack.set(toIdx, val)
}

//把指定索引处的值推入栈顶
func (self *luaState) PushValue(idx int) {
	val := self.stack.get(idx)
	self.stack.push(val)
}

//将栈顶值弹出，写入指定位置
func (self *luaState) Replace(idx int) {
	val := self.stack.pop()
	self.stack.set(idx, val)
}

//插入
func (self *luaState) Insert(idx int) {
	self.Rotate(idx, 1)
}

//删除指定索引处的值
func (self *luaState) Remove(idx int) {
	self.Rotate(idx, -1)
	self.Pop(1)
}

//将指定位置处的值进行位置旋转
func (self *luaState) Rotate(idx, n int) {
	t := self.stack.top - 1
	p := self.stack.absIndex(idx) - 1
	var m int
	if n >= 0 {
		m = t - n
	} else {
		m = p - n - 1
	}
	self.stack.reverse(p, m)
	self.stack.reverse(m+1, t)
	self.stack.reverse(p, t)
}
