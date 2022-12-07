package lc3_vm

import (
	"bufio"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func getChar() rune {
	charByte, _ := reader.ReadByte()
	return rune(charByte)
}

func putChar(char rune) {
	writer.WriteRune(char)
}
