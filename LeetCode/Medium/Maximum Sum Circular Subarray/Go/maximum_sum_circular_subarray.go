package main

import "fmt"

//-------------------using suffix sum + kadane---------------
func maximum_sum_circular_subarray(nums []int) int {
	n := len(nums)
	suffix_sum := make([]int, n)
	curr_sum := nums[n-1]
	suffix_sum[n-1] = nums[n-1]
	for i := n - 2; i >= 0; i-- {
		curr_sum += nums[i]
		suffix_sum[i] = max(suffix_sum[i+1], curr_sum)
	}
	prefix_sum := 0
	curr_sum = 0
	max_sum := nums[0]
	res := nums[0]
	for i := 0; i < n; i++ {
		curr_sum = max(curr_sum+nums[i], nums[i])
		max_sum = max(max_sum, curr_sum)
		prefix_sum += nums[i]
		if i+1 < n {
			res = max(res, prefix_sum+suffix_sum[i+1])
		}
	}
	return max(max_sum, res)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//-------------------using suffix sum + kadane---------------

func main() {
	fmt.Println(maximum_sum_circular_subarray([]int{1, -2, 3, -2}))
}
