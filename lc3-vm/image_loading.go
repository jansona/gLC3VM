package lc3_vm

import (
	"fmt"
	"io"
	"os"
)

func readImageFile(file *os.File, vm *LC3VM, needConvert bool) {
	origin, _ := readWord(file, needConvert)
	maxRead := MemoryMax - int(origin)

	for i := 0; i < maxRead; i++ {
		word, cnt := readWord(file, needConvert)
		if cnt != 2 {
			break
		}
		vm.WriteMemory(uint16(i)+origin, word)
	}
}

func readWord(file *os.File, needConvert bool) (uint16, int) {
	buffer := make([]byte, 2)
	cnt, err := file.Read(buffer)
	if err != nil && err != io.EOF {
		fmt.Println(err)
		os.Exit(1)
	}
	if cnt != 2 {
		return 0, cnt
	}
	tempWord := (uint16(buffer[0]) << 8) | (uint16(buffer[1]))
	if needConvert {
		tempWord = swap16(tempWord)
	}
	return tempWord, cnt
}

func swap16(rawWord uint16) uint16 {
	return (rawWord << 8) | (rawWord >> 8)
}
