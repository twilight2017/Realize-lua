package state

type luaStack struct {
	slots []luaValue
	top   int
}

//创建指定容量的栈
func newLuaStack(size int) *luaStack {
	return &luaStack{
		slots: make([]luaValue, size),
		top:   0,
	}
}

//检查栈空间是否还能容纳n个元素，若不能，则进行扩容
//扩容时用nil值进行填充
func (self *luaStack) check(n int) {
	free := len(self.slots) - self.top
	for i := free; i < n; i++ {
		self.slots = append(self.slots, nil)
	}
}

//push方法将值推入栈顶
func (self *luaStack) push(val luaValue) {
	if self.top == len(self.slots) {
		panic("stack overflow!")
	}
	self.slots[self.top] = val
	self.top++
}

//pop操作将栈顶元素出栈
func (self *luaStack) pop() (lua luaValue) {
	//stack empty
	if self.top < 1 {
		panic("stack underflow!")
	}
	lua = self.slots[self.top-1]
	self.top--
	self.slots[self.top] = nil
	return
}

//把索引转换成绝对索引
func (self *luaStack) absIndex(index int) int {
	if index >= 0 {
		return index
	} else {
		return -index
	}
}
