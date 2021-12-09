/*
Given an integer array nums, find a contiguous non-empty subarray within the array that has the
largest product, and return the product.

It is guaranteed that the answer will fit in a 32-bit integer.

A subarray is a contiguous subsequence of the array.

Example 1:
Input: nums = [2,3,-2,4]
Output: 6
Explanation: [2,3] has the largest product 6.

Example 2:
Input: nums = [-2,0,-1]
Output: 0
Explanation: The result cannot be 2, because [-2,-1] is not a subarray.
*/

package main

import "fmt"

func max_product(a []int) int {
	n := len(a)
	for i := 1; i < n; i++ {
		a[i] = a[i] * a[i-1]
	}
	m := -999999999
	for i := 0; i < n; i++ {
		m = max(a[i], m)
	}
	return m
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(max_product([]int{2, 3, -2, 4}))
}
