# TPU - Tuna Processing Unit

A custom 8-bit fantasy CPU emulator and assembler written in Go. Designed and implemented from scratch, including a bespoke instruction set architecture, a two-pass label-resolving assembler, and a cycle-accurate emulator with a full stack and subroutine model.

## Overview

TPU defines its own ISA, assembles `.tpu` source files into bytecode, and executes them on a virtual machine. The project covers the full toolchain — from source text to running instructions — with no external dependencies beyond the Go standard library.

## Architecture

- **8-bit registers and memory** — 16 general-purpose registers, 256 bytes of addressable memory
- **Separate stack space** — stack grows downward from `0xFF`, bounded at `0xC0` to prevent collision with program memory
- **Flags** — zero, greater, and less flags set by `CMP`, `ADD`, `SUB`, `INC`, `DEC`, and arithmetic operations
- **Subroutine support** — `CALL` pushes the return address onto the stack; `RET` restores it
- **Register packing** — two 4-bit register indices are packed into a single byte for two-operand instructions, keeping instruction sizes compact

## Instruction Set

| Instruction | Operands | Description |
|-------------|----------|-------------|
| `LOAD` | `Rd, imm` | Load immediate value into register |
| `LOADM` | `Rd, addr` | Load value from memory address into register |
| `STORE` | `Rs, addr` | Store register value to memory address |
| `MOV` | `Rd, Rs` | Copy register to register |
| `ADD` | `Rd, Rs` | Add Rs into Rd |
| `SUB` | `Rd, Rs` | Subtract Rs from Rd |
| `MUL` | `Rd, Rs` | Multiply Rd by Rs |
| `DIV` | `Rd, Rs` | Divide Rd by Rs (guards against divide-by-zero) |
| `INC` | `Rd` | Increment register |
| `DEC` | `Rd` | Decrement register |
| `CMP` | `Ra, Rb` | Compare two registers, set flags |
| `JMP` | `addr` | Unconditional jump |
| `JZ` | `addr` | Jump if zero flag set |
| `JNZ` | `addr` | Jump if zero flag not set |
| `JG` | `addr` | Jump if greater flag set |
| `JL` | `addr` | Jump if less flag set |
| `PUSH` | `Rs` | Push register onto stack |
| `POP` | `Rd` | Pop from stack into register |
| `CALL` | `addr` | Push return address, jump to subroutine |
| `RET` | | Return from subroutine |
| `PRINT` | `Rs` | Print register value to stdout |
| `HALT` | | Stop execution |

## Assembler

The assembler is a two-pass implementation:

1. **Tokenization** — strips comments (`;`), identifies labels, parses instructions and arguments
2. **Label resolution** — walks all instructions to compute byte offsets, building a label-to-address map before assembly begins
3. **Code generation** — emits bytecode using per-instruction handlers, resolving label references to concrete addresses

Registers are written as `R0`–`R15`. Numeric literals support decimal and hex (`0x` prefix). Labels are defined with a trailing colon and referenced by name in jump and call instructions.

## Project Structure

```
.
├── main.go                   Entry point — reads source file, assembles, and runs
├── src/
│   ├── isa/
│   │   └── isa.go            Opcode constants, register packing/unpacking, stack bounds
│   ├── assembler/
│   │   ├── assembler.go      Tokenizer, label resolver, and assembler driver
│   │   └── map.go            Per-instruction encoding handlers
│   └── tpu/
│       ├── tpu.go            CPU struct, fetch/decode loop, opcode dispatch table
│       └── components.go     Per-opcode execution handlers
```

## Usage

```bash
go run . path/to/program.tpu
```

## Built With

- Go (standard library only)