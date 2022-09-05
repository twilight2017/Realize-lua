func (self *luaState) PCall(nArgs, nResults, msgh int) (status int){
	caller := self.stack
	status = LUA_ERRRUN

	//catch error
	defer func(){

	}

	self.Call(nArgs, nResults)
	status = LUA_OK
	return
}