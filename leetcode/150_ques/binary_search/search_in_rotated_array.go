/*
There is an integer array nums sorted in ascending order (with distinct values).

Prior to being passed to your function, nums is rotated at an unknown pivot index k (0 <= k < nums.length)
such that the resulting array is [nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]] (0-indexed).
For example, [0,1,2,4,5,6,7] might be rotated at pivot index 3 and become [4,5,6,7,0,1,2].

Given the array nums after the rotation and an integer target, return the index of target if it is in nums,
or -1 if it is not in nums.

You must write an algorithm with O(log n) runtime complexity.

Example 1:
Input: nums = [4,5,6,7,0,1,2], target = 0
Output: 4

Example 2:
Input: nums = [4,5,6,7,0,1,2], target = 3
Output: -1

Example 3:
Input: nums = [1], target = 0
Output: -1
*/

package main

import "fmt"

func _search(arr []int, target int) int {
	n := len(arr)
	if n == 1 && arr[0] == target {
		return 0
	}
	l, r := 0, n-1
	for l <= r {
		mid := l + (r-l)/2
		if arr[mid] == target {
			return mid
		} else if target > arr[mid] {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return -1
}

func search(arr []int, target int) int {
	pivot := 0
	n := len(arr)
	for i := 1; i < n; i++ {
		if arr[i] < arr[i-1] {
			pivot = i
			break
		}
	}
	lower_search := _search(arr[:pivot], target)
	upper_search := _search(arr[pivot:], target)
	if lower_search != -1 {
		return lower_search
	}
	if upper_search != -1 {
		return pivot + upper_search
	}
	return -1
}

func main() {
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 0))
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 3))
	fmt.Println(search([]int{1}, 0))
}
