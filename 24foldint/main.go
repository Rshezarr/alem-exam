package main

import (
	"fmt"

	"github.com/01-edu/z01"
)

func main() {
	table := []int{21, 31, 20, 41}
	ac := 93
	FoldInt(Add, table, ac)
	FoldInt(Mul, table, ac)
	FoldInt(Sub, table, ac)
	fmt.Println()

	table = []int{0}
	FoldInt(Add, table, ac)
	FoldInt(Mul, table, ac)
	FoldInt(Sub, table, ac)
}

// Start

func FoldInt(f func(int, int) int, a []int, n int) {
	for i := 0; i < len(a); i++ {
		n = f(n, a[i])
	}

	for _, w := range Itoa(n) {
		z01.PrintRune(w)
	}
	z01.PrintRune('\n')
}

func Itoa(n int) string {
	var ch string
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

// End

func Add(a, b int) int {
	return a + b
}

func Sub(a, b int) int {
	return a - b
}

func Mul(a, b int) int {
	return a * b
}
