package main

import "fmt"

func search_in_rotated_sorted_array(nums []int, target int) int {
	return _search(nums, target, 0, len(nums)-1)
}

func _search(nums []int, target int, l, r int) int {
	if l > r {
		return -1
	}
	mid := (l + r) / 2
	if nums[mid] == target {
		return mid
	}
	if nums[l] <= nums[mid] {
		if target >= nums[l] && target < nums[mid] {
			return _search(nums, target, l, mid-1)
		}
		return _search(nums, target, mid+1, r)
	}
	if target > nums[mid] && target <= nums[r] {
		return _search(nums, target, mid+1, r)
	} else {
		return _search(nums, target, l, mid-1)
	}
}

func main() {
	fmt.Println(search_in_rotated_sorted_array([]int{4, 5, 6, 7, 0, 1, 2}, 0))
}
