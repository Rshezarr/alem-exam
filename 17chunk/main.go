package main

import "fmt"

func main() {
	Chunk([]int{}, 10)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 0)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 3)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 5)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 4)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 2)
}

func Chunk(slice []int, size int) {
	if slice == nil {
		return
	}

	if size == 0 {
		fmt.Println()
		return
	}

	temp := []int{}
	res := [][]int{}
	var j int

	for i := 0; i < len(slice); i++ {
		temp = append(temp, slice[i])
		j++
		if j == size {
			res = append(res, temp)
			j = 0
			temp = []int{}
		}
	}

	if j != 0 {
		res = append(res, temp)
	}

	fmt.Println(res)
}
