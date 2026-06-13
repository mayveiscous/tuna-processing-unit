package tpu

import (
	"fmt"
	"tpu/src/isa"
)

type OpcodeHandler func(cpu *TPU)

func LOAD(cpu *TPU) {
	reg := cpu.fetch()
	value := cpu.fetch()
	cpu.validateSingle(reg)
	cpu.registers[reg] = value
}

func ADD(cpu *TPU) {
	regA, regB := cpu.decodeRegs()

	result := cpu.registers[regA] + cpu.registers[regB]
	cpu.registers[regA] = result
	cpu.zero = result == 0
}

func SUB(cpu *TPU) {
	regA, regB := cpu.decodeRegs()

	result := cpu.registers[regA] - cpu.registers[regB]
	cpu.registers[regA] = result
	cpu.zero = result == 0
}

func MUL(cpu *TPU) {
	regA, regB := cpu.decodeRegs()

	result := cpu.registers[regA] * cpu.registers[regB]
	cpu.registers[regA] = result
	cpu.zero = result == 0
}

func DIV(cpu *TPU) {
	regA, regB := cpu.decodeRegs()

	if cpu.registers[regB] == 0 {
		fmt.Println("DIV: attempt to divide by zero...")
		return
	}
	
	result := cpu.registers[regA] / cpu.registers[regB]
	cpu.registers[regA] = result
	cpu.zero = result == 0
}

func MOV(cpu *TPU) {
	regA, regB := cpu.decodeRegs()
	cpu.registers[regA] = cpu.registers[regB]
}

func CMP(cpu *TPU) {
	regA, regB := cpu.decodeRegs()

	cpu.zero = cpu.registers[regA] == cpu.registers[regB]
	result := cpu.registers[regA] > cpu.registers[regB]

	cpu.greater = result
	cpu.less = !result
}

func JMP(cpu *TPU) {
	addr := cpu.fetch()
	cpu.pc = addr
}

func JZ(cpu *TPU) {
	addr := cpu.fetch()

	if cpu.zero {
		cpu.pc = addr
	}
}

func JNZ(cpu *TPU) {
	addr := cpu.fetch()

	if !cpu.zero {
		cpu.pc = addr
	}
}

func JG(cpu *TPU) {
	addr := cpu.fetch()

	if cpu.greater {
		cpu.pc = addr
	}
}

func JL(cpu *TPU) {
	addr := cpu.fetch()

	if cpu.less == true {
		cpu.pc = addr
	}
}

func PUSH(cpu *TPU) {
	reg := cpu.fetch()
	cpu.validateSingle(reg)

	if cpu.sp == isa.StackLimit {
		panic("stack overflow")
	}

	cpu.Memory[cpu.sp] = cpu.registers[reg]
	cpu.sp--
}

func POP(cpu *TPU) {
	reg := cpu.fetch()
	cpu.validateSingle(reg)

	if cpu.sp == isa.StackStart {
		panic("stack underflow")
	}

	cpu.sp++
	cpu.registers[reg] = cpu.Memory[cpu.sp]
}

func HALT(cpu *TPU) {
	cpu.running = false
}

func PRINT(cpu *TPU) {
	reg := cpu.fetch()
	cpu.validateSingle(reg)
	fmt.Printf("R%d = %d\n", reg, cpu.registers[reg])
}

func CALL(cpu *TPU) {
	address := cpu.fetch()
	cpu.Memory[cpu.sp] = cpu.pc
	cpu.sp--
	cpu.pc = address
}

func RET(cpu *TPU) {
	cpu.pc = cpu.Memory[cpu.sp]
	cpu.sp++
}

func LOADM(cpu *TPU) {
	reg := cpu.fetch()
	addr := cpu.fetch()

	cpu.validateSingle(reg)

	cpu.registers[reg] = cpu.Memory[addr]
}

func STORE(cpu *TPU) {
	reg := cpu.fetch()
	addr := cpu.fetch()

	cpu.validateSingle(reg)

	cpu.Memory[addr] = cpu.registers[reg]
}

func INC(cpu *TPU) {
	reg := cpu.fetch()
	cpu.validateSingle(reg)

	cpu.registers[reg]++
	cpu.zero = cpu.registers[reg] == 0
}

func DEC(cpu *TPU) {
	reg := cpu.fetch()
	cpu.validateSingle(reg)

	cpu.registers[reg]--
	cpu.zero = cpu.registers[reg] == 0
}
