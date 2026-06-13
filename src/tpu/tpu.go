package tpu

import (
	"fmt"
	"tpu/src/isa"
)

type TPU struct {
	registers [16]uint8
	Memory    [256]uint8
	pc        uint8
	sp        uint8
	zero      bool
	greater   bool
	less      bool
	running   bool
}

var opcodeTable = map[uint8]OpcodeHandler{
	isa.OP_LOAD:  LOAD,
	isa.OP_ADD:   ADD,
	isa.OP_SUB:   SUB,
	isa.OP_MUL:   MUL,
	isa.OP_DIV:   DIV,
	isa.OP_MOV:   MOV,
	isa.OP_CMP:   CMP,
	isa.OP_JMP:   JMP,
	isa.OP_JZ:    JZ,
	isa.OP_JNZ:   JNZ,
	isa.OP_JG:    JG,
	isa.OP_JL:    JL,
	isa.OP_PUSH:  PUSH,
	isa.OP_POP:   POP,
	isa.OP_PRINT: PRINT,
	isa.OP_HALT:  HALT,
	isa.OP_CALL:  CALL,
	isa.OP_STORE: STORE,
	isa.OP_LOADM: LOADM,
	isa.OP_RET:   RET,
	isa.OP_INC:   INC,
	isa.OP_DEC:   DEC,
}

func NewTPU() *TPU {
	return &TPU{
		sp: isa.StackStart,
	}
}

func (cpu *TPU) fetch() uint8 {
	val := cpu.Memory[cpu.pc]
	cpu.pc++
	return val
}

func (cpu *TPU) decodeRegs() (uint8, uint8) {
	arg := cpu.fetch()

	regA, regB := isa.UnpackRegs(arg)
	cpu.validateRegs(regA, regB)
	return regA, regB
}

func (cpu *TPU) validateRegs(regA uint8, regB uint8) {
	if regA >= uint8(len(cpu.registers)) || regB >= uint8(len(cpu.registers)) {
		panic(fmt.Sprintf("invalid register(s): R%d R%d", regA, regB))
	}
}

func (cpu *TPU) validateSingle(reg uint8) {
	if reg >= uint8(len(cpu.registers)) {
		panic(fmt.Sprintf("invalid register: R%d", reg))
	}
}

func (cpu *TPU) Run() {
	cpu.running = true

	for cpu.running {
		opcode := cpu.Memory[cpu.pc]
		cpu.pc++
		cpu.Execute(opcode)
	}
}

func (cpu *TPU) Execute(opcode uint8) {
	component, exists := opcodeTable[opcode]

	if !exists {
		panic(fmt.Sprintf("unknown opcode %02X", opcode))
	}

	component(cpu)
}