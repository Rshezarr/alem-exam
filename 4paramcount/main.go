package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	a := rune(len(os.Args)-1) + 48
	z01.PrintRune(a)
	z01.PrintRune('\n')
}
