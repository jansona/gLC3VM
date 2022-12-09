package main

import (
	"flag"
	"fmt"
	lc3vm "gLC3VM/vm-core"
	"os"
)

func main() {
	imagePath := flag.String("image-path", "", "the path of vm image to load")
	flag.Parse()

	vm := lc3vm.CreateVM()
	err := vm.LoadImageFromFile(*imagePath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	vm.Run()
}
