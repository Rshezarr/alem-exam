package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	arg := BasicAtoi(os.Args[1])
	if arg >= 4000 || arg < 1 {
		fmt.Printf("ERROR: cannot convert to roman digit\n")
		return
	}

	num_array := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	sym_array := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	symCalc_arrray := []string{"M", "(M-C)", "D", "(D-C)", "C", "(C-X)", "L", "(L-X)", "X", "(X-I)", "V", "(V-I)", "I"}
	var roman string
	var calcRoman string
	var i int

	for arg > 0 {
		k := arg / num_array[i]
		for j := 0; j < k; j++ {
			roman += sym_array[i]
			calcRoman += symCalc_arrray[i] + "+"
			arg -= num_array[i]
		}
		i++
	}
	fmt.Printf("%s\n%s\n", calcRoman[:len(calcRoman)-1], roman)
}

func BasicAtoi(s string) int {
	var num int

	for _, w := range s {
		if w < 48 || w > 57 {
			os.Exit(0)
		}
		num = num*10 + int(w) - 48
	}
	return num
}
