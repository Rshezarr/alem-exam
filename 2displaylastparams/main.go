package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 1 {
		for _, w := range os.Args[len(os.Args)-1] {
			z01.PrintRune(w)
		}
		z01.PrintRune('\n')
	} else {
		return
	}
}
