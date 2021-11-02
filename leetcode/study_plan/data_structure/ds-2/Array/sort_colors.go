/*
Given an array nums with n objects colored red, white, or blue, sort them in-place so that objects of the same color
are adjacent, with the colors in the order red, white, and blue.

We will use the integers 0, 1, and 2 to represent the color red, white, and blue, respectively.

You must solve this problem without using the library's sort function.
*/
package main

import "fmt"

func sort_colors(nums []int) []int {
	mp := make(map[int]int)
	n := len(nums)
	for i := 0; i < n; i++ {
		mp[nums[i]]++
	}
	i, j, k, zero, one, two := 0, 0, 0, mp[0], mp[1], mp[2]
	for i = 0; i < zero; i++ {
		nums[i] = 0
	}
	for j = i; j < i+one; j++ {
		nums[j] = 1
	}
	for k = j; k < j+two; k++ {
		nums[k] = 2
	}
	return nums
}

func main() {
	fmt.Println(sort_colors([]int{2, 0, 2, 1, 1, 0}))
}
