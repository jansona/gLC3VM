package glc3vm_test

import (
	vm_core "gLC3VM/vm-core"
	"testing"
)

var vm *vm_core.LC3VM

func TestInstruction(t *testing.T) {
	vm = vm_core.CreateVM()

	t.Run("test ADD one imm5", testOpAddOneImm5)
	t.Run("test ADD two register", testOpAddTwoRegister)
}

func testOpAddOneImm5(t *testing.T) {
	// Setup
	var addInstr uint16 = 0x1261
	var expected uint16 = 1

	// Execute & Verify
	vm_core.ExeCuteVmInstr(vm, addInstr)
	actual := vm.ReadRegister(vm_core.R_R1)
	if expected != actual {
		t.Errorf("ADD(%d, %d, Imm5: %d) = %d; expected: %d", vm_core.R_R1, vm_core.R_R1, 1, actual, expected)
	}

	// TearDown
	vm.Reset()
}

func testOpAddTwoRegister(t *testing.T) {
	// Setup
	var addInstr uint16 = 0x1042
	var expected uint16 = 3
	vm.WriteRegister(vm_core.R_R1, 1)
	vm.WriteRegister(vm_core.R_R2, 2)

	// Execute & Verify
	vm_core.ExeCuteVmInstr(vm, addInstr)
	actual := vm.ReadRegister(vm_core.R_R0)
	if expected != actual {
		t.Errorf("ADD(%d, %d, %d) = %d; expected: %d", vm_core.R_R0, vm_core.R_R1, vm_core.R_R2, actual, expected)
	}

	// TearDown
	vm.Reset()
}
