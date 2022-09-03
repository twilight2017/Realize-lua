package state

import . "luago/api"

type luaState struct {
	stack *luaStack
	//new
	proto *binchunk.Prototype
	pc    int //
}

//创建一个空的Lua表，将其推入栈顶
func (self *luaState) CreateTable(nArr, nRec int) {
	t := newLuaTable(nArr, nRec)
	self.stack.push(t)
}

//无法预估表的用法容量时，先创建一个表
func (self *luaState) NewTable(nArr, nRec int) {
	self.CreateTable(0, 0)
}

//从表中取值，把值推入栈顶并返回值的类型
func (self *luaState) GetTable(idx int) luaType {
	val := self.stack.get(idx)
	k := self.stack.pop() //键
	return self.getTable(t, k)
}

//根据键从表里取值的方法
func (self *luaState) getTable(t, k luaValue) LuaType {
	if tb1, ok := t.(*luaTable); ok {
		v := tb1.get(k)
		self.stack.push(v)
		return typeOf(v)
	}
	panic("not a table")
}

func (self *luaState) GetField(idx int, k string) LuaType {
	t := self.stack.get(idx)
	return self.getTable(t, k)
}
