/*
Given an integer array nums, find a contiguous non-empty subarray within the array
that has the largest product, and return the product.

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

import (
	"fmt"
	"math"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func max_prod_subarray(a []int) int {
	n := len(a)
	curr_so_far := 1
	max_so_far := int(math.MinInt32)

	for i := 0; i < n; i++ {
		curr_so_far *= a[i]
		max_so_far = max(max_so_far, curr_so_far)
		if curr_so_far == 0 {
			curr_so_far = 1
		}
	}

	curr_so_far = 1
	for i := n - 1; i >= 0; i-- {
		curr_so_far *= a[i]
		max_so_far = max(max_so_far, a[i])
		if curr_so_far == 0 {
			curr_so_far = 1
		}
	}
	return max_so_far
}

func main() {
	fmt.Println(max_prod_subarray([]int{2, 3, -2, 4}))
	fmt.Println(max_prod_subarray([]int{-2, 0, -1}))
	fmt.Println(max_prod_subarray([]int{2, 3, -2, -4}))
}
