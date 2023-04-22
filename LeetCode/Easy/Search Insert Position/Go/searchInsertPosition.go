package main

import "fmt"

func searchInsert(nums []int, target int) int {
	n := len(nums)

	l, h, mid := 0, n-1, -1
	for l <= h {
		mid = (l + h) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			h = mid - 1
		} else {
			l = mid + 1
		}
	}
	return l
}

func main() {
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 5))
}
