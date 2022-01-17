package main

import "fmt"

func max_square(arr [][]byte) int {
	n := len(arr)
	m := len(arr[0])
	dp := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = make([]int, m+1)
	}
	max_val := 0
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if arr[i-1][j-1] == '1' {
				dp[i][j] = 1 + min(dp[i-1][j-1], min(dp[i-1][j], dp[i][j-1]))
				max_val = max(max_val, dp[i][j])
			}

		}
	}
	return max_val * max_val
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(max_square(
		[][]byte{{'1', '0', '1', '0', '0'}, {'1', '0', '1', '1', '1'}, {'1', '1', '1', '1', '1'}, {'1', '0', '0', '1', '0'}},
	))
}
