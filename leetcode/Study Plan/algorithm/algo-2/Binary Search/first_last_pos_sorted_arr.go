/*
Given an array of integers nums sorted in non-decreasing order, find the starting and ending position of a given target value.

If target is not found in the array, return [-1, -1].

You must write an algorithm with O(log n) runtime complexity.
*/

package main

import "fmt"

func first_last_pos(arr []int, target int) []int {
	ans := []int{-1, -1}

	//Will run binary search two times
	l, h := 0, len(arr)-1
	for l <= h {
		mid := (l + h) / 2
		if arr[mid] == target {
			ans[0] = mid
			ans[1] = mid
			l = mid + 1
		} else if arr[mid] > target {
			h = mid - 1
		} else {
			l = mid + 1
		}
	}
	l, h = 0, len(arr)-1
	for l <= h {
		mid := (l + h) / 2
		if arr[mid] == target {
			ans[0] = mid
			h = mid - 1
		} else if arr[mid] > target {
			h = mid - 1
		} else {
			l = mid + 1
		}
	}

	return ans

}

func main() {
	fmt.Println(first_last_pos([]int{5, 7, 7, 8, 8, 10}, 8))
	fmt.Println(first_last_pos([]int{5, 7, 7, 8, 8, 10}, 6))
	fmt.Println(first_last_pos([]int{1}, 1))
}
