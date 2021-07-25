/*
Given an array nums of integers, return the length of the longest arithmetic subsequence in nums.

Recall that a subsequence of an array nums is a list nums[i1], nums[i2], ..., nums[ik] with
0 <= i1 < i2 < ... < ik <= nums.length - 1, and that a sequence seq is arithmetic if seq[i+1] - seq[i]
are all the same value (for 0 <= i < seq.length - 1).

Example 1:
Input: nums = [3,6,9,12]
Output: 4
Explanation:
The whole array is an arithmetic sequence with steps of length = 3.

Example 2:
Input: nums = [9,4,7,2,10]
Output: 3
Explanation:
The longest arithmetic subsequence is [4,7,10].

Example 3:
Input: nums = [20,1,15,3,10,5,8]
Output: 4
Explanation:
The longest arithmetic subsequence is [20,15,10,5].
*/

package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func longest_arithmetic_subsequence(arr []int) int {
	n := len(arr)
	if n <= 2 {
		return n
	}
	dp := make([]map[int]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make(map[int]int)
	}

	ans := 2
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			diff := arr[j] - arr[i]
			if _, ok := dp[i][diff]; ok {
				dp[j][diff] = dp[i][diff] + 1
			} else {
				dp[j][diff] = 2
			}
			ans = max(ans, dp[j][diff])
		}
	}

	return ans
}

func main() {
	fmt.Println(longest_arithmetic_subsequence([]int{3, 6, 9, 12}))
	fmt.Println(longest_arithmetic_subsequence([]int{9, 4, 7, 2, 10}))
	fmt.Println(longest_arithmetic_subsequence([]int{20, 1, 15, 3, 10, 5, 8}))
	fmt.Println(longest_arithmetic_subsequence([]int{83, 20, 17, 43, 52, 78, 68, 45}))
}
