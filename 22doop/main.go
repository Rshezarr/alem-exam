package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 4 {
		return
	}

	a := Atoi(os.Args[1])
	b := Atoi(os.Args[3])
	op := os.Args[2]

	// maxInt := int(^uint(0) >> 1)
	// minInt := -maxInt - 1

	if a > 0 && b > 0 && a+b < 0 {
		os.Exit(0)
	}
	switch op {
	case "+":
		c := a + b
		fmt.Printf("%d\n", c)
	case "-":
		c := a - b
		fmt.Printf("%d\n", c)
	case "*":
		c := a * b
		fmt.Printf("%d\n", c)
	case "/":
		if b == 0 {
			fmt.Printf("No division by 0\n")
			return
		}
		c := a / b
		fmt.Printf("%d\n", c)
	case "%":
		if b == 0 {
			fmt.Printf("No modulo by 0\n")
			return
		}
		c := a % b
		fmt.Printf("%d\n", c)
	}
}

func Atoi(s string) int {
	var num int
	var nega bool

	if len(s) > 0 {
		if s[0] == '-' {
			nega = true
			s = s[1:]
		}

		for _, w := range s {
			if w < 48 || w > 58 {
				os.Exit(0)
			}
			num = num*10 + int(w) - 48
		}

		if nega {
			num *= -1
		}
	}
	return num
}
