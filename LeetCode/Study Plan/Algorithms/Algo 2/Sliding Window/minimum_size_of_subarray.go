/*
Given an array of positive integers nums and a positive integer target,
return the minimal length of a contiguous subarray [numsl, numsl+1, ..., numsr-1, numsr] of which the sum is greater
than or equal to target.
If there is no such subarray, return 0 instead.
*/

package main

import (
	"fmt"
	"math"
)

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func minimum_size_of_subaarray(arr []int, target int) int {
	sum := 0
	i, j := 0, 0
	n := len(arr)
	ans := int(math.MaxInt32)
	for j < n {
		sum += arr[j]
		for sum >= target && i <= j {
			sum -= arr[i]
			ans = min(ans, j-i+1)
			i++
		}
		j++
	}
	if ans == int(math.MaxInt32) {
		return 0
	}
	return ans
}

func main() {
	fmt.Println(minimum_size_of_subaarray([]int{2, 3, 1, 2, 4, 3}, 7))
	fmt.Println(minimum_size_of_subaarray([]int{1, 4, 4}, 4))
}
