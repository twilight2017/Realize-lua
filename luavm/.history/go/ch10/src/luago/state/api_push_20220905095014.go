package state

func (self *luaState) PushGoFunction(f GoFunction) {
	self.stack.push(newGoClosure(f, 0))
}

func (self *luaState) PushGoClosure(f GoFunction, n int) {
	closure := newGoClosure(f, n) //先创建Go闭包
	for i := n; i > 0; i-- {      //从栈顶弹出指定数量的值，让它们变成闭包的Upvalue
		val := self.stack.pop()
		closure.upvals[n-1] = &upvalue{&val}
	}
	self.stack.push(closure)
}
