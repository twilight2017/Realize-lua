package state

//从栈顶将错误对象弹出
//调用Go语言内置的panic()函数将它抛出即可
func (self *luaState) Error() int {
	err := self.stack.pop()
	panic(err)
}
