package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Error")
		return
	}
	arg := os.Args[1]

	input := strings.Split(arg, " ")

	var num_arr []int
	l := len(num_arr)
	var str_arr []string

	for _, w := range input {
		if w != "" {
			str_arr = append(str_arr, w)
		}
	}

	for _, w := range str_arr {
		num, err := strconv.Atoi(w)

		if err != nil {
			if l < 2 {
				fmt.Println("Error")
				return
			}
			l := l
			switch w {
			case "+":
				num_arr[l-2] += num_arr[l-1]
				num_arr = num_arr[:l-1]
			case "-":
				num_arr[l-2] -= num_arr[l-1]
				num_arr = num_arr[:l-1]
			case "*":
				num_arr[l-2] *= num_arr[l-1]
				num_arr = num_arr[:l-1]
			case "/":
				num_arr[l-2] /= num_arr[l-1]
				num_arr = num_arr[:l-1]
			case "%":
				num_arr[l-2] %= num_arr[l-1]
				num_arr = num_arr[:l-1]
			case "":
				continue
			default:
				fmt.Println("Error")
				return
			}
		} else {
			num_arr = append(num_arr, num)
		}
	}

	if l != 1 {
		fmt.Println("Error")
		return
	}
	fmt.Println(num_arr[0])
}
