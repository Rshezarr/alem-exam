package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(ItoaBase(125, 16))
	fmt.Println(ItoaBase(math.MaxInt, 16))
	fmt.Println(ItoaBase(math.MinInt, 16))
}

func ItoaBase(value, base int) string {
	if value == 0 {
		return "0"
	}
	var m string
	var fl bool
	if value < 0 {
		fl = true
		m = "-"
	}

	hex := "0123456789ABCDEF"
	hex = hex[:base]
	var res string
	for value != 0 {
		if fl {
			res = string(hex[(value%base)*-1]) + res
		} else {
			res = string(hex[(value%base)]) + res
		}
		value /= base
	}
	return m + res
}
