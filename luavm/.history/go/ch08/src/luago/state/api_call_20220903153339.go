package state

import "fmt"
import "luago/binchunk"
import "luago/vm"

func (self *luaState) Load(chunk []byte, chunkName, mode string) int {
	proto := binchunk.Undump(chunk)
	c := newLuaClosure(proto)
	self.stack.push(c)
	return 0
}

//Call()方法对Lua函数进行调用
/*
1.把被调函数推入栈顶
2.把参数值依次推入栈顶
3.Call()方法结束后，参数值和函数会被弹出栈顶
*/
