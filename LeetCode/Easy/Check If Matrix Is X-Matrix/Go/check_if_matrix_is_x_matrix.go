package main

import "fmt"

func check_if_matrix_is_x_matrix(arr [][]int) bool {
	n := len(arr)
	for i := 0; i < n; i++ {
		if arr[i][i] == 0 || arr[i][n-1-i] == 0 {
			return false
		} else {
			arr[i][i] = 0
			arr[i][n-1-i] = 0
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if arr[i][j] != 0 {
				return false
			}
		}
	}
	return true
}

func main() {
	fmt.Println(check_if_matrix_is_x_matrix([][]int{
		{2, 0, 0, 1}, {0, 3, 1, 0}, {0, 5, 2, 0}, {4, 0, 0, 2},
	}))
}
