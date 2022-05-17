package main

import (
	"strconv"

	"github.com/01-edu/z01"
)

func main() {
	mul := func(acc int, cur int) int {
		return acc * cur
	}
	sum := func(acc int, cur int) int {
		return acc + cur
	}
	div := func(acc int, cur int) int {
		return acc / cur
	}
	as := []int{500, 2}
	ReduceInt(as, mul)
	ReduceInt(as, sum)
	ReduceInt(as, div)
}

func ReduceInt(a []int, f func(int, int) int) {
	nw := a[0]
	for i := 1; i < len(a); i++ {
		nw = f(nw, a[i])
	}
	n := strconv.Itoa(nw)
	for _, w := range n {
		z01.PrintRune(w)
	}
	z01.PrintRune('\n')
}
