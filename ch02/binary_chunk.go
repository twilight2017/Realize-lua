package binchunk

type binaryChunk struct {
	header                 //头部
	sizeUpvalue byte       //主函数upvalue数量
	mainFunc    *Prototype //主函数原型
}

//定义头部
type header struct{
	signature [4]byte //签名，用于识别这是一个lua二进制chunk
	version byte//版本号：记录该二进制chunk对应的lua版本号，其值由大小版本号计算得到
	format byte//格式:lua虚拟机在加载二进制chunk时，会检查其格式号，如不匹配，则拒绝加载该文件
	luacData  [6]byte//起进一步校验的作用
	/*以下5个字段记录这5种数据类型在二进制chunk种占用的字节数*/
	cintSize  byte
	sizeSize  byte
	instructionSize  byte
	luaIntegerSize  byte
	luaNumberSize  byte
	luacInt  int64 //该整数用于校验二进制chunk的大小端方式；不匹配时，拒绝加载
	luacNum  float64//用于检测二进制chunk所使用的浮点数格式
}

//定义常量
const(
	LUA_SIGNATURE  ="\x1bLua"
	LUAC_VERSION  =0x53
	LUAC_FORMAT  =0
	LUAC_DATA  ="\x19\x93\r\n\xla\n"
	CINT_SIZE  =4
	CSZIET_SIZE  =8
	INSTRUCTION_SIZE  =4
	LUA_INTEGER_SIZE  =8
	LUA_NUMBER_SIZE  =8
	LUAC_INT  =0x5678
	LUAC_NUM  =370.5
)

//定义函数原型
type Prototype struct{
	Source  string//源文件名：记录二进制chunk是由哪个源文件编译得到的
	LineDefined uint32
	LastLineDefined  uint32 //起止行号：记录原型对应的函数在源文件中的起止行号
	NumParams  byte //固定参数个数
	IsVararg  byte  //标记是否是Vararg函数
	MaxStackSize  byte//标记寄存器数量
	Code  []uint32//指令表
	Constants  []interface{}
	Upvalues  []Upvalue
	Ptotos  []*Prototype//子函数原型表
	LineInfo  []uint32//行号表，分别记录每条指令在源代码中的行号
	LocVars  []LocVar//局部变量表：用于记录局部变量名
	UpvalueNames  []string//Upvalue名列表：记录Upvalue在源代码中的名字

}

//定义常量表
const(
	TAG_NIL  =0x00
	TAG_BOOLEAN  =0x01
	TAG_NUMBER  =0x03
	TAG_INTEGER  =0x13
	TAG_SHORT_STR  =0x04
	TAG_LONG_STR  =0x14
)

//定义Upvalue表
type Upvalue struct{
	Instack byte
	Idx byte
}

//定义局部变量表
type LocVars{
	VarName string
	StartPC uint32
	EndPC uint32
}

//用于解析二进制chunk
func Undump(data []byte) *Prototype{
	reader := &reader(data)
	reader.checkHeader()//校验头部
	read.readByte()  //跳过Upvalue数量
	return reader.readProto("")//读取函数原型
}