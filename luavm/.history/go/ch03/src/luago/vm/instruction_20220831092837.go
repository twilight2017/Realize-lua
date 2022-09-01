package vm

type Instruction uint32

//从指令中取出操作码
func (self Instruction) Opcode() int {
	return int(self & 0x3F)
}
