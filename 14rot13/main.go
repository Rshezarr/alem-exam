package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		return
	}

	for _, w := range os.Args[1] {
		if w >= 'a' && w <= 'm' || w >= 'A' && w <= 'M' {
			w += 13
		} else if w >= 'n' && w <= 'z' || w >= 'N' && w <= 'Z' {
			w -= 13
		}
		z01.PrintRune(w)
	}
	z01.PrintRune('\n')
}
