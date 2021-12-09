/*
Given an integer array nums, return all the triplets [nums[i], nums[j], nums[k]] such that i != j, i != k, and j != k, and
nums[i] + nums[j] + nums[k] == 0.

Notice that the solution set must not contain duplicate triplets.
*/

package main

import (
	"fmt"
	"sort"
)

func two_sum(arr []int, i int, res *[][]int) {
	l, h := i+1, len(arr)-1
	for l < h {
		sum := arr[i] + arr[l] + arr[h]
		if sum == 0 {
			*res = append(*res, []int{arr[i], arr[l], arr[h]})
			l++
			h--
			for l < h && arr[l] == arr[h] {
				l++
			}
		} else if sum > 0 {
			h--
		} else {
			l++
		}
	}
}

func three_sum(arr []int) [][]int {
	n := len(arr)
	ans := make([][]int, 0)
	sort.Ints(arr)
	for i := 0; i < n && arr[i] <= 0; i++ {
		if i == 0 || arr[i] != arr[i-1] {
			two_sum(arr, i, &ans)
		}
	}
	return ans
}

func main() {
	fmt.Println(three_sum([]int{-1, 0, 1, 2, -1, -4}))
}
