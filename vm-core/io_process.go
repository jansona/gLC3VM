package vm_core

import (
	"bufio"
	"fmt"
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
	fmt.Printf("%c", char)
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
