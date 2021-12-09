/*
There is an integer array nums sorted in ascending order (with distinct values).

Prior to being passed to your function, nums is possibly rotated at an unknown pivot index k (1 <= k < nums.length)
such that the resulting array is [nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]] (0-indexed).
For example, [0,1,2,4,5,6,7] might be rotated at pivot index 3 and become [4,5,6,7,0,1,2].

Given the array nums after the possible rotation and an integer target, return the index of target if it is in nums,
or -1 if it is not in nums.

You must write an algorithm with O(log n) runtime complexity.
*/

package main

import "fmt"

func search_rotated_sorted_array(arr []int, target int) int {
	l, h := 0, len(arr)-1
	return _search_rotated_sorted_array(arr, l, h, target)
}

func _search_rotated_sorted_array(arr []int, l, h, target int) int {
	if l > h {
		return -1
	}
	mid := (l + h) / 2
	if arr[mid] == target {
		return mid
	}
	if arr[l] <= arr[mid] {
		if target >= arr[l] && target <= arr[mid] {
			return _search_rotated_sorted_array(arr, l, mid-1, target)
		}
		return _search_rotated_sorted_array(arr, mid+1, h, target)
	}
	if target >= arr[mid] && target <= arr[h] {
		return _search_rotated_sorted_array(arr, mid+1, h, target)
	} else {
		return _search_rotated_sorted_array(arr, l, mid-1, target)
	}
}

func main() {
	fmt.Println(search_rotated_sorted_array([]int{4, 5, 6, 7, 0, 1, 2}, 0))
}
