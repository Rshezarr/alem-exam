package main

import "fmt"

func main() {
	fmt.Println(Split("HelloHAareHAyou", "HA"))
}

func Split(s, sep string) []string {
	var arr []string
	var str string

	for i := 0; i < len(s)-len(sep); i++ {
		if s[i:i+len(sep)] == sep {
			arr = append(arr, str)
			str = ""
			i = i + len(sep) - 1
		} else {
			str += string(s[i])
		}
	}
	str += s[len(s)-len(sep):]
	arr = append(arr, str)

	return arr
}
