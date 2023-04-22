/*
Given an integer array nums, find a contiguous non-empty subarray within the array that has the largest product,
and return the product.

The test cases are generated so that the answer will fit in a 32-bit integer.

A subarray is a contiguous subsequence of the array.
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

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max_product_bf(arr []int) int {
	n := len(arr)
	ans := arr[0]
	for i := 0; i < n; i++ {
		temp := 1
		for j := i; j < n; j++ {
			temp *= arr[j]
			ans = max(ans, temp)
		}
	}
	return ans
}

func max_product(arr []int) int {
	max_so_far := arr[0]
	min_so_far := arr[0]
	res := max_so_far
	n := len(arr)

	for i := 1; i < n; i++ {
		curr := arr[i]
		temp := max(curr, max(max_so_far*curr, min_so_far*curr))
		min_so_far = min(curr, min(max_so_far*curr, min_so_far*curr))
		max_so_far = temp
		res = max(res, max_so_far)
	}
	return res
}

func main() {
	fmt.Println(max_product([]int{2, 3, -2, 4}))
	fmt.Println(max_product([]int{-2, 0, -1}))
	fmt.Println(max_product([]int{-2, 3, -4}))
	fmt.Println(max_product([]int{2, -5, -2, -4, 3}))
}
