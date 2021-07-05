/*
Given an array of integers nums which is sorted in ascending order, and an integer target,
write a function to search target in nums. If target exists, then return its index. Otherwise, return -1.

You must write an algorithm with O(log n) runtime complexity.

Example 1:

Input: nums = [-1,0,3,5,9,12], target = 9
Output: 4
Explanation: 9 exists in nums and its index is 4


Example 2:

Input: nums = [-1,0,3,5,9,12], target = 2
Output: -1
Explanation: 2 does not exist in nums so return -1
*/

package main

import "fmt"

func binary_search(arr []int, target int) int {
	if len(arr) == 0 {
		return -1
	}
	low, high := 0, len(arr)
	for low <= high {
		mid := low + (high-low)/2
		if arr[mid] == target {
			return mid
		} else if arr[mid] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}

func main() {
	fmt.Println(binary_search([]int{-1, 0, 3, 5, 9, 12}, 9))
	fmt.Println(binary_search([]int{-1, 0, 3, 5, 9, 12}, 2))
}
