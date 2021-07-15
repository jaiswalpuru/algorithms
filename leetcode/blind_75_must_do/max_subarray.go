/*
Given an integer array nums, find the contiguous subarray (containing at least one number)
which has the largest sum and return its sum.

Example 1:
Input: nums = [-2,1,-3,4,-1,2,1,-5,4]
Output: 6
Explanation: [4,-1,2,1] has the largest sum = 6.

Example 2:
Input: nums = [1]
Output: 1

Example 3:
Input: nums = [5,4,-1,7,8]
Output: 23

Algo : Kadane's Algorithm
*/

package main

import (
	"fmt"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func largest_sum(arr []int) int {
	sum, temp := arr[0], arr[0]
	n := len(arr)
	for i := 1; i < n; i++ {
		temp = max(temp+arr[i], arr[i])
		sum = max(sum, temp)
	}
	return sum
}

func main() {
	fmt.Println(largest_sum([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
	fmt.Println(largest_sum([]int{1}))
	fmt.Println(largest_sum([]int{5, 4, -1, 7, 8}))
}
