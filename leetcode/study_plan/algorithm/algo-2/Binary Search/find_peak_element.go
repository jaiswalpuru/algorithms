/*
A peak element is an element that is strictly greater than its neighbors.

Given an integer array nums, find a peak element, and return its index. If the array contains multiple peaks,
return the index to any of the peaks.

You may imagine that nums[-1] = nums[n] = -∞.

You must write an algorithm that runs in O(log n) time.
*/

package main

import "fmt"

func find_peak_element(arr []int) int {
	return _find(arr, 0, len(arr)-1)
}

func _find(arr []int, l, h int) int {
	if l == h {
		return l
	}
	mid := (l + h) / 2
	if arr[mid] > arr[mid+1] {
		return _find(arr, l, mid)
	}
	return _find(arr, mid+1, h)
}

func main() {
	fmt.Println(find_peak_element([]int{1, 2, 3, 1}))
	fmt.Println(find_peak_element([]int{1, 2, 1, 3, 5, 6, 4}))
	fmt.Println(find_peak_element([]int{1, 2}))
	fmt.Println(find_peak_element([]int{0}))
	fmt.Println(find_peak_element([]int{1, 3, 2}))
	fmt.Println(find_peak_element([]int{1, 2, 3}))
	fmt.Println(find_peak_element([]int{2, 1}))
}
