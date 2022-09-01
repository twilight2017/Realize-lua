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
