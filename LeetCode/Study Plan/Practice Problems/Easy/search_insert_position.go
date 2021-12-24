/*
Given a sorted array of distinct integers and a target value, return the index if the target is found.
If not, return the index where it would be if it were inserted in order.

You must write an algorithm with O(log n) runtime complexity.
*/

package main

import "fmt"

func search_position(nums []int, target int) int {
	n := len(nums)

	l, h := 0, n-1
	for l <= h {
		mid := (l + h) / 2
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
	fmt.Println(search_position([]int{1, 3, 5, 6}, 5))
	fmt.Println(search_position([]int{1, 3, 5, 6}, 7))
}
