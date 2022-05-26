package main

import "fmt"

func main() {
	fmt.Println(AtoiBase("125", "0123456789"))
	fmt.Println(AtoiBase("1111101", "01"))
	fmt.Println(AtoiBase("7D", "0123456789ABCDEF"))
	fmt.Println(AtoiBase("uoi", "choumi"))
	fmt.Println(AtoiBase("bbbbbab", "-ab"))
}

func AtoiBase(s string, base string) int {
	mapa := make(map[rune]int)

	for i, v := range base {
		if _, ok := mapa[v]; ok || (v == '+' || v == '-') {
			return 0
		}
		mapa[v] = i
	}
	var res int
	for i := 0; i < len(s); i++ {
		res = res*len(base) + mapa[rune(s[i])]
	}
	return res
}
