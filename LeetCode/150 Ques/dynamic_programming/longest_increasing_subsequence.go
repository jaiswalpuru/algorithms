/*
Given an integer array nums, return the length of the longest strictly increasing subsequence.

A subsequence is a sequence that can be derived from an array by deleting some or no elements
without changing the order of the remaining elements. For example, [3,6,2,7] is a subsequence
of the array [0,3,1,6,2,2,7].

Example 1:
Input: nums = [10,9,2,5,3,7,101,18]
Output: 4
Explanation: The longest increasing subsequence is [2,3,7,101], therefore the length is 4.

Example 2:
Input: nums = [0,1,0,3,2,3]
Output: 4

Example 3:
Input: nums = [7,7,7,7,7,7,7]
Output: 1
*/

package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func lis(arr []int) int {
	n := len(arr)
	if n == 1 {
		return 1
	}

	max_len := 1
	dp := make([]int, n)

	for i := 0; i < n; i++ {
		dp[i] = 1
	}

	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			if arr[j] < arr[i] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		max_len = max(max_len, dp[i])
	}

	return max_len
}

func main() {
	fmt.Println(lis([]int{10, 9, 2, 5, 3, 7, 101, 18}))
	fmt.Println(lis([]int{0, 1, 0, 3, 2, 3}))
	fmt.Println(lis([]int{7, 7, 7, 7, 7, 7, 7}))
}
