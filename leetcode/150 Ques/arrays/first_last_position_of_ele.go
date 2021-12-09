/*
Given an array of integers nums sorted in ascending order, find the starting and
ending position of a given target value.

If target is not found in the array, return [-1, -1].

You must write an algorithm with O(log n) runtime complexity.



Example 1:

Input: nums = [5,7,7,8,8,10], target = 8
Output: [3,4]
Example 2:

Input: nums = [5,7,7,8,8,10], target = 6
Output: [-1,-1]
Example 3:

Input: nums = [], target = 0
Output: [-1,-1]
*/

package main

import "fmt"

func lower_bound(arr []int, target int) int {
	l := 0
	h := len(arr)
	res := -1
	for l <= h {
		mid := l + (h-l)/2
		if arr[mid] > target {
			h = mid - 1
		} else if arr[mid] < target {
			l = mid + 1
		} else {
			res = mid
			h = mid - 1
		}
	}
	return res
}

func upper_bound(arr []int, target int) int {
	l := 0
	h := len(arr)
	res := -1
	for l <= h {
		mid := l + (h-l)/2
		if arr[mid] > target {
			h = mid - 1
		} else if arr[mid] < target {
			l = mid + 1
		} else {
			res = mid
			l = mid + 1
		}
	}
	return res
}

//array is sorted
func first_and_last_pos(arr []int, target int) (int, int) {
	n := len(arr)
	if n == 0 {
		return -1, -1
	}
	first, last := lower_bound(arr, target), upper_bound(arr, target)
	return first, last
}

func main() {
	fmt.Println(first_and_last_pos([]int{5, 7, 7, 8, 8, 10}, 8))
	fmt.Println(first_and_last_pos([]int{5, 7, 7, 8, 8, 10}, 6))
	fmt.Println(first_and_last_pos([]int{}, 0))
}
