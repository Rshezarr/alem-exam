package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	num, err := strconv.Atoi(os.Args[1])
	if err != nil || num <= 1 {
		return
	}
	Fprime(num)
	fmt.Println()
}

func Fprime(n int) {
	for i := 2; i < n*2; i++ {
		if n%i == 0 {
			fmt.Print(i)
			if n/i != 1 {
				fmt.Print("*")
			}
			Fprime(n / i)
			break
		}
	}
}
