package main

import (
	"fmt"
	"math"
)

func lis(arr []int) int {
	n := len(arr)
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = -1
	}
	return lis_dp_memo(arr, &dp, n-1)
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
func lis_dp_memo(arr []int, dp *[]int, i int) int {
	if (*dp)[i] != -1 {
		return (*dp)[i]
	}

	if i == 0 {
		(*dp)[i] = 1
		return (*dp)[i]
	}

	longest := 1
	for j := 0; j < i; j++ {
		if arr[i] > arr[j] {
			longest = max(longest, 1+lis_dp_memo(arr, dp, j))
		}
	}
	(*dp)[i] = longest
	return (*dp)[i]
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
	fmt.Println(lis([]int{10, 9, 2, 5, 3, 7, 101, 18}))
}
