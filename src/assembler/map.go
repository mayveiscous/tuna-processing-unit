package assembler

import (
	"fmt"
	"strconv"
	"strings"

	"tpu/src/isa"
)

type InstructionHandler func(args []string, labels map[string]uint8) []uint8

func parseReg(reg string) uint8 {
	if !strings.HasPrefix(reg, "R") {
		panic(fmt.Sprintf("invalid register '%s'", reg))
	}

	val, err := strconv.ParseUint(reg[1:], 10, 8)
	if err != nil {
		panic(fmt.Sprintf("invalid register '%s'", reg))
	}

	return uint8(val)
}

func parseNum(num string) uint8 {
	val, err := strconv.ParseUint(num, 0, 8)
	if err != nil {
		panic(fmt.Sprintf("invalid number '%s'", num))
	}
	return uint8(val)
}

func resolveAddress(arg string, labels map[string]uint8) uint8 {
	if addr, ok := labels[arg]; ok {
		return addr
	}

	return parseNum(arg)
}

func resolveReg(args []string) (uint8, uint8) {
	regA := parseReg(args[0])
	regB := parseReg(args[1])
	return regA, regB
}

func LOAD(args []string, labels map[string]uint8) []uint8 {
	return []uint8{isa.OP_LOAD, parseReg(args[0]), parseNum(args[1])}
}

func LOADM(args []string, labels map[string]uint8) []uint8 {
	return []uint8{isa.OP_LOADM, parseReg(args[0]), parseNum(args[1])}
}

func STORE(args []string, labels map[string]uint8) []uint8 {
	return []uint8{isa.OP_STORE, parseReg(args[0]), parseNum(args[1])}
}

func MOV(args []string, labels map[string]uint8) []uint8 {
	regA, regB := resolveReg(args)
	return []uint8{isa.OP_MOV, isa.PackRegs(regA, regB)}
}

func CMP(args []string, labels map[string]uint8) []uint8 {
	regA, regB := resolveReg(args)
	return []uint8{isa.OP_CMP, isa.PackRegs(regA, regB)}
}

func HALT(args []string, labels map[string]uint8) []uint8 {
	return []uint8{isa.OP_HALT}
}

func ADD(args []string, labels map[string]uint8) []uint8 {
	regA, regB := resolveReg(args)
	return []uint8{isa.OP_ADD, isa.PackRegs(regA, regB)}
}

func SUB(args []string, labels map[string]uint8) []uint8 {
	regA, regB := resolveReg(args)
	return []uint8{isa.OP_SUB, isa.PackRegs(regA, regB)}
}

func MUL(args []string, labels map[string]uint8) []uint8 {
	regA, regB := resolveReg(args)
	return []uint8{isa.OP_MUL, isa.PackRegs(regA, regB)}
}

func DIV(args []string, labels map[string]uint8) []uint8 {
	regA, regB := resolveReg(args)
	return []uint8{isa.OP_DIV, isa.PackRegs(regA, regB)}
}

func INC(args []string, labels map[string]uint8) []uint8 {
	return []uint8{isa.OP_INC, parseReg(args[0])}
}

func DEC(args []string, labels map[string]uint8) []uint8 {
	return []uint8{isa.OP_DEC, parseReg(args[0])}
}

func JMP(args []string, labels map[string]uint8) []uint8 {
	addr := resolveAddress(args[0], labels)
	return []uint8{isa.OP_JMP, addr}
}

func JZ(args []string, labels map[string]uint8) []uint8 {
	addr := resolveAddress(args[0], labels)
	return []uint8{isa.OP_JZ, addr}
}

func JNZ(args []string, labels map[string]uint8) []uint8 {
	addr := resolveAddress(args[0], labels)
	return []uint8{isa.OP_JNZ, addr}
}

func JG(args []string, labels map[string]uint8) []uint8 {
	addr := resolveAddress(args[0], labels)
	return []uint8{isa.OP_JG, addr}
}

func JL(args []string, labels map[string]uint8) []uint8 {
	addr := resolveAddress(args[0], labels)
	return []uint8{isa.OP_JL, addr}
}

func PUSH(args []string, labels map[string]uint8) []uint8 {
	return []uint8{isa.OP_PUSH, parseReg(args[0])}
}

func POP(args []string, labels map[string]uint8) []uint8 {
	return []uint8{isa.OP_POP, parseReg(args[0])}
}

func CALL(args []string, labels map[string]uint8) []uint8 {
	return []uint8{isa.OP_CALL, resolveAddress(args[0], labels)}
}

func RET(args []string, labels map[string]uint8) []uint8 {
	return []uint8{isa.OP_RET}
}

func PRINT(args []string, labels map[string]uint8) []uint8 {
	return []uint8{isa.OP_PRINT, parseReg(args[0])}
}
