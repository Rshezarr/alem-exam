package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) == 2 {
		arg := os.Args[1]
		var res string
		splt := Split(arg)
		for i, w := range splt {
			if i == len(splt)-1 {
				res += w
			} else {
				res += w + " "
			}
		}
		for _, w := range res {
			z01.PrintRune(w)
		}
		z01.PrintRune('\n')
	}
}

func Split(s string) []string {
	var temp string
	var arr []string

	for i, w := range s {
		if i == len(s)-1 && w != ' ' {
			temp += string(w)
			arr = append(arr, temp)
		} else if w != ' ' {
			temp += string(w)
		} else {
			if len(temp) >= 1 {
				arr = append(arr, temp)
			}
			temp = ""
		}
	}
	return arr
}
