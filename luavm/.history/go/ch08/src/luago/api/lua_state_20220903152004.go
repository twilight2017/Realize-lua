type LuaState interface {
	Load(chunk []byte, chunkName, mode string) int //Load()方法加载二进制chunk
	Call(nArgs, nResults int)
}