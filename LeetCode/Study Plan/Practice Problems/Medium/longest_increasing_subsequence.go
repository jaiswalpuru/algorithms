/*
Given an integer array nums, return the length of the longest strictly increasing subsequence.

A subsequence is a sequence that can be derived from an array by deleting some or no elements without changing
the order of the remaining elements. For example, [3,6,2,7] is a subsequence of the array [0,3,1,6,2,2,7].
*/

package main

import "fmt"

func lis(arr []int) int {
	if len(arr) == 1 {
		return 1
	}
	max := 0
	_lis(arr, len(arr), &max)
	return max
}

func _lis(arr []int, n int, max *int) int {
	if n == 1 {
		return 1
	}

	res, max_ending_here := 0, 1
	for i := 1; i < n; i++ {
		res = _lis(arr, i, max)
		if arr[i-1] < arr[n-1] && res+1 > max_ending_here {
			max_ending_here = res + 1
		}
	}

	if *max < max_ending_here {
		*max = max_ending_here
	}
	return max_ending_here
}

func lis_dp(arr []int) int {
	n := len(arr)
	if n == 1 {
		return 1
	}
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = 1
	}
	res := 1
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			if arr[j] < arr[i] {
				dp[i] = max(dp[i], dp[j]+1)
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
	fmt.Println(lis_dp([]int{10, 9, 2, 5, 3, 7, 101, 18}))
	fmt.Println(lis_dp([]int{0, 1, 0, 3, 2, 3}))
	fmt.Println(lis_dp([]int{7, 7, 7, 7, 7, 7}))
	fmt.Println(lis([]int{4, 10, 4, 3, 8, 9}))
	fmt.Println(lis([]int{0}))
}
