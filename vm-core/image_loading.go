package vm_core

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"os"
)

func readImageFile(file *os.File, vm *LC3VM) {
	origin, _ := readWord(file)
	maxRead := MemoryMax - int(origin)

	for i := 0; i < maxRead; i++ {
		word, cnt := readWord(file)
		if cnt != 2 {
			break
		}
		vm.WriteMemory(uint16(i)+origin, word)
	}
}

func readWord(file *os.File) (uint16, int) {
	var word uint16

	byteArr := make([]byte, 2)
	cnt, err := file.Read(byteArr)
	if err != nil {
		fmt.Println(err)
		return 0, 0
	}
	buffer := bytes.NewBuffer(byteArr)
	err = binary.Read(buffer, binary.BigEndian, &word)
	if err != nil {
		fmt.Println(err)
		return 0, 0
	}

	return word, cnt
}

func swap16(rawWord uint16) uint16 {
	return (rawWord << 8) | (rawWord >> 8)
}
