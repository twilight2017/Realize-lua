package api

type LuaVM interface {
	CloseUpvalues(a int)
}

func (self *luaState) CloseUpvalues(a int){
	for i. openuv := range self.stack.openuvs{
		if i >= a-1{
			val := *openuv.val
			openuv.val = &val
			delete(self.stack.openuvs, i)
		}
	}
}
