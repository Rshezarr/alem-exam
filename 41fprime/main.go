package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) == 2 {
		num, err := strconv.Atoi(os.Args[1])
		if err != nil {
			return
		}
		Fprime(num)
	}
	z01.PrintRune('\n')
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
