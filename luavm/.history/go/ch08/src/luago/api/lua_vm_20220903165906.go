type LuaVM interface {
	RegisterCount() int
	LoadVararg(n int)
	LoadProto(idx int)
}