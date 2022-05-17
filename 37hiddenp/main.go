package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) == 3 {
		var arr []rune
		a1 := []rune(os.Args[1])
		a2 := []rune(os.Args[2])
		for i := 0; i < len(a1); i++ {
			for j := 0; j < len(a2); j++ {
				if a1[i] == a2[j] {
					arr = append(arr, a2[j])
					a2[j] = rune(0)
					break
				}
				a2[j] = rune(0)
			}
		}

		if len(a1) != len(arr) {
			z01.PrintRune('0')
		} else {
			z01.PrintRune('1')
		}
	}
}
