package main

import (
	"fmt"
)

//-------------------------Recursion TLE-------------------------
func minimum_path_sum(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	return _minimum_path_sum(grid, m-1, n-1)
}

func _minimum_path_sum(arr [][]int, m, n int) int {
	if m == 0 && n == 0 {
		return arr[m][n]
	}

	if m < 0 || n < 0 {
		return 1e9
	}

	up := arr[m][n] + _minimum_path_sum(arr, m-1, n)
	left := arr[m][n] + _minimum_path_sum(arr, m, n-1)
	return min(up, left)
}

//------------------------------------------------------

//--------------------------Minimum path sum memoization----------------------------
func minimum_path_memo(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			dp[i][j] = -1
		}
	}

	return _minimum_path_memo(grid, &dp, m-1, n-1)

}

func _minimum_path_memo(grid [][]int, dp *[][]int, m, n int) int {
	if m == 0 && n == 0 {
		return grid[m][n]
	}

	if m < 0 || n < 0 {
		return 1e9
	}

	if (*dp)[m][n] != -1 {
		return (*dp)[m][n]
	}

	up := grid[m][n] + _minimum_path_memo(grid, dp, m-1, n)
	left := grid[m][n] + _minimum_path_memo(grid, dp, m, n-1)
	(*dp)[m][n] = min(left, up)
	return (*dp)[m][n]
}

//--------------------------Minimum path sum memoization ends----------------------------

//--------------------------Minimum path sum bottom up----------------------------
func minimum_path_dp(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			dp[i][j] = -1
		}
	}

	dp[0][0] = grid[0][0]
	for i := 1; i < m; i++ {
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}

	for i := 1; i < n; i++ {
		dp[0][i] = dp[0][i-1] + grid[0][i]
	}

	for i := 1; i < m; i++ {
		for j := 1; j < m; j++ {
			dp[i][j] = grid[i][j] + min(dp[i-1][j], dp[i][j-1])
		}
	}

	return dp[m-1][n-1]
}

//--------------------------Minimum path sum bottom up----------------------------

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	fmt.Println(minimum_path_dp([][]int{
		{1, 3, 1},
		{1, 5, 1},
		{4, 2, 1},
	}))
}
