package main

import "github.com/01-edu/z01"

func main() {
	for _, w := range "Hello World!" {
		z01.PrintRune(w)
	}
	z01.PrintRune('\n')
}
