/*
You are given an integer array nums and an integer k.

For each index i where 0 <= i < nums.length, change nums[i] to be either nums[i] + k or nums[i] - k.

The score of nums is the difference between the maximum and minimum elements in nums.

Return the minimum score of nums after changing the values at each index.
*/

package main

import (
	"fmt"
	"sort"
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

func smallest_range(arr []int, k int) int {
	n := len(arr)
	sort.Ints(arr)

	ans := arr[n-1] - arr[0]
	for i := 0; i < n-1; i++ {
		a := arr[i]
		b := arr[i+1]
		high := max(arr[n-1]-k, a+k)
		low := min(arr[0]+k, b-k)
		ans = min(ans, high-low)
	}
	return ans
}

func main() {
	fmt.Println(smallest_range([]int{1, 3, 6}, 3))
}
