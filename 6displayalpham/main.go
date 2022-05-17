package main

import "github.com/01-edu/z01"

// func main() {
// 	a := "aBcDeFgHiJkLmNoPqRsTuVwXyZ"

// 	for _, w := range a {
// 		z01.PrintRune(w)
// 	}
// 	z01.PrintRune('\n')
// }

// OR

func main() {
	for i := 'a'; i <= 'z'; i++ {
		if i%2 == 0 {
			n := i - 32
			z01.PrintRune(n)
		} else {
			z01.PrintRune(i)
		}
	}
	z01.PrintRune('\n')
}
