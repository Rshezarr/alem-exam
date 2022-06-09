package main

import (
	"unicode"

	"github.com/01-edu/z01"
)

func main() {
	PrintMemory([10]byte{'h', 'e', 'l', 'l', 'o', 16, 21, '*'})
}

func PrintMemory(arr [10]byte) {
	var fres string
	for i := 0; i < len(arr); i++ {
		fres += PrintHex(int(arr[i])) + " "
	}
	Print(fres[:11])
	Print(fres[12:23])
	Print(fres[24:29])
	var sres string
	for _, v := range arr {
		if unicode.IsGraphic(rune(v)) {
			sres += string(v)
		} else {
			sres += "."
		}
	}
	Print(sres)
}

func PrintHex(n int) string {
	if n == 0 {
		return "00"
	}
	hex := "0123456789abcdef"
	var res string
	for n != 0 {
		res = string(hex[n%16]) + res
		n /= 16
	}

	if len(res) == 1 {
		return "0" + res
	}
	return res
}

func Print(s string) {
	for _, v := range s {
		z01.PrintRune(v)
	}
	z01.PrintRune('\n')
}
