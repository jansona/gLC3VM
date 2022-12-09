package vm_core

import (
	"errors"
	"fmt"
	"os"
)

func ExeCuteVmInstr(vm *LC3VM, instr uint16) {
	opcode := instr >> 12

	switch opcode {
	case OP_ADD:
		opADD(vm, instr)
	case OP_AND:
		opAND(vm, instr)
	case OP_NOT:
		opNOT(vm, instr)
	case OP_BR:
		opBR(vm, instr)
	case OP_JMP:
		opJMP(vm, instr)
	case OP_JSR:
		opJSR(vm, instr)
	case OP_LD:
		opLD(vm, instr)
	case OP_LDI:
		opLDI(vm, instr)
	case OP_LDR:
		opLDR(vm, instr)
	case OP_LEA:
		opLEA(vm, instr)
	case OP_ST:
		opST(vm, instr)
	case OP_STI:
		opSTI(vm, instr)
	case OP_STR:
		opSTR(vm, instr)
	case OP_TRAP:
		handleTrap(vm, instr)
	case OP_RES:
	case OP_RTI:
	default:
		fmt.Println(errors.New(fmt.Sprintf("unknow opcode %x", opcode)))
		os.Exit(1)
	}
}

func opADD(vm *LC3VM, instr uint16) {
	r0 := (instr >> 9) & 0x7
	r1 := (instr >> 6) & 0x7
	immFlag := (instr >> 5) & 0x1

	result := vm.ReadMemory(r1)
	if immFlag != 0 {
		imm5 := signExtend(instr&0x1F, 5)
		result += imm5
	} else {
		r2 := instr & 0x7
		result += vm.ReadRegister(r2)
	}
	vm.WriteRegister(r0, result)

	vm.UpdateVmFlag(r0)
}

func opAND(vm *LC3VM, instr uint16) {
	r0 := (instr >> 9) & 0x7
	r1 := (instr >> 6) & 0x7
	immFlag := (instr >> 5) & 0x1

	result := vm.ReadRegister(r1)
	if immFlag != 0 {
		imm5 := signExtend(instr&0x1F, 5)
		result &= imm5
	} else {
		r2 := instr & 0x7
		result &= vm.ReadRegister(r2)
	}
	vm.WriteRegister(r0, result)

	vm.UpdateVmFlag(r0)
}

func opNOT(vm *LC3VM, instr uint16) {
	r0 := (instr >> 9) & 0x7
	r1 := (instr >> 6) & 0x7

	vm.WriteRegister(r0, ^vm.ReadRegister(r1))
	vm.UpdateVmFlag(r0)
}

func opBR(vm *LC3VM, instr uint16) {
	pcOffset := signExtend(instr&0x1FF, 9)
	condition := (instr >> 9) & 0x7
	if condition&vm.ReadRegister(R_COND) != 0 {
		vm.WriteRegister(R_PC, vm.ReadRegister(R_PC)+pcOffset)
	}
}

func opJMP(vm *LC3VM, instr uint16) {
	r1 := (instr >> 6) & 0x7
	vm.WriteRegister(R_PC, vm.ReadRegister(r1))
}

func opJSR(vm *LC3VM, instr uint16) {
	flag := (instr >> 11) & 0x1
	vm.WriteRegister(R_R7, vm.ReadRegister(R_PC))
	if flag != 0 {
		pcOffset := signExtend(instr&0x7FF, 11)
		vm.WriteRegister(R_PC, vm.ReadRegister(R_PC)+pcOffset)
	} else {
		r1 := (instr >> 6) & 0x7
		vm.WriteRegister(R_PC, vm.ReadRegister(r1))
	}
}

func opLD(vm *LC3VM, instr uint16) {
	r0 := (instr >> 9) & 0x7
	pcOffset := signExtend(instr&0x1FF, 9)
	addr := vm.ReadRegister(R_PC) + pcOffset
	vm.WriteRegister(r0, vm.ReadMemory(addr))

	vm.UpdateVmFlag(r0)
}

func opLDR(vm *LC3VM, instr uint16) {
	r0 := (instr >> 9) & 0x7
	r1 := (instr >> 6) & 0x7
	offset := signExtend(instr&0x3F, 6)
	addr := vm.ReadRegister(r1) + offset
	vm.WriteRegister(r0, vm.ReadMemory(addr))

	vm.UpdateVmFlag(r0)
}

func opLEA(vm *LC3VM, instr uint16) {
	r0 := (instr >> 9) & 0x7
	pcOffset := signExtend(instr&0x1FF, 9)
	addr := vm.ReadRegister(R_PC) + pcOffset
	vm.WriteRegister(r0, addr)

	vm.UpdateVmFlag(r0)
}

func opST(vm *LC3VM, instr uint16) {
	r0 := (instr >> 9) & 0x7
	pcOffset := signExtend(instr&0x1FF, 9)
	addr := vm.ReadRegister(R_PC) + pcOffset

	vm.WriteMemory(addr, vm.ReadRegister(r0))
}

func opSTI(vm *LC3VM, instr uint16) {
	r0 := (instr >> 9) & 0x7
	pcOffset := signExtend(instr&0x1FF, 9)
	indirectAddr := vm.ReadRegister(R_PC) + pcOffset
	addr := vm.ReadMemory(indirectAddr)

	vm.WriteMemory(addr, vm.ReadRegister(r0))
}

func opSTR(vm *LC3VM, instr uint16) {
	r0 := (instr >> 9) & 0x7
	r1 := (instr >> 6) & 0x7
	offset := signExtend(instr&0x3F, 6)
	addr := vm.ReadRegister(r1) + offset

	vm.WriteMemory(addr, vm.ReadRegister(r0))
}

func opLDI(vm *LC3VM, instr uint16) {
	r0 := (instr >> 9) & 0x7
	pcOffset := signExtend(instr&0x1FF, 9)
	indirectAddr := vm.ReadRegister(R_PC) + pcOffset
	finalAddr := vm.ReadMemory(indirectAddr)
	vm.WriteRegister(r0, vm.ReadMemory(finalAddr))

	vm.UpdateVmFlag(r0)
}

func signExtend(data uint16, bitWise int) uint16 {
	if (data>>(bitWise-1))&1 != 0 {
		data |= 0xFFFF << bitWise
	}
	return data
}
