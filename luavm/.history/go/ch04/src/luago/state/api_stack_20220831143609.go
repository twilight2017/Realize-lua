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
	val := self.stack.get(fromIdx)
	self.stack.set(toIdx, val)
}

//把指定索引处的值推入栈顶
func (self *luaState) PushValue(idx int) {
	val := self.stack.get(idx)
	self.stack.push(val)
}

//将栈顶的值弹出，写入指定位置
func (self *luaState) Replace(idx int) {
	val := self.stack.get(idx)
	self.stack.push(val)
}

//将栈顶的值弹出，插入指定的位置
func (self *luaState) Insert(idx int) {
	self.Rotate(idx, 1)
}

//删除指定索引处的值，然后将该值上面的值全部下移一个位置
func (self *luaState) Remove(idx int) {
	self.Rotate(idx, -1)
	self.Pop(1)
}

//Rotate函数将栈[idxt, top]区间的值向栈顶方向旋转n个位置
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

//调转指定序列的元素位置
func (self *luaStack) reverse(from, to int) {
	slots := self.slots
	for from < to {
		slots[from], slots[to] = slots[to], slots[from]
		from++
		to--
	}
}

//将栈顶索引设为指定值，如果指定值小于当前栈顶索引，则相当于执行弹出操作
//否则则相当于进行扩容操作
func (self *luaState) SetTop(idx int) {
	newTop := self.stack.absIndex(idx) //首先将待设置索引转换为指定值
	if newTop < 0 {
		panic("stack overflow")
	}
	n := self.stack.top - newTop //计算需弹出元素的个数
	if n > 0 {
		for i := 0; i < n; i++ {
			self.stack.pop()
		}
	} else if n < 0 {
		for i := 0; i < n; i++ {
			self.stack.push(nil)
		}
	}
}
