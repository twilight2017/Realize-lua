package api

type LuaType = int
type ArithOp = int
type CompareOp = int

type LuaState interface {
	Arith(op ArithOp)                          //执行算术和按位运算
	Compare(idx1, idx2 int, op CompareOp) bool //执行比较操作
	Len(idx int)                               //执行取长度运算
	Concat(n int)                              //执行字符串拼接运算
}
