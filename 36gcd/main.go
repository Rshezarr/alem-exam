package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	arg1, err1 := strconv.Atoi(os.Args[1])
	arg2, err2 := strconv.Atoi(os.Args[2])
	if err1 != nil || err2 != nil {
		return
	}
	gcd(arg1, arg2)
}

func gcd(num1 int, num2 int) {
	var gcdnum int

	for i := 1; i <= num1 && i <= num2; i++ {
		if num1%i == 0 && num2%i == 0 {
			gcdnum = i
		}
	}
	fmt.Println(gcdnum)
}
