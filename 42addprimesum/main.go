package main

import (
	"fmt"
	"os"
)

func main() {
	primes := SieveOfEratosthenes(Atoi(os.Args[1]))

	var sum int
	for _, prime := range primes {
		sum1 := sum + prime
		if sum1 < sum {
			fmt.Println(sum)
		} else {
			sum = sum1
		}
	}
	fmt.Println(sum)

}

func SieveOfEratosthenes(n int) []int {
	integers := make([]bool, n+1)
	for i := 2; i < n+1; i++ {
		integers[i] = true
	}

	for p := 2; p*p <= n; p++ {
		if integers[p] {
			for i := p * 2; i <= n; i += p {
				integers[i] = false
			}
		}
	}

	var primes []int
	for p := 2; p <= n; p++ {
		if integers[p] {
			primes = append(primes, p)
		}
	}

	return primes
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
