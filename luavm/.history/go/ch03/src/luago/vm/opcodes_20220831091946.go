package vm

const (
	IABC = iota
	IABx
	IAsBx
	IAx
)

//定义操作码常量
const (
	OP_MOVE = iota
	OP_LOADK
	OP_LOADKX
	OP_LOADBOOL
	OP_LOADNIL
	OP_GETUPVAL
	OP_GETTABUP
	OP_GETTABLE
	OP_SETTABUP
	OP_SETUPVAL
	OP_SETTABLE
	OP_NEWTABLE
	OP_SELF
	OP_ADD
	OP_SUB
	OP_MUL
	OP_MOD
	OP_POW
	OP_DIV
	OP_IDIV
	OP_BAND
	OP_BOR
	OP_BOXR
	OP_SHL
	OP_SHR
	OP_UNM
	OP_BNOT
	OP_NOT
	OP_LEN
	OP_CONCAT
	OP_JMP
	OP_EQ
	OP_LT
	OP_LE
	OP_TEST
	OP_TESTSET
	OP_CALL
	OP_TALLCALL
	OP_RETURN
	OP_FORLOOP
	OP_FORPREP
	OP_TFORCALL
	OP_TFORLOOP
	OP_SETLIST
	OP_CLOSURE
	OP_VARAGE
	OP_EXTRAARG
)

const (
	OpArgN = iota //argument is not used
	OpARgU        //argument is used
	OpArgR        //argument is a register or a jump offset
	OpArgK        //argument is a constant or register/constant
)

type opcode struct {
	testFlag byte //operator is a test (next instruction must be a jump)
	setAFlag byte //instruction set register A
	argBMode byte // B arg mode
	argCMode byte // C arg mode
	opMode   byte //op mode
	name     string
}