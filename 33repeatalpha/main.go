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
		for i := 0; i < OrderCount(w); i++ {
			z01.PrintRune(w)
		}
	}

	z01.PrintRune('\n')
}

func OrderCount(r rune) int {
	if r >= 'a' && r <= 'z' {
		return int(r - 96)
	} else if r >= 'A' && r <= 'Z' {
		return int(r - 64)
	}
	return 1
}
