package main

import (
	"fmt"
	"os"
)

func BasicAtoi(s string) int {
	var num int
	for _, w := range s {
		num = num*10 + int(w) - 48
	}
	return num
}

func main() {
	for _, w := range os.Args[1] {
		if w < 48 || w > 57 {
			fmt.Println("ERROR: cannot convert to roman digit")
			return
		}
	}

	arg := BasicAtoi(os.Args[1])

	if arg > 3999 {
		fmt.Println("ERROR: cannot convert to roman digit")
		return
	}

	fmt.Println(RomanNumbers(arg))

}

func RomanNumbers(num int) string {
	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	symbols := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	var roman string
	var i int

	for num > 0 {
		k := num / values[i]
		for j := 0; j < k; j++ {
			roman += symbols[i]
			num -= values[i]
		}
		i++
	}
	return roman
}
