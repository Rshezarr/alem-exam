package main

import (
	"unicode"

	"github.com/01-edu/z01"
)

func main() {
	PrintMemory([10]byte{'h', 'e', 'l', 'l', 'o', 16, 21, '*'})
}

func PrintMemory(arr [10]byte) {
	var temp string
	for _, v := range arr {
		temp += string(PrintHex(int(v))) + " "
	}

	Println(temp[:11])
	Println(temp[12:23])
	Println(temp[24:29])
	var str string
	for _, v := range arr {
		if unicode.IsGraphic(rune(v)) {
			str += string(v)
		} else {
			str += "."
		}
	}
	Println(str)
}

func PrintHex(n int) string {
	var res string
	if n == 0 {
		return "00"
	}
	base := "0123456789abcdef"
	for n != 0 {
		res = string(base[n%16]) + res
		n /= 16
	}
	return res
}

func Println(s string) {
	for _, ch := range s {
		z01.PrintRune(ch)
	}
	z01.PrintRune('\n')
}
