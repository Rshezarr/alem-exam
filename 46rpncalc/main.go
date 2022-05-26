package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Print("Error\n")
		return
	}
	arg := os.Args[1]
	input := strings.Split(arg, " ")
	var arr []int
	var newinput []string
	for _, w := range input {
		if w != "" {
			newinput = append(newinput, w)
		}
	}
	for _, w := range newinput {
		num, err := strconv.Atoi(w)
		if err != nil {
			if len(arr) < 2 {
				fmt.Print("Error\n")
				return
			}
			switch w {
			case "+":
				arr[len(arr)-2] += arr[len(arr)-1]
				arr = arr[:len(arr)-1]
			case "-":
				arr[len(arr)-2] -= arr[len(arr)-1]
				arr = arr[:len(arr)-1]
			case "/":
				arr[len(arr)-2] /= arr[len(arr)-1]
				arr = arr[:len(arr)-1]
			case "%":
				arr[len(arr)-2] %= arr[len(arr)-1]
				arr = arr[:len(arr)-1]
			case "*":
				arr[len(arr)-2] *= arr[len(arr)-1]
				arr = arr[:len(arr)-1]
			case "":
				continue
			default:
				fmt.Println("Error")
				return
			}
		} else {
			arr = append(arr, num)
		}
	}
	if len(arr) != 1 {
		fmt.Println("Error")
		return
	}
	fmt.Println(arr[0])
}
