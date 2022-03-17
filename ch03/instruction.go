package ch03

import "fmt"

type Instruction uint32

const (
	MAXARG_Bx  = 1<<18 - 1
	MAXARG_sBx = MAXARG_Bx >> 1
)

//从指令中提取操作码
func (self Instruction) Opcode() int {
	return int(self & 0x3F)
}

//从iABC指令中提取参数
func (self Instruction) ABC() (a, b, c int) {
	a = int(self >> 6 & 0xFF)
	c = int(self >> 14 & 0x1FF)
	b = int(self >> 23 & 0x1FF)
	return
}

//从iABx指令中提取参数
func (self Instruction) ABx() (a, bx int) {
	a = int(self >> 6 & 0xFF)
	bx = int(self >> 14)
}

//AsBx()从iAsBx指令中提取参数
func (self Instruction) AsBx() (a, sbx int) {
	a, bx := self.ABx
	return a, bx - MAXARG_sBx
}

//从iAx模式指令中提取参数
func (self Instruction) Ax() int {
	return int(self >> 6)
}

//返回指令的操作码名字
func (self Instruction) OpName() string {
	return opcodes[self.Opcode()].name
}

//返回编码格式
func (self Instruction) OpMode() byte {
	return opcodes[self.Opcode()].OpMode
}

//返回操作数B的使用模式
func (self Instruction) BMode() byte {
	return opcodes[self.Opcode()].BMode
}

//返回操作数C的使用模式
func (self Instruction) CMode() byte {
	return opcodes[self.Opcode()].CMode
}

//打印指令的操作数
func printOperands(i Instruction) {
	switch i.OpMode() {
	case IABC:
		a, b, c := i.ABC()
		fmt.Printf("%d", a)
		if i.BMode() != OpArgN {
			if b > 0xFF {
				fmt.Printf("%d", -1-b&0xFF)
			} else {
				fmt.Printf("%d", b)
			}
		}
		if i.CMode != OpArgN {
			if c > 0xFF {
				fmt.Printf("%d", -1-c&0xFF)
			} else {
				fmt.Printf("%d", c)
			}
		}
	case IABx:
		a, bx := i.ABx()
		fmt.Printf("%d", a)
		if i.BMode() == OpArgK {
			fmt.Printf("%d", -1-bx)
		} else if i.BMode() == OpArgU {
			fmt.Printf("%d", bx)
		}
	case IAsBx:
		a, sbx := i.AsBx()
		fmt.Printf("%d %d", a, sbx)
	case IAx:
		ax := i.Ax()
		fmt.Printf("%d", -1-ax)
	}

}
