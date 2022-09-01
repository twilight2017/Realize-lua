package state

import "fmt"
import . "luago/api"

//把给定的lua类型转换为对应的字符串表示
func (self *luaState) TypeName(tp LuaType) string {
    switch tp{
		case LUA_TNONE:return "no value"
	    case LUA_TNIL: return "nil"
		case LUA_RBOOLEAN: return "boolean"
	case LUA_TNUMBER:return "number"
	case LUA_TSTRING: return "string"
	case LUA_TTABLE: return "table"
	case LUA_TFUNCTION return "function"
	case LUA_TTHREAD return "thread"
	default: "userdata"
	}
}
