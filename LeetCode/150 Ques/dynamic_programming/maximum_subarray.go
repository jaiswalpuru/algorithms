/*
Given an integer array nums, find the contiguous subarray (containing at least one number) which has
the largest sum and return its sum.

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
*/

package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//dp approach
func maximum_subarray(arr []int) int {
	n := len(arr)
	res := make([]int, n)
	res[0] = arr[0]
	t := arr[0]
	for i := 1; i < n; i++ {
		t = max(arr[i], t+arr[i])
		res[i] = max(res[i-1], t)
	}
	return res[n-1]
}

func _maximum_subarray_recursive(arr []int, sum *int, temp int, ind int) {
	if ind == len(arr) {
		return
	}
	temp = max(arr[ind], temp+arr[ind])
	*sum = max(*sum, temp)
	_maximum_subarray_recursive(arr, sum, temp, ind+1)
}

func maximum_subarray_recursive(arr []int) int {
	sum := arr[0]
	temp := arr[0]
	_maximum_subarray_recursive(arr, &sum, temp, 1)
	return sum
}

func main() {
	fmt.Println(maximum_subarray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
	fmt.Println(maximum_subarray([]int{5, 4, -1, 7, 8}))
	fmt.Println(maximum_subarray_recursive([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
	fmt.Println(maximum_subarray_recursive([]int{5, 4, -1, 7, 8}))
	fmt.Println(maximum_subarray([]int{1}))
	fmt.Println(maximum_subarray_recursive([]int{1}))
}
