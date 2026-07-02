package tpu

import "fmt"

type CommandHandler func(t *TPU)

func ExecuteCommands(cmds []string, t *TPU) {
    for _, cmd := range cmds {
        if fn, ok := Commands[cmd]; ok {
            fn(t)
        } else {
            fmt.Printf("unknown command: %s\n", cmd)
        }
    }
}

func Debug(t *TPU) {
	fmt.Println("=== TPU DEBUG ===")
	fmt.Printf("PC: 0x%02X   SP: 0x%02X\n", t.pc, t.sp)
	fmt.Printf("Flags: Z=%v G=%v L=%v\n", t.zero, t.greater, t.less)
	fmt.Printf("Running: %v\n", t.running)
	fmt.Println("Registers:")
	for i, r := range t.registers {
		fmt.Printf("  R%02d = 0x%02X (%d)\n", i, r, r)
	}
	fmt.Println("=================")
}

func Registers(t *TPU) {
	fmt.Println("=== REGISTERS ===")
	for i, r := range t.registers {
		fmt.Printf("  R%02d = 0x%02X (%d)\n", i, r, r)
	}
	fmt.Println("=================")
}

func Stack(t *TPU) {
	fmt.Println("=== STACK ===")
	for addr := uint8(0xFF); addr > t.sp; addr-- {
		fmt.Printf("  0x%02X: 0x%02X (%d)\n", addr, t.Memory[addr], t.Memory[addr])
	}
	if t.sp == 0xFF {
		fmt.Println("  (empty)")
	}
	fmt.Println("=============")
}

func DumpMemory(t *TPU) {
	fmt.Println("=== MEMORY ===")
	for i := 0; i < 256; i += 16 {
		fmt.Printf("  0x%02X: ", i)
		for j := 0; j < 16; j++ {
			fmt.Printf("%02X ", t.Memory[i+j])
		}
		fmt.Println()
	}
	fmt.Println("==============")
}

func Step(t *TPU) {
	if !t.running {
		fmt.Println("TPU is not running")
		return
	}
	t.Execute(t.Memory[t.pc])
	fmt.Printf("Stepped — PC now at 0x%02X\n", t.pc)
}

func Reset(t *TPU) {
	t.registers = [16]uint8{}
	t.pc = 0
	t.sp = 0xFF
	t.zero = false
	t.greater = false
	t.less = false
	t.running = false
	fmt.Println("TPU reset.")
}

func Status(t *TPU) {
	fmt.Printf("PC: 0x%02X  SP: 0x%02X  Z=%v G=%v L=%v  Running: %v\n",
		t.pc, t.sp, t.zero, t.greater, t.less, t.running)
}