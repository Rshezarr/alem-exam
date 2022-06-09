package main

import (
	"github.com/01-edu/z01"
)

func main() {
	a := "9876543210"

	for _, w := range a {
		z01.PrintRune(w)
	}

	//OR
	// for i := 9; i >= 0; i-- {
	// 	z01.PrintRune(rune(i + 48))
	// }
	z01.PrintRune('\n')
}
