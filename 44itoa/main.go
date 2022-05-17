package main

import "fmt"

func main() {
	fmt.Println(Itoa(125))
	fmt.Println(Itoa(-125))
}

func Itoa(n int) string {
	var s string

	if n < 0 {
		n = -n
		s = "-"
	}

	digits := "0123456789"

	if n < 10 {
		return s + digits[n:n+1]
	}
	return s + Itoa(n/10) + digits[n%10:n%10+1]
}
