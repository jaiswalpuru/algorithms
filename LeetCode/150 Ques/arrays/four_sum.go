/*
Given an array nums of n integers, return an array of all the unique quadruplets
[nums[a], nums[b], nums[c], nums[d]] such that:

0 <= a, b, c, d < n
a, b, c, and d are distinct.
nums[a] + nums[b] + nums[c] + nums[d] == target
You may return the answer in any order.

Example 1:

Input: nums = [1,0,-1,0,-2,2], target = 0
Output: [[-2,-1,1,2],[-2,0,0,2],[-1,0,0,1]]

Example 2:

Input: nums = [2,2,2,2,2], target = 8
Output: [[2,2,2,2]]
*/

package main

import (
	"fmt"
	"sort"
)

func four_sum(arr []int, target int) [][]int {
	n := len(arr)
	res := [][]int{}
	sort.Ints(arr)

	for i := 0; i < n-3; i++ {
		for j := i + 1; j < n-2; j++ {
			l := j + 1
			r := n - 1

			for l < r {
				if arr[i]+arr[j]+arr[l]+arr[r] == target {
					res = append(res, []int{arr[i], arr[j], arr[l], arr[r]})
					l++
					r--
				} else if arr[i]+arr[j]+arr[l]+arr[r] < target {
					l++
				} else {
					r--
				}
			}
		}
	}

	return res
}

func main() {
	fmt.Println(four_sum([]int{1, 0, -1, 0, -2, 2}, 0))
	fmt.Println(four_sum([]int{2, 2, 2, 2}, 8))
}
