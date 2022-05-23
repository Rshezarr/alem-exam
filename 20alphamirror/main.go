package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	arg := os.Args[1]
	for _, w := range arg {
		if w >= 'a' && w <= 'z' {
			w = 122 - w + 97
		} else if w >= 'A' && w <= 'Z' {
			w = 90 - w + 65
		}
		z01.PrintRune(w)
	}
	z01.PrintRune('\n')
}
