package main

import (
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) == 2 {
		arg, _ := strconv.Atoi(os.Args[1])
		flag := IsPower(arg)
		if flag == 0 {
			Println("true")
		} else {
			Println("false")
		}
	}
}

func IsPower(n int) int {
	if n == 0 {
		return 1
	}
	return n & (n - 1)
}

func Println(s string) {
	for _, w := range s {
		z01.PrintRune(w)
	}
	z01.PrintRune('\n')
}
