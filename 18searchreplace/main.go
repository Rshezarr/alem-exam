package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) == 4 {
		l1 := os.Args[2][0]
		l2 := os.Args[3][0]
		word := os.Args[1]
		for _, w := range word {
			if w == rune(l1) {
				w = rune(l2)
			}
			z01.PrintRune(w)
		}
		z01.PrintRune('\n')
	}
}
