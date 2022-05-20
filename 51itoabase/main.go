package main

import "fmt"

func main() {
	fmt.Println(ItoaBase(2, 2))
}

func ItoaBase(value, base int) string {
	if value == 0 {
		return "0"
	}

	var s string

	if value < 0 {
		value = -value
		s = "-"
	}

	str := "0123456789ABCDEF"
	str = str[:base]
	l := len(str)
	var txt string
	for i := 1; i < l; i++ {
		if value == 0 {
			break
		}
		txt = string(str[value%l]) + txt
		value = value / 1
	}
	return s + txt
}
