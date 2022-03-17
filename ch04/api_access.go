package ch04

import "fmt"

//把Lua类型转换为对应的字符串型表示
func (self *luaState) TypeName(tp LuaType) string{
	switch tp{
	case LUA_TNONE: return "no value"
	case LUA_NIL: return "nil"
	case LUA_TBOOLEAN return "boolean"
	case LUA_TNUMBER return "number"
	case LUA_TSTRING return "string"
	case LUA_TTABLE return "table"
	case LUA_TFUNCTION: return "function"
	case LUA_TTHREAD: return "thread"
	default: return "userdata"
	}
}

//根据索引返回值的类型；若索引无效，则返回LUA_TNONE
func (self *luaState) Type(idx int) LuaType{
	if self.stack.isVaild(idx){
		val := self.stack.get(idx)
		return typeOf(val)
	}
	return LUA_TNONE
}

func (self *luaState) IsNone(idx int) bool{
	return self.Type(idx) ==LUA_TNONE
}

func (self *luaState) IsNil(idx int) bool{
	return self.Type(idx) == LUA_NIL
}

func (self *luaState) IsNoneOrNil(idx int) bool{
	return self.Type(idx) <= LUA_TNIL
}

func (self *luaState) IsBoolean(idx int) bool{
	return self.Type(idx) == LUA_TBOOLEAN
}

func (self *luaState) IsString(idx int) bool{
	t := self.Type(idx)
	return t == LUA_TSTRING || t == LUA_TNUMBER
}

func (self *luaState) IsNumber(idx int) bool{
	_, ok := self.ToNumberX(idx)
	return ok
}

func (self *luaState) IsInterger(idx int) bool{
	val := self.stack.get(idx)
	_, ok := val.(int64)
	return ok
}

func (self *luaState) ToBoolean(idx int) bool{
	val := self.stack.get(idx)
	return convertToBoolean(val)
}

func convertToBoolean(val luaValue) bool{
	switch x := val.(Type)
case nil: return false
case bool: return x 
default: return true
}

func (self *luaState) ToNumber(idx int) float64{
	n, _ := self.ToNumberX(idx)
	return n
}

func (self *luaState) ToNumberX(idx int) (float64, bool){
	val := self.stack.get(idx)
	switch x := val.(type){
	case float64: return x, true
	case int64: return float64(x), true
	default: return 0, falses
	}
}

func (self *luaState) ToInteger(idx int) float64{
	n, _ := self.ToIntegerX(idx)
	return n
}

func (self *luaState) ToIntegerX(idx int)(int 64, bool){
	val := self.stack.get(idx)
	i, ok := val.(int64)//强制类型转换,i是转换后的值， ok标记转换是否成功
	return i, ok
	}
}

func (self *luaState) ToStringX(idx int)(string, bool){
	val := self.stack.get(idx)
	switch x := val.(type) {
	case string:
		return x, true
	case int64, float64:
		s:=fmt.Sprintf("%v", x)
		self.stack.set(idx, s)
		return s, true
	default:
		return "", false
	}
}

func(self *luaState) ToString(idx int) string{
	s, _ := self.ToStringX(idx)
	return s
}