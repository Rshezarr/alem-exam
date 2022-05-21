package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 4 {
		return
	}
	ErrorCheck(os.Args[1], os.Args[3])

	a := Atoi(os.Args[1])
	b := Atoi(os.Args[3])
	op := os.Args[2]

	maxInt := int(^uint(0) >> 1)
	minInt := -maxInt - 1

	if op == "/" && b == 0 {
		fmt.Println("No division by 0")
	} else if op == "%" && b == 0 {
		fmt.Println("No modulo by 0")
	} else {
		switch op {
		case "+":
			// Overflow(a, b, maxInt, minInt, "add")
			c := a + b
			if c < 0 && a <= maxInt || a > minInt {
				return
			}
			fmt.Printf("%d\n", c)
		case "-":
			//	Overflow(a, b, maxInt, minInt, "sub")
			c := a - b
			fmt.Printf("%d\n", c)
		case "*":
			//	Overflow(a, b, maxInt, minInt, "mul")
			c := a * b
			fmt.Printf("%d\n", c)
		case "/":
			c := a / b
			fmt.Printf("%d\n", c)
		case "%":
			c := a % b
			fmt.Printf("%d\n", c)
		}
		return
	}
}

func ErrorCheck(num1, num2 string) {
	for _, w := range num1 + num2 {
		if w != 45 && (w < 48 || w > 57) {
			os.Exit(0)
		}
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
				return 0
			}
			num = num*10 + int(w) - 48
		}

		if nega {
			num *= -1
		}
	}
	return num
}

// func Overflow(a, b, maxInt, minInt int, op string) {
// 	if op == "add" {
// 		if len(os.Args[1]) >= 19 {
// 		}
// 	}

// 	if op == "sub" {
// 	}

// 	if op == "mul" {
// 	}
// }
