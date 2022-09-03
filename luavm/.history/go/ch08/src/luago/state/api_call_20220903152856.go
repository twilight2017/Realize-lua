package state

import "fmt"
import "luago/binchunk"
import "luago/vm"

func (self *luaState) Load(chunk []byte, chunkName, mode string) int {
	proto := binchunk.Undump(chunk)
	c := newLuaClosure(proto)
	self.stack.push(c)
}
