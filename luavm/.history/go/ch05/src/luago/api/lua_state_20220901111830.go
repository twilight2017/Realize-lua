package api

type LuaType = int
type ArithOp = int
type CompareOp = int

type LuaState interface {
	Arith(op ArithOp)
	Compare(idx1, idx2 int, op CompareOp) bool
	Len(idx int)
	Concat(n int)
}
