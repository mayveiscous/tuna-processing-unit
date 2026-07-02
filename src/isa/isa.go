package isa

const (
	OP_LOAD  = 0x01
	OP_ADD   = 0x02
	OP_SUB   = 0x03
	OP_MUL   = 0x04
	OP_DIV   = 0x05
	OP_MOV   = 0x06
	OP_CMP   = 0x07
	OP_JMP   = 0x08
	OP_JZ    = 0x09
	OP_JNZ   = 0x0A
	OP_PUSH  = 0x0B
	OP_POP   = 0x0C
	OP_PRINT = 0x0D
	OP_CALL  = 0x0E
	OP_RET   = 0x0F
	OP_LOADM = 0x10
	OP_STORE = 0x11
	OP_INC = 0x12
	OP_DEC = 0x13
	OP_JG = 0x14
	OP_JL = 0x15
	OP_HALT  = 0xBE
)

const (
	StackStart = 0xFF
	StackLimit = 0xC0
)

var Commands []string


func PackRegs(a, b uint8) uint8 {
	return (a << 4) | b
}

func UnpackRegs(v uint8) (uint8, uint8) {
	return v >> 4, v & 0x0F
}

func RegisterCommands(commands []string) {
	Commands = append(Commands, commands...)
}