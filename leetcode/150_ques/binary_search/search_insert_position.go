/*
Given a sorted array of distinct integers and a target value, return the index if the target is found.
If not, return the index where it would be if it were inserted in order.

You must write an algorithm with O(log n) runtime complexity.

Example 1:
Input: nums = [1,3,5,6], target = 5
Output: 2

Example 2:
Input: nums = [1,3,5,6], target = 2
Output: 1

Example 3:
Input: nums = [1,3,5,6], target = 7
Output: 4

Example 4:
Input: nums = [1,3,5,6], target = 0
Output: 0

Example 5:
Input: nums = [1], target = 0
Output: 0
*/
package main

import "fmt"

func search_insert_position(arr []int, target int) int {
	n := len(arr)
	if n == 0 {
		return n
	}
	l, r := 0, n-1
	for l <= r {
		mid := l + (r-l)/2
		if arr[mid] == target {
			return mid
		} else if arr[mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return l
}

func main() {
	fmt.Println(search_insert_position([]int{1, 3, 5, 6}, 5))
	fmt.Println(search_insert_position([]int{1, 3, 5, 6}, 2))
	fmt.Println(search_insert_position([]int{1, 3, 5, 6}, 7))
	fmt.Println(search_insert_position([]int{1, 3, 5, 6}, 0))
	fmt.Println(search_insert_position([]int{1}, 0))
}
