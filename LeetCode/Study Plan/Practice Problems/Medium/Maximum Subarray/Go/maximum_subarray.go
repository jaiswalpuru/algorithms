package main

import (
	"fmt"
	"math"
)

//------------------divide and conquer------------
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maximum_subarray_divide_and_conquer(nums []int) int {
	return max_sub_sum(nums, 0, len(nums)-1)
}
func max_sub_sum(nums []int, l, r int) int {
	if l > r {
		return math.MinInt64
	}
	mid := (l + r) / 2
	fmt.Println(mid)
	curr := 0
	best_left_sum := 0
	best_right_sum := 0
	for i := mid - 1; i >= l; i-- {
		curr += nums[i]
		best_left_sum = max(best_left_sum, curr)
	}
	curr = 0
	for i := mid + 1; i <= r; i++ {
		curr += nums[i]
		best_right_sum = max(best_right_sum, curr)
	}
	combined := nums[mid] + best_left_sum + best_right_sum
	left := max_sub_sum(nums, l, mid-1)
	right := max_sub_sum(nums, mid+1, r)
	return max(combined, max(left, right))
}

//------------------divide and conquer------------

//----------------kadane algorithm---------------
func maximum_subarray_kadane(nums []int) int {
	sum := nums[0]
	res := sum
	n := len(nums)
	for i := 1; i < n; i++ {
		sum = max(sum+nums[i], nums[i])
		res = max(res, sum)
	}
	return res
}

//----------------kadane algorithm---------------

func main() {
	fmt.Println(maximum_subarray_kadane([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
}
