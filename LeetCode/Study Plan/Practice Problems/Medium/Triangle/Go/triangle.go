package main

import (
	"fmt"
	"math"
)

//-------------------------recursion----------------------------
func min_path_sum(triangle [][]int) int {
	row, col := 0, 0
	return _min_path_sum(triangle, row, col)
}

func _min_path_sum(grid [][]int, i, j int) int {

	if i == len(grid)-1 {
		return grid[i][j]
	}

	down := grid[i][j] + _min_path_sum(grid, i+1, j)
	diagonal := grid[i][j] + _min_path_sum(grid, i+1, j+1)
	return min(down, diagonal)
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

//-------------------------recursion----------------------------

//-------------------------memoization----------------------------
func min_path_sum_memo(triangle [][]int) int {
	row := len(triangle)
	dp := make([][]int, row)
	for i := 0; i < row; i++ {
		dp[i] = make([]int, len(triangle[i]))
		for j := 0; j < len(triangle[i]); j++ {
			dp[i][j] = math.MinInt64
		}
	}
	return memo(triangle, 0, 0, &dp)
}

func memo(grid [][]int, i, j int, dp *[][]int) int {
	if i == len(grid)-1 {
		return grid[i][j]
	}

	if (*dp)[i][j] != math.MinInt64 {
		return (*dp)[i][j]
	}

	down := grid[i][j] + memo(grid, i+1, j, dp)
	diag := grid[i][j] + memo(grid, i+1, j+1, dp)
	(*dp)[i][j] = min(down, diag)
	return (*dp)[i][j]
}

//-------------------------memoization---------------------------

//-------------------------dp bottom up---------------------------

func min_path_sum_dp(grid [][]int) int {
	n := len(grid)
	dp := make([][]int, n)

	for i := 0; i < n; i++ {
		dp[i] = make([]int, len(grid[i]))
	}

	for i := 0; i < n; i++ {
		dp[n-1][i] = grid[n-1][i]
	}

	for i := n - 2; i >= 0; i-- {
		for j := i; j >= 0; j-- {
			down := grid[i][j] + dp[i+1][j]
			diag := grid[i][j] + dp[i+1][j+1]
			dp[i][j] = min(down, diag)
		}
	}
	return dp[0][0]
}

//-------------------------dp bottom up---------------------------

func main() {
	fmt.Println(min_path_sum_dp([][]int{
		{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3},
	}))
}
