package lc3_vm

// register & memory
const (
	R_R0 = iota
	R_R1
	R_R2
	R_R3
	R_R4
	R_R5
	R_R6
	R_R7
	R_PC
	R_COND
	R_COUNT
)

const (
	MemoryMax int = 1 << 16

	MR_KBSR = 0xFE00
	MR_KBDR = 0xFE02
)

// instruction & trap
const (
	OP_BR = iota
	OP_ADD
	OP_LD
	OP_ST
	OP_JSR
	OP_AND
	OP_LDR
	OP_STR
	OP_RTI
	OP_NOT
	OP_LDI
	OP_STI
	OP_JMP
	OP_RES
	OP_LEA
	OP_TRAP
)

const (
	TRAP_GETC = iota + 32
	TRAP_OUT
	TRAP_PUTS
	TRAP_IN
	TRAP_PUTSP
	TRAP_HALT
)
