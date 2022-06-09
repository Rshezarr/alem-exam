package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		return
	}

	code := os.Args[1]

	var arr [2048]byte
	var bracks []int
	pos := -1
	var pow int

	for i, r := range code {
		if r == '[' {
			bracks = append(bracks, i)
		}
	}

	for i := 0; i < len(code); i++ {
		switch code[i] {
		case '+':
			arr[pow]++
		case '-':
			arr[pow]--
		case '>':
			pow++
		case '<':
			pow--
		case '.':
			z01.PrintRune(rune(arr[pow]))
		case ']':
			if arr[pow] == 0 {
				pos--
			} else {
				for j := i; i >= 0; j-- {
					if code[j] == '[' && bracks[pos] == j {
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
