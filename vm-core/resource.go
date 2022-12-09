package vm_core

import (
	"fmt"
	"github.com/nsf/termbox-go"
	"os"
)

type LC3VM struct {
	memory    [MemoryMax]uint16
	reg       [R_COUNT]uint16
	keyBuffer []rune
	status    int
}

func CreateVM() *LC3VM {
	vm := LC3VM{
		status:    INIT,
		keyBuffer: make([]rune, 0),
	}

	return &vm
}

func (vm *LC3VM) LoadImageFromFile(imagePath string) error {
	file, err := os.Open(imagePath)
	if err != nil {
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
		vm.handleInput()

		instr := vm.loadInstr(vm.ReadRegister(R_PC))
		vm.WriteRegister(R_PC, vm.ReadRegister(R_PC)+1)
		ExeCuteVmInstr(vm, instr)
	}

	defer vm.finalize()
}

func (vm *LC3VM) Terminate() {
	vm.status = HALT
}

func (vm *LC3VM) Reset() {
	for _, idx := range vm.reg {
		vm.WriteRegister(idx, 0)
	}
	for _, idx := range vm.memory {
		vm.StoreInMemory(idx, 0)
	}
	vm.status = INIT
	vm.keyBuffer = make([]rune, 0)
}

func (vm *LC3VM) ReadRegister(addr uint16) uint16 {
	return vm.reg[addr]
}

func (vm *LC3VM) WriteRegister(addr uint16, data uint16) {
	vm.reg[addr] = data
}

func (vm *LC3VM) LoadFromMemory(addr uint16) uint16 {
	return vm.memory[addr]
}

func (vm *LC3VM) StoreInMemory(addr uint16, data uint16) {
	vm.memory[addr] = data
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

func (vm *LC3VM) AppendKeyBuffer(key rune) {
	vm.keyBuffer = append(vm.keyBuffer, key)
}

func (vm *LC3VM) PopKeyBuffer() rune {
	key := vm.keyBuffer[0]
	vm.keyBuffer = vm.keyBuffer[1:]
	return key
}

func (vm *LC3VM) initialize() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}

	go processInput(vm)
}

func (vm *LC3VM) finalize() {
	termbox.Close()
}

func (vm *LC3VM) handleInput() {
	kbsrVal := vm.LoadFromMemory(MR_KBSR)
	if (kbsrVal&0x8000 == 0) && len(vm.keyBuffer) > 0 {
		vm.StoreInMemory(MR_KBSR, kbsrVal|0x8000)
		vm.StoreInMemory(MR_KBDR, uint16(vm.keyBuffer[0]))
	}
}

func (vm *LC3VM) loadInstr(addr uint16) uint16 {
	if addr == MR_KBDR {
		vm.StoreInMemory(MR_KBSR, vm.LoadFromMemory(MR_KBSR)&0x7FFF)
	}
	return vm.LoadFromMemory(addr)
}
