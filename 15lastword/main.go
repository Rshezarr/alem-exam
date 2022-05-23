package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	arg := os.Args[1]
	var fst int
	var last int

	for i := len(arg) - 1; i >= 0; i-- {
		if last != 0 && arg[i] == ' ' {
			fst = i + 1
			break
		}

		if arg[i] != ' ' {
			if last != 0 {
				continue
			}
			last = i + 1
		}
	}

	Println(arg[fst:last])
}

func Println(s string) {
	for _, w := range s {
		z01.PrintRune(w)
	}
	z01.PrintRune('\n')
}
