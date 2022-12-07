package main

import (
	"fmt"
	lc3vm "gLC3VM/lc3-vm"
	"os"
)

func main() {
	vm := lc3vm.CreateVM()
	err := vm.LoadImageFromFile("C:\\Users\\jan_s\\Documents\\tomo_pi_mi\\project\\go_learning\\gLC3VM\\program_on_lc3\\hangman.obj")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	vm.Run()
}
