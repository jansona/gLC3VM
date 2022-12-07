package lc3_vm

import "fmt"

func handleTrap(vm *LC3VM, instr uint16) {
	vm.WriteRegister(R_R7, vm.ReadRegister(R_PC))

	switch instr & 0xFF {
	case TRAP_GETC:
		trapGETC(vm)
	case TRAP_OUT:
		trapOUT(vm)
	case TRAP_PUTS:
		trapPUTS(vm)
	case TRAP_IN:
		trapIN(vm)
	case TRAP_PUTSP:
		trapPUTSP(vm)
	case TRAP_HALT:
		trapHALT(vm)
	}
}

func trapGETC(vm *LC3VM) {
	vm.WriteRegister(R_R0, uint16(getChar()))
	vm.UpdateVmFlag(R_R0)
}

func trapOUT(vm *LC3VM) {
	putChar(rune(vm.ReadRegister(R_R0)))
	// TODO 需要flush stdout吗？
}

func trapPUTS(vm *LC3VM) {
	strAddr := vm.ReadRegister(R_R0)
	charData := vm.ReadMemory(strAddr)
	for charData != 0x0 {
		putChar(rune(charData))
		strAddr++
	}
	// TODO 需要flush stdout吗？
}

func trapIN(vm *LC3VM) {
	fmt.Println("Enter a character: ")
	char := getChar()
	putChar(char)
	// TODO 需要flush stdout吗？
	vm.WriteRegister(R_R0, uint16(char))
	vm.UpdateVmFlag(R_R0)
}

func trapPUTSP(vm *LC3VM) {
	strAddr := vm.ReadRegister(R_R0)
	charData := vm.ReadMemory(strAddr)
	for charData != 0x0 {
		putChar(rune(charData & 0xFF))
		if charData>>8 != 0 {
			putChar(rune(charData >> 8))
		}
		strAddr++
	}
	// TODO 需要flush stdout吗？
}

func trapHALT(vm *LC3VM) {
	fmt.Println("HALT")
	// TODO 需要flush stdout吗？
	vm.status = HALT
}
