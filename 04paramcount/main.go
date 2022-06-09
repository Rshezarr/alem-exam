package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	a := BasicItoa(len(os.Args) - 1)
	for _, v := range a {
		z01.PrintRune(v)
	}
	z01.PrintRune('\n')
}

func BasicItoa(n int) string {
	base := "0123456789"
	var res string
	if n < 10 {
		return res + base[n:n+1]
	}

	return res + BasicItoa(n/10) + base[n%10:n%10+1]
}
