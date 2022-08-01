package main

import (
	"fmt"
)

//--------------------------------recursion will give TLE----------------------------------------
func unique_paths_recursion(m, n int) int {
	return _recur(m-1, n-1)
}

func _recur(m, n int) int {
	return _recur(m-1, n-1)
}

func _recur(m, n int) int {
	if m == 0 && n == 0 {
		return 1
	}
	if m < 0 || n < 0 {
		return 0
	}
	down := _recur(m-1, n)
	right := _recur(m, n-1)
	return down + right
}

//----------------------------------------------------------------------------------------------------

//----------------------------------------memoization------------------------------------
func unique_paths_memo(m, n int) int {
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			dp[i][j] = -1
		}
	}

	return memo(&dp, m-1, n-1)
}

func memo(dp *[][]int, m, n int) int {
	if m < 0 || n < 0 {
		return 0
	}
	if m == 0 && n == 0 {
		return 1
	}

	if (*dp)[m][n] != -1 {
		return (*dp)[m][n]
	}

	left := memo(dp, m-1, n)
	right := memo(dp, m, n-1)
	(*dp)[m][n] = left + right
	return (*dp)[m][n]
}

//------------------------------------------------------------------------------------------

//----------------------------------using bottom up-----------------------------------------
func unique_paths_dp(m, n int) int {

	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	dp[0][0] = 1
	for i := 0; i < m; i++ {
		dp[i][0] = 1
	}

	for j := 0; j < n; j++ {
		dp[0][j] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m-1][n-1]
}

//------------------------------------------------------------------------------------------

func main() {
	fmt.Println(unique_paths_dp(3, 7))
}
