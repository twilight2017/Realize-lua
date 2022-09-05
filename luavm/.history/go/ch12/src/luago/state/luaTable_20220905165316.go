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
	return self.keys[key]
}

func (self *luaTable) initKeys(){
	self.keys := make(map[luaValue]luaValue)
	var key luaValue = nil
	for i, v := range self.arr{
		if v != nil{
			self.keys[key] = int64(i +1)
			key = int64(i + 1)
		}
	}
	for k, v := range self._map{
		if v !=nil{
			self.keys[key] = k
			key = k
		}
	}
}


//表的索引由参数指定
//上一个键从栈顶弹出
//如果栈顶弹出的是nil，说明刚开始遍历表，把表的第一个键值对推入栈顶并返回true