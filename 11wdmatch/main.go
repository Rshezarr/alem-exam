package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) == 3 {
		arg1 := []rune(os.Args[1])
		arg2 := []rune(os.Args[2])
		var arr []rune

		for i := 0; i < len(arg1); i++ {
			for j := 0; j < len(arg2); j++ {
				if arg1[i] == arg2[j] {
					arr = append(arr, arg2[j])
					arg2[j] = rune(0)
					break
				}
				arg2[j] = rune(0)
			}
		}

		if len(arg1) != len(arr) {
			return
		}

		for _, w := range arr {
			z01.PrintRune(w)
		}
		z01.PrintRune('\n')
	} else {
		return
	}
}
