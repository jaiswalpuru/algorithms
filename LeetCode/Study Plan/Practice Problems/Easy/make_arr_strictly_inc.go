/*
Given a 0-indexed integer array nums, return true if it can be made strictly increasing after removing exactly one element,
or false otherwise. If the array is already strictly increasing, return true.

The array nums is strictly increasing if nums[i - 1] < nums[i] for each index (1 <= i < nums.length).
*/

package main

import "fmt"

func can_be_increasing(a []int) bool {
	n := len(a)
	b, c := -1, -1
	for i := 0; i < n-1; i++ {
		if a[i] >= a[i+1] {
			if b != -1 {
				return false
			}
			b = i
			c = i + 1
		}
	}

	if b < 1 || c == n-1 {
		return true
	}
	if a[b-1] < a[c] || a[b] < a[c+1] {
		return true
	}
	return false
}

func main() {
	fmt.Println(can_be_increasing([]int{1, 2, 10, 5, 7}))
	fmt.Println(can_be_increasing([]int{2, 3, 1, 2}))
}
