package lc3_vm

import (
	"fmt"
	"os"
)

type LC3VM struct {
	memory [MemoryMax]uint16
	reg    [R_COUNT]uint16
	status int
}

func CreateVM() *LC3VM {
	vm := LC3VM{
		status: INIT,
	}

	return &vm
}

func (vm *LC3VM) LoadImageFromFile(imagePath string) error {
	file, err := os.Open(imagePath)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}(file)

	readImageFile(file, vm)
	return nil
}

func (vm *LC3VM) Run() {
	vm.initialize()

	vm.WriteRegister(R_COND, FL_ZRO)
	vm.WriteRegister(R_PC, PC_START)

	vm.status = RUNNING
	for vm.status == RUNNING {
		instr := vm.loadInstr(vm.ReadRegister(R_PC))
		vm.WriteRegister(R_PC, vm.ReadRegister(R_PC)+1)
		ExeCuteVmInstr(vm, instr)
	}

	vm.finalize()
}

func (vm *LC3VM) ReadMemory(addr uint16) uint16 {
	return vm.memory[addr]
}

func (vm *LC3VM) WriteMemory(addr uint16, data uint16) {
	vm.memory[addr] = data
}

func (vm *LC3VM) ReadRegister(addr uint16) uint16 {
	return vm.reg[addr]
}

func (vm *LC3VM) WriteRegister(addr uint16, data uint16) {
	vm.reg[addr] = data
}

func (vm *LC3VM) UpdateVmFlag(regAddr uint16) {
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

func (vm *LC3VM) initialize() {

}

func (vm *LC3VM) finalize() {

}

func (vm *LC3VM) loadInstr(addr uint16) uint16 {
	if addr == MR_KBSR {
		if checkKey() != 0 {
			vm.WriteMemory(MR_KBSR, 1<<15)
			charByte := getChar()
			vm.WriteMemory(MR_KBDR, uint16(charByte))
		} else {
			vm.WriteMemory(MR_KBSR, 0)
		}
	}
	return vm.ReadMemory(addr)
}

// TODO
func checkKey() uint16 {
	return 0
}
