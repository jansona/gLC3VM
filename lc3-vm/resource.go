package lc3_vm

type LC3VM struct {
	memory []uint16
	reg    []uint16
}

func CreateVM() *LC3VM {
	vm := LC3VM{
		memory: make([]uint16, MemoryMax),
		reg:    make([]uint16, R_COUNT),
	}

	return &vm
}
