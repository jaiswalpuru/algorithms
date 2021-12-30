/*
Given an integer array arr and an integer k, modify the array by repeating it k times.

For example, if arr = [1, 2] and k = 3 then the modified array will be [1, 2, 1, 2, 1, 2].

Return the maximum sub-array sum in the modified array. Note that the length of the sub-array can be 0 and its sum in
that case is 0.

As the answer can be very large, return the answer modulo 109 + 7.
*/

package main

import (
	"fmt"
	"math"
)

const (
	mod = 1_000_000_007
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func kadane(arr []int) int {
	max_val, curr_max := arr[0], arr[0]
	n := len(arr)
	for i := 1; i < n; i++ {
		curr_max = max(arr[i], curr_max+arr[i])
		max_val = max(max_val, curr_max) % mod
	}
	return max_val
}

func k_concatenation(arr []int, k int) int {
	n := len(arr)

	max_sum := kadane(arr)
	if k == 1 {
		return max_sum
	}

	total_sum := 0
	for i := 0; i < n; i++ {
		total_sum = (total_sum + arr[i]) % mod
	}

	//prefixSum
	prefix_sum := 0
	maximum := int(math.MinInt32)
	for i := 0; i < n; i++ {
		prefix_sum += arr[i]
		maximum = max(maximum, prefix_sum)
	}
	if maximum < 0 {
		prefix_sum = 0
	} else {
		prefix_sum = maximum
	}

	//suffix sum
	suffix_sum := 0
	maximum = int(math.MinInt32)
	for i := n - 1; i >= 0; i-- {
		suffix_sum += arr[i]
		maximum = max(maximum, suffix_sum)
	}
	if maximum < 0 {
		suffix_sum = 0
	} else {
		suffix_sum = maximum
	}

	if total_sum > 0 {
		return max((total_sum*(k-2)+prefix_sum+suffix_sum)%mod, (total_sum*k)%mod)
	}
	return max(max_sum, (prefix_sum+suffix_sum)%mod)
}

func main() {
	fmt.Println(k_concatenation([]int{1, 2}, 3))
	// fmt.Println(k_concatenation([]int{2, -5, 1, 0, -2, -2, 2}, 2))
}
