package ast

type Exp interface{}

//布尔,nil,vararg表达式仅需要记录行号
type NilExp struct{ Line int }
type TrueExp struct{ Line int }
type FalseExp struct{ Line int }
type VarargExp struct{ Line int }

//记录行号，并记录解析后的数值
type IntegerExp struct {
	Line int
	Val  int64
}
type FloatExp struct {
	Line int
	Val  float64
}
type StringExp struct {
	Line int
	Str  string
}

//名字表达式需要记录行号和标识符
type NameExp struct {
	Line int
	Name string
}

//定义一元运算符表达式
type UnopExp struct {
	Line int //line of operator
	Op   int //operator
	Exp  Exp
}

//定义二元运算符表达式
type BinopExp struct {
	Line int //line of operator
	Op   int //operator
	Exp1 Exp
	Exp2 Exp
}

//拼接操作
type ConcatExp struct {
	Line int
	Exps []Exp
}

type TableConstructExp struct {
	Line      int //line of '{'
	LastLine  int //line of '}'
	KeyExps   []Exp
	ValueExps []Exp
}

//函数表达式
type FuncDefExp struct {
	Line int
}

//圆括号表达式
type ParseExp struct {
	Exp Exp
}

//表访问表达式
type TableAccessExp struct {
	LastLine  int //line of ']'
	PrefixExp Exp
	KeyExp    Exp
}

//函数调用表达式
type FuncCallExp struct {
	//圆括号所在行号
	Line      int //line of `(`
	LastLine  int //line of ')'
	PrefixExp Exp
	NameExp   *StringExp
	Args      []Exp
}
