package main

import (
	"fmt"
	"math"
)

func lis(arr []int) int {
	n := len(arr)
	memo := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		memo[i] = make([]int, n+1)
		for j := 0; j <= n; j++ {
			memo[i][j] = -1
		}
	}
	return _memo(arr, 0, -1, n, &memo)
}

//recursive solution
func _lis(arr []int, start int, prev int) int {
	if start == len(arr) {
		return 0
	}

	exclude := _lis(arr, start+1, prev)

	include := 0
	if arr[start] > prev {
		include = _lis(arr, start+1, arr[start]) + 1
	}

	return max(exclude, include)
}

//dp with memoization
func _memo(arr []int, ind, prev, n int, memo *[][]int) int {
	if ind == n {
		return 0
	}
	if (*memo)[ind][prev+1] != -1 {
		return (*memo)[ind][prev+1]
	}
	exclude := _memo(arr, ind+1, prev, n, memo)
	include := 0
	if prev == -1 || arr[ind] > arr[prev] {
		include = max(include, 1+_memo(arr, ind+1, ind, n, memo))
	}
	(*memo)[ind][prev+1] = max(include, exclude)
	return (*memo)[ind][prev+1]
}

//pure dp solution
func lis_dp(arr []int) int {
	n := len(arr)
	dp := make([]int, n)

	res := math.MinInt64
	for i := 0; i < n; i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if arr[i] > arr[j] {
				dp[i] = max(dp[i], 1+dp[j])
			}
		}
		res = max(res, dp[i])
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(lis([]int{1, 3, 6, 7, 9, 4, 10, 5, 6}))
}
