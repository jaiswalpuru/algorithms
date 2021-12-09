/*
Given a binary array nums and an integer k, return the maximum number of consecutive 1's in
the array if you can flip at most k 0's.

Example 1:
Input: nums = [1,1,1,0,0,0,1,1,1,1,0], k = 2
Output: 6
Explanation: [1,1,1,0,0,1,1,1,1,1,1]
Bolded numbers were flipped from 0 to 1. The longest subarray is underlined.

Example 2:
Input: nums = [0,0,1,1,0,0,1,1,1,0,1,1,0,0,0,1,1,1,1], k = 3
Output: 10
Explanation: [0,0,1,1,1,1,1,1,1,1,1,1,0,0,0,1,1,1,1]
Bolded numbers were flipped from 0 to 1. The longest subarray is underlined.
*/

package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func max_consequtive_ones(arr []int, k int) int {
	ans := 0
	n := len(arr)
	i := 0

	for j := 0; j < n; j++ {
		if arr[j] == 0 {
			k--
		}

		for k < 0 {
			if arr[i] == 0 {
				k++
			}
			i++
		}

		ans = max(ans, j-i+1)
	}

	return ans
}

func main() {
	fmt.Println(max_consequtive_ones([]int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}, 2))
}
