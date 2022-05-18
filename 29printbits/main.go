package main

import (
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) == 2 {
		arg, err := strconv.Atoi(os.Args[1])
		if err != nil {
			PrintReverse("00000000")
			return
		}

		var s string
		for len(s) != 8 {
			s = string(rune((arg%2)+48)) + s
			arg /= 2
		}

		PrintReverse(s)
	}
}

func PrintReverse(s string) {
	for _, w := range s {
		z01.PrintRune(w)
	}
	z01.PrintRune('\n')
}
