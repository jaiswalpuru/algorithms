/*
You are given a m x n matrix grid. Initially, you are located at the top-left corner (0, 0), and in each step,
you can only move right or down in the matrix.

Among all possible paths starting from the top-left corner (0, 0) and ending in the bottom-right corner (m - 1, n - 1),
find the path with the maximum non-negative product. The product of a path is the product of all integers in the grid
cells visited along the path.

Return the maximum non-negative product modulo 109 + 7. If the maximum product is negative, return -1.

Notice that the modulo is performed after getting the maximum product.
*/

package main

import (
	"fmt"
)

const (
	mod = 1000000007
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func non_neg_prod_max(arr [][]int) int {
	n := len(arr)
	m := len(arr[0])

	max_dp := make([][]int, n)
	min_dp := make([][]int, n)

	for i := 0; i < n; i++ {
		max_dp[i] = make([]int, m)
		min_dp[i] = make([]int, m)
	}

	max_dp[0][0] = arr[0][0]
	min_dp[0][0] = arr[0][0]

	for i := 1; i < n; i++ {
		max_dp[i][0] = max_dp[i-1][0] * arr[i][0]
		min_dp[i][0] = min_dp[i-1][0] * arr[i][0]
	}

	for i := 1; i < m; i++ {
		max_dp[0][i] = max_dp[0][i-1] * arr[0][i]
		min_dp[0][i] = min_dp[0][i-1] * arr[0][i]
	}

	//check if zero is present in the array
	zero_flag := false
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if arr[i][j] == 0 {
				zero_flag = true
				break
			}
		}
	}

	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			max_dp[i][j] = max(0, arr[i][j]*max_dp[i-1][j])
			max_dp[i][j] = max(max_dp[i][j], arr[i][j]*max_dp[i][j-1])
			max_dp[i][j] = max(max_dp[i][j], arr[i][j]*min_dp[i-1][j])
			max_dp[i][j] = max(max_dp[i][j], arr[i][j]*min_dp[i][j-1])

			min_dp[i][j] = min(0, arr[i][j]*max_dp[i-1][j])
			min_dp[i][j] = min(min_dp[i][j], arr[i][j]*max_dp[i][j-1])
			min_dp[i][j] = min(min_dp[i][j], arr[i][j]*min_dp[i-1][j])
			min_dp[i][j] = min(min_dp[i][j], arr[i][j]*min_dp[i][j-1])
		}
	}

	res := max(max_dp[n-1][m-1], min_dp[n-1][m-1])
	if res < 0 || (res == 0 && !zero_flag) {
		return -1
	}

	return res % mod
}

func main() {
	fmt.Println(non_neg_prod_max([][]int{
		{1, -2, 1},
		{1, -2, 1},
		{3, -4, 1},
	}))
	fmt.Println(non_neg_prod_max([][]int{
		{-1, -2, -3},
		{-2, -3, -3},
		{-3, -3, -2},
	}))
}
