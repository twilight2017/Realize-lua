package api

type LuaState interface {
	Load(chunk []byte, chunkName, mode string) int //Load()方法加载二进制chunk，把函数主原型实例化为闭包并推入栈顶
	Call(nArgs, nResults int)
}
