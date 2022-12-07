package lc3_vm

func ExeCuteVmInstr(vm *LC3VM, instr uint16) {
	opcode := instr >> 12

	switch opcode {
	case OP_ADD:
		opADD(vm, instr)
	case OP_AND:
	case OP_NOT:
	case OP_BR:
	case OP_JMP:
	case OP_JSR:
	case OP_LD:
	case OP_LDI:
		opLDI(vm, instr)
	case OP_LDR:
	case OP_LEA:
	case OP_ST:
	case OP_STI:
	case OP_TRAP:
	case OP_RES:
	case OP_RTI:
	default:

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
}

func opLDI(vm *LC3VM, instr uint16) {
	r0 := (instr >> 9) & 0x7
	pcOffset := signExtend(instr&0x1FF, 9)
	indirectAddr := vm.ReadRegister(R_PC) + pcOffset
	finalAddr := vm.ReadMemory(indirectAddr)
	vm.WriteRegister(r0, vm.ReadMemory(finalAddr))
}

func UpdateVmFlag(vm *LC3VM, regAddr uint16) {
	regData := vm.ReadRegister(regAddr)
	condition := FL_POS
	switch {
	case regData == 0:
		condition = FL_ZRO
	case regData>>15 != 0:
		condition = FL_NEG
	}
	vm.WriteRegister(R_COND, uint16(condition))
}

func signExtend(data uint16, bitWise int) uint16 {
	if (data>>(bitWise-1))&1 != 0 {
		data |= 0xFFFF << bitWise
	}
	return data
}
