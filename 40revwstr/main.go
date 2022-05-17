package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) == 2 {
		Print(SplitAndJoin(os.Args[1]))
	}
}

func SplitAndJoin(s string) string {
	var str string
	var arr []string
	var res string
	for i, w := range s {
		if i == len(s)-1 && w != ' ' {
			str += string(w)
			arr = append(arr, str)
		} else if w != ' ' {
			str += string(w)
		} else {
			if len(str) >= 1 {
				arr = append(arr, str)
			}
			str = ""
		}
	}

	for i, w := range arr {
		if i != len(arr)-1 {
			res = " " + w + res
		} else {
			res = w + res
		}
	}
	return res
}

func Print(s string) {
	for _, w := range s {
		z01.PrintRune(w)
	}
	z01.PrintRune('\n')
}
