package api

// const (
// 	LUA_TNONE = iota - 1 //-1
// 	LUA_TNIL
// 	LUA_TBOOLEAN
// 	LUA_TLIGHTUSERDATA
// 	LUA_TNUMBER
// 	LUA_TSTRING
// 	LUA_TTABLE
// 	LUA_TFUNCTION
// 	LUA_TTHREAD
// )
type luaType struct{
	LUA_TNONE = iota - 1 //-1
	LUA_TNIL
	LUA_TBOOLEAN
	LUA_TLIGHTUSERDATA
	LUA_TNUMBER
	LUA_TSTRING
	LUA_TTABLE
	LUA_TFUNCTION
	LUA_TTHREAD
}