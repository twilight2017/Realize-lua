package api

type LuaState interface {
	Error() int
	PCall(nArgs, nResults, msgh int) int
}
