package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 3 {
		return
	}
	arg1 := os.Args[1]
	arg2 := os.Args[2]

	args := arg1 + arg2

	mapC := make(map[rune]bool)

	for _, w := range args {
		if !mapC[w] {
			z01.PrintRune(w)
			mapC[w] = true
		}
	}
	z01.PrintRune('\n')
}
