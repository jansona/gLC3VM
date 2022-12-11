package glc3vm_test

import (
	vm_core "gLC3VM/vm-core"
	"testing"
)

var vm *vm_core.LC3VM

func TestInstruction(t *testing.T) {
	vm = vm_core.CreateVM()

	t.Run("test ADD imm5", testOpAddImm5)
	t.Run("test ADD two register", testOpAddTwoRegister)

	t.Run("test AND imm5", testOpAndImm5)
	t.Run("test AND two register", testOpAndTwoRegister)

	t.Run("test NOT", testOpNot)
}

func testOpAddImm5(t *testing.T) {
	// Setup
	var instr uint16 = 0x1261
	var expected uint16 = 1

	// Execute & Verify
	vm_core.ExeCuteVmInstr(vm, instr)
	actual := vm.ReadRegister(vm_core.R_R1)
	if expected != actual {
		t.Errorf("ADD(%d, %d, Imm5: %d) = %d; expected: %d", vm_core.R_R1, vm_core.R_R1, 1, actual, expected)
	}

	// TearDown
	vm.Reset()
}

func testOpAddTwoRegister(t *testing.T) {
	// Setup
	var instr uint16 = 0x1042
	var expected uint16 = 3
	vm.WriteRegister(vm_core.R_R1, 1)
	vm.WriteRegister(vm_core.R_R2, 2)

	// Execute & Verify
	vm_core.ExeCuteVmInstr(vm, instr)
	actual := vm.ReadRegister(vm_core.R_R0)
	if expected != actual {
		t.Errorf("ADD(%d, %d, %d) = %d; expected: %d", vm_core.R_R0, vm_core.R_R1, vm_core.R_R2, actual, expected)
	}

	// TearDown
	vm.Reset()
}

func testOpAndImm5(t *testing.T) {
	var instr uint16 = 0x5021
	var expected uint16 = 1
	vm.WriteRegister(vm_core.R_R0, 3)

	vm_core.ExeCuteVmInstr(vm, instr)
	actual := vm.ReadRegister(vm_core.R_R0)
	if expected != actual {
		t.Errorf("AND(%d, %d, Imm5: %d) = %d; expected: %d", vm_core.R_R0, vm_core.R_R0, 1, actual, expected)
	}

	vm.Reset()
}

func testOpAndTwoRegister(t *testing.T) {
	var instr uint16 = 0x5001
	var expected uint16 = 3
	vm.WriteRegister(vm_core.R_R0, 7)
	vm.WriteRegister(vm_core.R_R1, 3)

	vm_core.ExeCuteVmInstr(vm, instr)
	actual := vm.ReadRegister(vm_core.R_R0)
	if expected != actual {
		t.Errorf("AND(%d, %d, %d) = %d; expected: %d", vm_core.R_R0, vm_core.R_R0, vm_core.R_R1, actual, expected)
	}

	vm.Reset()
}

func testOpNot(t *testing.T) {
	var instr uint16 = 0x907F
	var expected uint16 = 0xFFFE
	vm.WriteRegister(vm_core.R_R1, 1)

	vm_core.ExeCuteVmInstr(vm, instr)
	actual := vm.ReadRegister(vm_core.R_R0)
	if expected != actual {
		t.Errorf("NOT(%d, %d) = %d; expected: %d", vm_core.R_R0, vm_core.R_R1, actual, expected)
	}

	vm.Reset()
}
