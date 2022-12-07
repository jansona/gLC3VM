package lc3_vm

import (
	"fmt"
	"io"
	"os"
)

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

	readImageFile(file, vm.memory)
	return nil
}

func readImageFile(file *os.File, memory []uint16) {
	origin, _ := readWord(file)
	maxRead := MemoryMax - int(origin)

	for i := 0; i < maxRead; i++ {
		word, cnt := readWord(file)
		if cnt != 2 {
			break
		}
		memory[i+int(origin)] = word
	}
}

func readWord(file *os.File) (uint16, int) {
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
	//tempWord = swap16(tempWord)
	return tempWord, cnt
}

func swap16(rawWord uint16) uint16 {
	return (rawWord << 8) | (rawWord >> 8)
}
