package main

import (
	"fmt"
	"math"
)

//----------------------------brute force----------------------------
func minimum_falling_path_sum(arr [][]int) int {
	n := len(arr)
	val := math.MaxInt64
	for i := n - 1; i >= 0; i-- {
		val = min(val, dfs(arr, n-1, i))
	}
	return val
}

func dfs(arr [][]int, i, j int) int {
	if i >= len(arr) || i < 0 || j < 0 || j >= len(arr) {
		return math.MaxInt64
	}

	if i == 0 && j < len(arr) {
		return arr[i][j]
	}

	up := dfs(arr, i-1, j)
	up_left := dfs(arr, i-1, j-1)  //diagnally left in upwards direction
	up_right := dfs(arr, i-1, j+1) //diagonally right in upwards direction
	return arr[i][j] + min(up, min(up_left, up_right))
}

//-------------------------------------------------------------------

//------------------------------memoization-------------------------------------
func minimum_falling_path_sum_memo(arr [][]int) int {
	n := len(arr)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			dp[i][j] = -1
		}
	}

	val := math.MaxInt64
	for i := n - 1; i >= 0; i-- {
		val = min(val, _minimum_falling_path_sum_memo(arr, n-1, i, &dp))
	}
	return val
}

func _minimum_falling_path_sum_memo(arr [][]int, i, j int, dp *[][]int) int {
	if i >= len(arr) || i < 0 || j < 0 || j >= len(arr) {
		return math.MaxInt64
	}

	if (*dp)[i][j] != -1 {
		return (*dp)[i][j]
	}

	if i == 0 && j < len(arr) {
		return arr[i][j]
	}

	up := _minimum_falling_path_sum_memo(arr, i-1, j, dp)
	up_left := _minimum_falling_path_sum_memo(arr, i-1, j-1, dp)
	up_right := _minimum_falling_path_sum_memo(arr, i-1, j+1, dp)
	(*dp)[i][j] = arr[i][j] + min(up, min(up_left, up_right))
	return (*dp)[i][j]
}

//-------------------------------------------------------------------------------

//--------------------------------------dp-----------------------------------------
func min_fall_path_sum_dp(arr [][]int) int {
	n := len(arr)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		dp[0][i] = arr[0][i]
	}

	var up, up_left, up_right int
	for i := 1; i < n; i++ {
		for j := 0; j < n; j++ {
			up = arr[i][j] + dp[i-1][j]
			up_left = arr[i][j]
			if j-1 >= 0 {
				up_left += dp[i-1][j-1]
			} else {
				up_left += 1e8
			}
			up_right = arr[i][j]
			if j+1 < n {
				up_right += dp[i-1][j+1]
			} else {
				up_right += 1e8
			}
			dp[i][j] = min(up, min(up_left, up_right))
		}
	}

	min_val := math.MaxInt64
	for i := 0; i < n; i++ {
		min_val = min(min_val, dp[n-1][i])
	}
	return min_val
}

//---------------------------------------------------------------------------------

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	fmt.Println(min_fall_path_sum_dp([][]int{
		{2, 1, 3},
		{6, 5, 4},
		{7, 8, 9},
	}))
}
