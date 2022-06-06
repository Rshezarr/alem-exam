package main

import "fmt"

func main() {
	fmt.Println(ItoaBase(2, 2))
}

func ItoaBase(value, base int) string {
	if value == 0 {
		return "0"
	}

	var res string

	if value < 0 {
		value = -value
		res = "-"
	}

	b := "0123456789ABCDEF"
	b = b[:base]
	l := len(b)

	for i := 1; i < l; i++ {
		if value == 0 {
			break
		}
		res = string(b[value%l]) + res
		value = value / l
	}
	return res
}
