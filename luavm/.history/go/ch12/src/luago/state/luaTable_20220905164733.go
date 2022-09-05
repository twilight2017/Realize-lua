package state

type luaTable struct {
	keys map[luaValue]luaValue
}

//Go语言的map不能保证顺序遍历
func (self *luaTable) nextKey(key luaValue) luaValue {
	if self.keys == nil || key == nil {
		self.initKeys()
		self.changed = false
	}
}

func (self *luaTable) initKeys(){
	self.keys := make(map[luaValue]luaValue)
	var key luaValue = nil
	for i, v := range self.arr{
		if v != nil{
			self.keys[key] = int64(i +1)
			key = int64(i +1)
		}
	} 
}
