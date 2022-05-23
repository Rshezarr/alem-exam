package main

import (
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	arg, err := strconv.Atoi(os.Args[1])
	if err != nil {
		Print("ERROR")
		return
	}

	hex := strconv.FormatInt(int64(arg), 16)

	// var s string
	// hex := "0123456789abcdef"
	// for arg > 0 {
	// 	s = string(hex[arg%16]) + s
	// 	arg /= 16
	// }
	Print(hex)
}

func Print(s string) {
	for _, w := range s {
		z01.PrintRune(w)
	}
	z01.PrintRune('\n')
}
