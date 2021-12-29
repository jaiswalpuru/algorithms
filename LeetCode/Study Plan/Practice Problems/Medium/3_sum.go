/*
Given an integer array nums, return all the triplets [nums[i], nums[j], nums[k]] such that i != j, i != k, and j != k,
and nums[i] + nums[j] + nums[k] == 0.

Notice that the solution set must not contain duplicate triplets.
*/

package main

import (
	"fmt"
	"sort"
)

func three_sum(arr []int) [][]int {
	sort.Ints(arr)
	n := len(arr)
	res := make([][]int, 0)
	for i := 0; i < n && arr[i] <= 0; i++ {
		if i == 0 || arr[i-1] != arr[i] {
			two_sum_map(arr, i, &res)
		}
	}
	return res
}

//using two pointers approach
func two_sum_2p(arr []int, i int, res *[][]int) {
	l, h := i+1, len(arr)-1
	for l < h {
		sum := arr[i] + arr[l] + arr[h]
		if sum < 0 {
			l++
		} else if sum > 0 {
			h--
		} else {
			*res = append(*res, []int{arr[i], arr[l], arr[h]})
			l++
			h--
			//disregarding duplicates
			for l < h && arr[l] == arr[l-1] {
				l++
			}
		}
	}
}

//using map approach
func two_sum_map(arr []int, i int, res *[][]int) {
	mp := make(map[int]struct{})
	n := len(arr)

	for j := i + 1; j < n; j++ {
		inverse_sum := -arr[i] - arr[j]
		if _, ok := mp[inverse_sum]; ok {
			*res = append(*res, []int{arr[i], arr[j], inverse_sum})
			for j+1 < n && arr[j] == arr[j+1] {
				j++
			}
		} else {
			mp[arr[j]] = struct{}{}
		}
	}
}

func main() {
	fmt.Println(three_sum([]int{-1, 0, 1, 2, -1, -4}))
}
