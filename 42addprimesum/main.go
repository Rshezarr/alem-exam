package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	n := Atoi(os.Args[1])
	var sum int
	for i := n; i >= 2; i-- {
		if IsPrime(i) {
			sum += i
		}
		continue
	}
	Print(Itoa(sum))
}

func IsPrime(n int) bool {
	if n <= 1 {
		return false
	}

	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func Print(s string) {
	for _, w := range s {
		z01.PrintRune(w)
	}
	z01.PrintRune('\n')
}

func Atoi(s string) int {
	var bl bool
	var num int

	if len(s) > 0 {
		if s[0] == '-' {
			bl = true
			s = s[1:]
		}

		for _, w := range s {
			if w < 48 || w > 57 {
				return 0
			}
			num = num*10 + int(w) - 48
		}

		if bl {
			num *= -1
		}
	}

	return num
}

func Itoa(n int) string {
	var ch string
	if n < 0 {
		n = -n
		ch = "-"
	}
	digits := "0123456789"
	if n < 10 {
		return ch + digits[n:n+1]
	}
	return ch + Itoa(n/10) + digits[n%10:n%10+1]
}
