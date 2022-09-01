package state

import "fmt"
import . "luago/api"

//把给定的lua类型转换为对应的字符串表示
func (self *luaState) TypeName(tp LuaType) string {
    switch tp{
		case LUA_TNONE:return "no value",
	    case LUA_TNIL: return "nil",
	    case LUA_RBOOLEAN: return "boolean",
	}
}
