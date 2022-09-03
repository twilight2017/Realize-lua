package api

type LuaState interface {
	NewTable()
	CreateTable(nArr, nRec int)
	GetTable(idx int) LuaType
	GetField(idx int, k string)
	GetI(idx int, i int64) LuaType
	SetTable(idx int)
	SetField(idx int, k string)
	SetI(idx int, n int64)
}
