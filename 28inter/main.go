package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) == 3 {
		arg1 := os.Args[1]
		arg2 := os.Args[2]

		var str string

		for i := 0; i < len(arg1); i++ {
			for j := 0; j < len(arg2); j++ {
				if arg1[i] == arg2[j] {
					str += string(arg2[j])
				}
			}
		}

		mapStr := make(map[rune]bool)

		for _, w := range str {
			mapStr[w] = true
		}

		for _, w := range str {
			if mapStr[w] {
				z01.PrintRune(w)
				mapStr[w] = false
			}
		}
		z01.PrintRune('\n')
	}
}
