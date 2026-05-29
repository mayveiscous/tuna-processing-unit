package assembler

import (
	"fmt"
	"strings"
	"tpu/src/isa"
)

type Line struct {
	Instruction string
	Args        []string
}

type InstructionDef struct {
	Opcode  uint8
	Size    uint8
	Handler InstructionHandler
}

var Instructions = map[string]InstructionDef{
	"HALT":  {isa.OP_HALT, 1, HALT},
	"LOAD":  {isa.OP_LOAD, 3, LOAD},
	"ADD":   {isa.OP_ADD, 2, ADD},
	"SUB":   {isa.OP_SUB, 2, SUB},
	"MUL":   {isa.OP_MUL, 2, MUL},
	"DIV":   {isa.OP_DIV, 2, DIV},
	"MOV":   {isa.OP_MOV, 2, MOV},
	"CMP":   {isa.OP_CMP, 2, CMP},
	"JMP":   {isa.OP_JMP, 2, JMP},
	"JZ":    {isa.OP_JZ, 2, JZ},
	"JNZ":   {isa.OP_JNZ, 2, JNZ},
	"PUSH":  {isa.OP_PUSH, 2, PUSH},
	"POP":   {isa.OP_POP, 2, POP},
	"PRINT": {isa.OP_PRINT, 2, PRINT},
	"CALL":  {isa.OP_CALL, 2, CALL},
	"RET":   {isa.OP_RET, 1, RET},
	"STORE": {isa.OP_STORE, 3, STORE},
	"LOADM": {isa.OP_LOADM, 3, LOADM},
	"INC":   {isa.OP_INC, 2, INC},
	"DEC":   {isa.OP_DEC, 2, DEC},
	"JG":    {isa.OP_JG, 2, JG},
	"JL":    {isa.OP_JL, 2, JL},
}

func ResolveLabels(lines []Line) map[string]uint8 {
	labels := map[string]uint8{}
	var counter uint8 = 0

	for _, line := range lines {
		if strings.HasSuffix(line.Instruction, ":") {
			name := strings.TrimSuffix(line.Instruction, ":")
			labels[name] = counter
			continue
		}

		def, exists := Instructions[line.Instruction]
		if !exists {
			panic(fmt.Sprintf("unknown instruction '%s'", line.Instruction))
		}
		counter += def.Size
	}

	return labels
}

func Tokenize(source string) []Line {
	lines := []Line{}

	for _, raw := range strings.Split(source, "\n") {
		if idx := strings.Index(raw, ";"); idx != -1 {
			raw = raw[:idx]
		}

		raw = strings.TrimSpace(raw)
		if raw == "" {
			continue
		}

		fields := strings.Fields(raw)
		lines = append(lines, Line{
			Instruction: fields[0],
			Args:        fields[1:],
		})
	}

	return lines
}

func Assemble(lines []Line, labels map[string]uint8) ([]uint8, error) {
	bytes := []uint8{}

	for _, line := range lines {
		if strings.HasSuffix(line.Instruction, ":") {
			continue
		}

		def, exists := Instructions[line.Instruction]
		if !exists {
			return nil, fmt.Errorf("unknown instruction '%s'", line.Instruction)
		}

		bytes = append(bytes, def.Handler(line.Args, labels)...)
	}

	return bytes, nil
}
