package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 3 {
		return
	}
	arg1 := []rune(os.Args[1])
	arg2 := []rune(os.Args[2])

	var str []rune

	for i := 0; i < len(arg1); i++ {
		for j := 0; j < len(arg2); j++ {
			if arg1[i] == arg2[j] {
				str = append(str, arg2[j])
				arg2[j] = 0
				break
			}
			arg2[j] = 0
		}
	}

	if len(str) != len(arg1) {
		return
	}

	mapStr := make(map[rune]bool)

	for _, w := range str {
		if !mapStr[w] {
			z01.PrintRune(w)
			mapStr[w] = true
		}
	}
	z01.PrintRune('\n')
}
