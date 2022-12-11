package glc3vm

import (
	vm_core "gLC3VM/vm-core"
	"testing"
)

var vm *vm_core.LC3VM

func TestResource(t *testing.T) {
	vm = vm_core.CreateVM()

	t.Run("test Run with one NOT op", testRunWithOneNOTOP)
}

func testRunWithOneNOTOP(t *testing.T) {
	var notInstr uint16 = 0x907F
	var trapHaltInstr uint16 = 0xF025
	vm.StoreInMemory(0x3000, notInstr)
	vm.StoreInMemory(0x3001, trapHaltInstr)
	var expectedCalResult uint16 = 0xFFFE
	vm.WriteRegister(vm_core.R_R1, 1)

	vm.Run()
	actualCalResult := vm.ReadRegister(vm_core.R_R0)
	if expectedCalResult != actualCalResult {
		t.Errorf("NOT(%d, %d) = %d; expected: %d", vm_core.R_R0, vm_core.R_R1, actualCalResult, expectedCalResult)
	}

	vm.Reset()
}
