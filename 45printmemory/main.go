package main

import (
	"unicode"

	"github.com/01-edu/z01"
)

func main() {
	PrintMemory([10]byte{'h', 'e', 'l', 'l', 'o', 16, 21, '*'})
}

func PrintMemory(arr [10]byte) {
	var firs_res string
	for _, v := range arr {
		firs_res += PrintHex(int(v)) + " "
	}
	Println(firs_res[:11])
	Println(firs_res[12:23])
	Println(firs_res[24:29])
	var sec_res string
	for _, v := range arr {
		if unicode.IsGraphic(rune(v)) {
			sec_res += string(v)
		} else {
			sec_res += "."
		}
	}
	Println(sec_res)
}

func PrintHex(n int) string {
	var res string
	base := "0123456789abcdef"
	if n == 0 {
		return "00"
	}
	for n != 0 {
		res = string(base[n%16]) + res
		n /= 16
	}
	return res
}

func Println(s string) {
	for _, w := range s {
		z01.PrintRune(w)
	}
	z01.PrintRune('\n')
}
