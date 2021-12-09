/*
You are a professional robber planning to rob houses along a street. Each house has a certain amount of money stashed.
All houses at this place are arranged in a circle. That means the first house is the neighbor of the last one. Meanwhile,
adjacent houses have a security system connected, and it will automatically contact the police if two adjacent houses were
broken into on the same night.

Given an integer array nums representing the amount of money of each house, return the maximum amount of money you can rob
tonight without alerting the police.
*/

package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func house_robber(nums []int) int {
	n := len(nums)

	if n == 0 {
		return 0
	}

	if n == 1 {
		return nums[0]
	}

	return max(_house_robber(nums[1:]), _house_robber(nums[:n-1]))
}

func _house_robber(nums []int) int {
	n := len(nums)
	t1, t2 := 0, 0
	for i := 0; i < n; i++ {
		temp := t1
		t1 = max(nums[i]+t2, t1)
		t2 = temp
	}
	return t1
}

func main() {
	fmt.Println(house_robber([]int{1, 2, 3}))
	fmt.Println(house_robber([]int{2, 3, 2}))
	fmt.Println(house_robber([]int{1, 2, 3, 4, 5}))
	fmt.Println(house_robber([]int{1, 2, 3, 4, 5, 6}))
}
