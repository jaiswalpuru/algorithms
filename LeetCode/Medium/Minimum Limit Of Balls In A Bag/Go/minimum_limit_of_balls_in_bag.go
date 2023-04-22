package main

import "fmt"
import "math"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func possible(mid, op int, nums []int) bool {
	for i := 0; i < len(nums); i++ {
		ops := (nums[i]-1) / mid
		if ops <= op {
			op -= ops
		} else {
			return false
		}
	}
	return true
}

func minimum_limit_of_balls_in_bag(nums []int, op int) int {
	max_val := math.MinInt64
	n := len(nums)

	for i := 0; i < n; i++ {
		max_val = max(max_val, nums[i])
	}

	l, h := 1, max_val
	res := 0
	for l <= h {
		mid := (l + h) / 2
		if possible(mid, op, nums) {
			res = mid
			h = mid - 1
		} else {
			l = mid + 1
		}
	}
	return res
}

func main() {
	fmt.Println(minimum_limit_of_balls_in_bag([]int{2, 4, 8, 2}, 4))
}
