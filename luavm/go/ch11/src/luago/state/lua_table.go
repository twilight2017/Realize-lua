package state

//把每种类型和元表关联起来
type luaTable struct {
	metatable *luaTable
	arr       []luaValue
	_map      map[luaValue]luaValue
}

func (self *luaTable) hasMetafield(fieldName string) bool {
	return self.metatable != nil &&
		self.metatable.get(fieldName) != nil
}
