package api

type LuaType = int
type ArithOp = int   //新增的类型别名
type CompareOp = int //新增的类型别名

type LuaState interface {
	Arith(op ArithOp)                          //用于执行算术和按位运算
	Compare(idx1, idx2 int, op CompareOp) bool //用于执行比较运算
	Len(idx int)                               //用于执行取长度运算
	Concat(n int)                              //用于执行字符串拼接运算
}
