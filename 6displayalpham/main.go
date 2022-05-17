package main

import "github.com/01-edu/z01"

func main() {
	a := "aBcDeFgHiJkLmNoPqRsTuVwXyZ"

	for _, w := range a {
		z01.PrintRune(w)
	}
	z01.PrintRune('\n')
}
