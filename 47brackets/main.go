package main

import (
	"fmt"
	"os"
)

func main() {
	arg := os.Args[1:]
	if len(os.Args) > 1 {
		for _, v := range arg {
			fmt.Println(Brackets(v))
		}
	} else {
		fmt.Println()
	}
}

func Brackets(s string) string {
	var b string
	for _, v := range s {
		if (v == ')' || v == ']' || v == '}') && len(b) == 0 {
			return "Error"
		}
		if v == '(' || v == '[' || v == '{' {
			b += string(v)
		}
		if v == ')' && b[len(b)-1] == '(' {
			b = b[:len(b)-1]
		}
		if v == ']' && b[len(b)-1] == '[' {
			b = b[:len(b)-1]
		}
		if v == '}' && b[len(b)-1] == '{' {
			b = b[:len(b)-1]
		}
	}
	if len(b) == 0 {
		return "OK"
	}
	return "Error"
}
