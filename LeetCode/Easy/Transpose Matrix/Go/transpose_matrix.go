package main

import "fmt"

func transpose_matrix(arr [][]int) [][]int {
	res := make([][]int, len(arr[0]))
	for i := 0; i < len(arr[0]); i++ {
		res[i] = make([]int, len(arr))
	}

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[0]); j++ {
			res[j][i] = arr[i][j]
		}
	}
	return res
}

func main() {
	fmt.Println(transpose_matrix([][]int{
		{1, 2, 3}, {4, 5, 6}, {7, 8, 9},
	}))
}
