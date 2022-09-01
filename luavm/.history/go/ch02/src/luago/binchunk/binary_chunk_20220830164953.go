package binchunk

type binaryChunk struct {
	header                 //头部
	sizeUpvalue byte       //主函数upvalue数量
	mainFunc    *Prototype //主函数原型
}
