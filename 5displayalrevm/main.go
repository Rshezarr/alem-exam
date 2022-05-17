package main

import "github.com/01-edu/z01"

// func main() {
// 	a := "zYxWvUtSrQpOnMlKjIhGfEdCbA"

// 	for _, w := range a {
// 		z01.PrintRune(w)
// 	}
// 	z01.PrintRune('\n')
// }

// OR

func main() {
	for i := 'z'; i >= 'a'; i-- {
		if i%2 == 1 {
			n := i - 32
			z01.PrintRune(n)
		} else {
			z01.PrintRune(i)
		}
	}
	z01.PrintRune('\n')
}
