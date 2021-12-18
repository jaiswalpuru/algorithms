/*
A peak element is an element that is strictly greater than its neighbors.

Given an integer array nums, find a peak element, and return its index. If the array contains multiple peaks,
return the index to any of the peaks.

You may imagine that nums[-1] = nums[n] = -∞.

You must write an algorithm that runs in O(log n) time.
*/

package main

import "fmt"

func find_peak_element(a []int) int {
	l, h := 0, len(a)-1
	for l < h {
		mid := (l + h) / 2
		if a[mid] > a[mid+1] {
			h = mid
		} else {
			l = mid + 1
		}
	}
	return l
}

func main() {
	fmt.Println(find_peak_element([]int{1, 2, 3, 1}))
	fmt.Println(find_peak_element([]int{1}))
}
