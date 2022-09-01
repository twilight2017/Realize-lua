package state

type luaState struct {
	stack *luaStack
}

func New() *luaState {
	return &luaState{
		stack: newLuaStack(20),
	}
}

//返回栈顶索引
func (self *luaState) GetTop() int {
	return self.stack.top
}

//把索引转换为绝对索引
func (self *luaState) AbxIndex(idx int) int {
	return self.stack.absIndex(idx)
}

//检查栈容量并扩容
func (self *luaState) CheckStack(n int) bool {
	self.stack.check(n) //直接调用已经封装好的方法即可
	return true
}

//pop()方法从栈顶弹出n个值
func (self *luaState) Pop(n int) {
	for i := 0; i < n; i++ {
		self.stack.pop()
	}
}

//copy方法把值从一个位置复制到另一个位置
func (self *luaState) Copy(fromIdx, toIdx int) {
	val := self.stack.get()
}
