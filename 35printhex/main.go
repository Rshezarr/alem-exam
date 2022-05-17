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
			Print("ERROR")
			return
		}

		hex := strconv.FormatInt(int64(arg), 16)
		Print(hex)
	}
}

func Print(s string) {
	for _, w := range s {
		z01.PrintRune(w)
	}
}
