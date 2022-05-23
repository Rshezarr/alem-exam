package main

import (
	"fmt"

	"github.com/01-edu/z01"
)

func main() {
	table := []int{1, 2, 3}
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
	var sum int
	for _, w := range a {
		sum += w
	}

	res := f(n, sum)

	for _, w := range Itoa(res) {
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
