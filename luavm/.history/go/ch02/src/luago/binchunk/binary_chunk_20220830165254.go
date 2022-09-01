package binchunk

type binaryChunk struct {
	header                 //头部
	sizeUpvalue byte       //主函数upvalue数量
	mainFunc    *Prototype //主函数原型
}

type header struct {
	signature       [4]byte
	version         byte
	format          byte
	luacData        [6]byte
	cintSize        byte
	sizeSize        byte
	instructionSize byte
	luaIntegerSize  size
	luaNumberSize   byte
	luaCInt         int64
	luacNum         float64
}
