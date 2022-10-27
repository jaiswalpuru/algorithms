package main

import "fmt"

func check_if_every_row_and_col_contains_all_numbers(arr [][]int) bool {
	n := len(arr)
	for i := 0; i < n; i++ {
		r, c := make(map[int]bool), make(map[int]bool)
		for j := 0; j < n; j++ {
			r[arr[i][j]] = true
			c[arr[j][i]] = true
		}
		for j := 1; j <= n; j++ {
			if !r[j] || !c[j] {
				return false
			}
		}
	}
	return true
}

func main() {
	fmt.Println(check_if_every_row_and_col_contains_all_numbers([][]int{
		{1, 2, 3}, {3, 1, 2}, {2, 3, 1},
	}))
}
