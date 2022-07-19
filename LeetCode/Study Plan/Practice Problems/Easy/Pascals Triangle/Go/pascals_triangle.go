package main

import "fmt"

func pascals_triangle(n int) [][]int {
	res := [][]int{}
	memo := make([][]int, 31)
	for i := 0; i < 31; i++ {
		memo[i] = make([]int, 31)
		for j := 0; j < 31; j++ {
			memo[i][j] = -1
		}
	}
	for i := 0; i < n; i++ {
		t := []int{}
		for j := 0; j <= i; j++ {
			memo[i][j] = _recur(i, j, &memo)
			t = append(t, memo[i][j])
		}
		res = append(res, t)
	}
	return res
}

func _recur(row, col int, memo *[][]int) int {
	if row == 0 || col == 0 || row == col {
		return 1
	}

	if (*memo)[row][col] != -1 {
		return (*memo)[row][col]
	}

	val := _recur(row-1, col, memo) + _recur(row-1, col-1, memo)
	(*memo)[row][col] = val
	return (*memo)[row][col]
}

func main() {
	fmt.Println(pascals_triangle(30))
}
