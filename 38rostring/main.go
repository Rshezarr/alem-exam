package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) == 2 {
		arg := os.Args[1]
		res := SplitAndJoin(arg)
		Print(res)
	}
}

func SplitAndJoin(s string) string {
	var temp string
	var arr []string
	var res string

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

	for i, v := range arr[1:] {
		if i == len(arr)-2 {
			res += v + " " + arr[0]
		} else {
			res += v + " "
		}
	}
	return res
}

func Print(s string) {
	for _, v := range s {
		z01.PrintRune(v)
	}
	z01.PrintRune('\n')
}
