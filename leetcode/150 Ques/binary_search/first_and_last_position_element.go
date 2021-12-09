/*
Given an array of integers nums sorted in ascending order, find the starting and ending position
of a given target value.

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
	res := -1
	low, high := 0, len(arr)-1
	for low <= high {
		mid := low + (high-low)/2
		if arr[mid] < target {
			low = mid + 1
		} else if arr[mid] > target {
			high = mid - 1
		} else {
			res = mid
			high = mid - 1
		}
	}
	return res
}

func upper_bound(arr []int, target int) int {
	res := -1
	low, high := 0, len(arr)-1
	for low <= high {
		mid := low + (high-low)/2
		if arr[mid] == target {
			res = mid
			low = mid + 1
		} else if arr[mid] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return res
}

func search(arr []int, target int) (int, int) {
	return lower_bound(arr, target), upper_bound(arr, target)
}

func main() {
	fmt.Println(search([]int{5, 7, 7, 8, 8, 10}, 8))
	fmt.Println(search([]int{5, 7, 7, 8, 8, 10}, 6))
	fmt.Println(search([]int{}, 0))
}
