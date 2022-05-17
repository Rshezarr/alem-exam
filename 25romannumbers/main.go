package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 2 {
		arg := os.Args[1]
		num := BasicAtoi(arg)
		fmt.Printf("%s \n", ToRoman(num))
	}
}

func ToRoman(num int) string {
	romans := ""

	numbers := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	symbols := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	mapRoman := make(map[int]string)
	i := 0
	z := 0
	for num > 0 {
		k := num / numbers[i]
		for j := 0; j < k; j++ {
			mapRoman[z] = symbols[i]
			romans += symbols[i]
			num -= numbers[i]
			z++
		}
		i++
	}

	calcRoman := ""
	for j := 0; j < z; j++ {
		if len(mapRoman[j]) == 1 {
			calcRoman += mapRoman[j]
		} else {
			calcRoman += "(" + string(mapRoman[j][1]) + "-" + string(mapRoman[j][0]) + ")"
		}
		if j != z-1 {
			calcRoman += "+"
		}
	}
	return calcRoman + "\n" + romans
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
