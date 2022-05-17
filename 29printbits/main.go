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
			s += Itoa(arg % 2)
			arg /= 2
		}

		PrintReverse(s)
		z01.PrintRune('\n')
	}
}

func Itoa(n int) string {
	ch := ""
	if n < 0 {
		n = -n
		ch = "-"
	}
	digits := "0123456789"
	if n < 10 {
		return ch + digits[n:n+1]
	}
	return ch + Itoa(n/10) + digits[n%10:n%10+1]
}

func PrintReverse(s string) {
	res := ""
	for _, w := range s {
		res = string(w) + res
	}

	for _, e := range res {
		z01.PrintRune(e)
	}
}
