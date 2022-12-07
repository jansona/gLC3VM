package lc3_vm

import (
	"bufio"
	"github.com/nsf/termbox-go"
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
	writer.Flush()
}

func processInput(vm *LC3VM) {
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			vm.AppendKeyBuffer(ev.Ch)
			if ev.Key == termbox.KeyCtrlC {
				vm.Terminate()
			}
		}
	}
}
