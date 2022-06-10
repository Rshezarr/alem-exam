package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) == 1 {
		Print("options: abcdefghijklmnopqrstuvwxyz")
		return
	}

	res := make([]bool, 32)

	for _, arg := range os.Args[1:] {
		if arg[0] == '-' && len(arg) == 1 {
			Print("Invalid Option")
			return
		}

		if arg[:2] == "-h" {
			Print("options: abcdefghijklmnopqrstuvwxyz")
			return
		}

		for _, w := range arg[1:] {
			if w < 'a' || w > 'z' {
				Print("Invalid Option")
				return
			}
			res[31-(w-'a')] = true
		}
	}

	for i, w := range res {
		if w {
			z01.PrintRune('1')
		} else {
			z01.PrintRune('0')
		}

		if (i+1)%8 == 0 && i != len(res)-1 {
			z01.PrintRune(' ')
		}
	}
	z01.PrintRune('\n')
}

func Print(s string) {
	for _, w := range s {
		z01.PrintRune(w)
	}
	z01.PrintRune('\n')
}
