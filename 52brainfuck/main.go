package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		return
	}

	arg := os.Args[1]

	var arr [2048]byte
	var pow int
	var brack []int
	pos := -1

	for i, v := range arg {
		if v == '[' {
			brack = append(brack, i)
		}
	}

	for i := 0; i < len(arg); i++ {
		switch arg[i] {
		case '+':
			arr[pow]++
		case '-':
			arr[pow]--
		case '<':
			pow--
		case '>':
			pow++
		case '.':
			z01.PrintRune(rune(arr[pow]))
		case ']':
			if arr[pow] == 0 {
				pos--
			} else {
				for j := i; j >= 0; j-- {
					if arg[j] == '[' && j == brack[pos] {
						i = j
						break
					}
				}
			}
		case '[':
			pos++
		}
	}
}
