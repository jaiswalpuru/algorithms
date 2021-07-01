/*
Given an integer array nums, return all the triplets [nums[i], nums[j], nums[k]]
such that i != j, i != k, and j != k, and nums[i] + nums[j] + nums[k] == 0.

Notice that the solution set must not contain duplicate triplets.



Example 1:

Input: nums = [-1,0,1,2,-1,-4]
Output: [[-1,-1,2],[-1,0,1]]


Example 2:

Input: nums = []
Output: []


Example 3:

Input: nums = [0]
Output: []
*/

package main

import (
	"fmt"
)

func three_sum(arr []int, target int) [][]int {
	n := len(arr)
	res := [][]int{}

	for i := 0; i < n; i++ {
		k := target - arr[i]
		mp := make(map[int]bool)
		for j := i + 1; j < n; j++ {
			if _, ok := mp[k-arr[j]]; ok {
				res = append(res, []int{arr[i], arr[j], k - arr[j]})
			} else {
				mp[arr[j]] = true
			}
		}
	}
	temp := [][]int{}
	for i := 0; i < len(res); i++ {
		for j := i + 1; j < len(res); j++ {
			if (res[i][0] == res[j][0] || res[i][1] == res[j][0] || res[i][2] == res[j][0]) &&
				(res[i][0] == res[j][1] || res[i][1] == res[j][1] || res[i][2] == res[j][1]) &&
				(res[i][0] == res[j][2] || res[i][1] == res[j][2] || res[i][2] == res[j][2]) {
				res[j][0] = -99999
				res[j][1] = -99999
				res[j][2] = -99999
			}
		}
	}

	for i := 0; i < len(res); i++ {
		if res[i][0] != -99999 {
			temp = append(temp, res[i])
		}
	}

	return temp
}

func main() {
	arr := []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(three_sum(arr, 0))
}
