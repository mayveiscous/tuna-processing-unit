package main

import (
	"fmt"
	"os"
	"tpu/src/assembler"
	"tpu/src/tpu"
)

func main() {
	if len(os.Args) < 2 {
		 fmt.Println("Usage: tpu <file>")
		 os.Exit(1)
	}

	src, err := os.ReadFile(os.Args[1])
	if err != nil {
		 fmt.Println("failed to read file:", err)
		 os.Exit(1)
	}

	lines := assembler.Tokenize(string(src))
	labels := assembler.ResolveLabels(lines)
	program, err := assembler.Assemble(lines, labels)
	if err != nil {
		 fmt.Println("assembly error:", err)
		 os.Exit(1)
	}

	cpu := tpu.NewTPU()
	copy(cpu.Memory[:], program)
	cpu.Run()

	// fmt.Println("FINAL REGISTERS:")
	// fmt.Println(cpu)
}