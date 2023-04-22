package main

import "fmt"

func find_minimum_in_rotated_sorted_array(nums []int) int {
	res := nums[0]
	nums = nums[1:]
	l, r := 0, len(nums)-1
	for l <= r {
		mid := (l + r) / 2
		if res > nums[mid] {
			res = nums[mid]
			r = mid - 1
		}
		if res < nums[mid] {
			l = mid + 1
		}
	}
	return res
}

func main() {
	fmt.Println(find_minimum_in_rotated_sorted_array([]int{3, 4, 5, 1, 2}))
}
