package state

//把每种类型和元表关联起来
type luaTable struct {
	metatable *luaTable
	arr       []luaValue
	_map      map[luaValue]luaValue
}
