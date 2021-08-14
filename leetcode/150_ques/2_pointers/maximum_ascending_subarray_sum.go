/*
Given an array of positive integers nums, return the maximum possible sum of an ascending subarray in nums.

A subarray is defined as a contiguous sequence of numbers in an array.

A subarray [numsl, numsl+1, ..., numsr-1, numsr] is ascending if for all i where l <= i < r, numsi < numsi+1.
Note that a subarray of size 1 is ascending.

Example 1:
Input: nums = [10,20,30,5,10,50]
Output: 65
Explanation: [5,10,50] is the ascending subarray with the maximum sum of 65.

Example 2:
Input: nums = [10,20,30,40,50]
Output: 150
Explanation: [10,20,30,40,50] is the ascending subarray with the maximum sum of 150.

Example 3:
Input: nums = [12,17,15,13,10,11,12]
Output: 33
Explanation: [10,11,12] is the ascending subarray with the maximum sum of 33.

Example 4:
Input: nums = [100,10,1]
Output: 100
*/

package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func max_ascending_subarray_sum(arr []int) int {
	n := len(arr)
	ans, temp := arr[0], arr[0]
	for i := 1; i < n; i++ {
		if arr[i-1] < arr[i] {
			temp += arr[i]
		} else {
			ans = max(ans, temp)
			temp = arr[i]
		}
	}
	ans = max(temp, ans)
	return ans
}

func main() {
	fmt.Println(max_ascending_subarray_sum([]int{10, 20, 30, 5, 10, 50}))
}
