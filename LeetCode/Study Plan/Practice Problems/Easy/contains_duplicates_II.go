/*
Given an integer array nums and an integer k, return true if there are two distinct indices i and j
in the array such that nums[i] == nums[j] and abs(i - j) <= k.
*/

package main

import (
	"fmt"
	"math"
)

func contains_dup(arr []int, k int) bool {
	n := len(arr)
	if len(arr) == 1 {
		return false
	}

	mp := make(map[int]int)
	for i := 0; i < n; i++ {
		if _, ok := mp[arr[i]]; ok {
			ind := mp[arr[i]]
			if int(math.Abs(float64(ind-i))) <= k {
				return true
			} else {
				mp[arr[i]] = i
			}
		} else {
			mp[arr[i]] = i
		}
	}
	return false
}

func main() {
	fmt.Println(contains_dup([]int{1, 2, 3, 1}, 3))
	fmt.Println(contains_dup([]int{1, 0, 1, 1}, 1))
	fmt.Println(contains_dup([]int{1, 2, 3, 1, 2, 3}, 2))
}
