package state

type luaState struct {
	stack *luaStack
}

func New() *luaState{
	return &luaState{
		stack: newLuaStack(20)
	}
}

func (self *luaState) pushLuaStack(stack *luaStack){
	stack.prev = self.stack
	self.stack = stack
}