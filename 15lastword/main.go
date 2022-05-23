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

	var s string

	for i := len(arg) - 1; i >= 0; i-- {
		if arg[i] != ' ' {
			s = string(arg[i]) + s
		} else if arg[i] == ' ' && s != "" {
			break
		}
	}

	Println(s)
}

func Println(s string) {
	for _, w := range s {
		z01.PrintRune(w)
	}
	z01.PrintRune('\n')
}
