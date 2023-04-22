package main

import "fmt"

//------------------------------recursion brute force it might give TLE------------------------------
func unique_path_recursion(arr [][]int) int {
	m, n := len(arr), len(arr[0])
	return _recur(arr, m-1, n-1)
}

func _recur(arr [][]int, m, n int) int {
	if m == 0 && n == 0 {
		if arr[m][n] == 1 {
			return 0
		}
		return 1
	}

	if m < 0 || n < 0 {
		return 0
	}

	if arr[m][n] == 1 {
		return 0
	}

	left := _recur(arr, m, n-1)
	right := _recur(arr, m-1, n)
	return left + right
}

//----------------------------------------------------------------------------------------------------

//-----------------------------------memoization-----------------------------------
func unique_path_memo(arr [][]int) int {
	m, n := len(arr), len(arr[0])

	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			dp[i][j] = -1
		}
	}

	return memo(arr, &dp, m-1, n-1)
}

func memo(arr [][]int, dp *[][]int, m, n int) int {
	if m == 0 && n == 0 {
		if arr[m][n] == 1 {
			return 0
		}
		return 1
	}

	if m < 0 || n < 0 {
		return 0
	}

	if arr[m][n] == 1 {
		return 0
	}

	if (*dp)[m][n] != -1 {
		return (*dp)[m][n]
	}

	left := memo(arr, dp, m, n-1)
	right := memo(arr, dp, m-1, n)
	(*dp)[m][n] = left + right
	return (*dp)[m][n]
}

//---------------------------------------------------------------------------------

//-----------------------------------------dp--------------------------------------
func unique_path_dp(arr [][]int) int {
	m, n := len(arr), len(arr[0])

	dp := make([][]int, m)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		if arr[i][0] == 1 {
			break
		}
		dp[i][0] = 1
	}

	for i := 0; i < n; i++ {
		if arr[0][i] == 1 {
			break
		}
		dp[0][i] = 1
	}

	for i := 1; i < n; i++ {
		for j := 0; j < n; j++ {
			if arr[i][j] == 0 {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			} else {
				dp[i][j] = 0
			}
		}
	}
	return dp[m-1][n-1]
}

//---------------------------------------------------------------------------------

func main() {
	fmt.Println(unique_path_recursion([][]int{
		{0, 0, 0},
		{0, 1, 0},
		{0, 0, 0},
	}))
}
