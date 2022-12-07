package main

import (
	"fmt"
	lc3_vm "gLC3VM/lc3-vm"
	"os"
)

func main() {
	vm := lc3_vm.CreateVM()
	err := vm.LoadImageFromFile("C:\\Users\\jan_s\\Documents\\tomo_pi_mi\\project\\go_learning\\gLC3VM\\program_on_lc3\\2048.obj")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
