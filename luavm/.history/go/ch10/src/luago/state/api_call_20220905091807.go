package state

import "luago/binchunk"

//Lua函数全部都是闭包
//Load()方法在加载闭包时，会看它是否需要Upvalue。如果需要，那么第一个Upvalue会被初始化为全局环境，其他Upvalue会被初始化成nil
func (self *luaState) Load(chunk []byte, chunkName, mode string) int {
	proto := binchunk.Undump(chunk)
	c := newLuaClosure(proto)
	self.stack.push(c)

}
