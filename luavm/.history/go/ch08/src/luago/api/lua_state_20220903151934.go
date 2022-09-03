type LuaState interface {
	Load(chunk []byte, chunkName, mode string) int
	Call
}