/*
Given an array of distinct integers nums and a target integer target, return the number of possible combinations that
add up to target.

The test cases are generated so that the answer can fit in a 32-bit integer.
*/

package main

import (
	"fmt"
)

//---------------------------Brute force using backtracking---------------------------
func combination_sum_IV(arr []int, target int) int {
	res := [][]int{}
	_combination_sum(arr, []int{}, &res, target, 0)
	fmt.Println(res)
	return len(res)
}

func _combination_sum(arr []int, set []int, res *[][]int, target int, start int) {
	if target < 0 {
		return
	}
	if target == 0 {
		*res = append(*res, append([]int{}, set...))
		return
	}

	for i := 0; i < len(arr); i++ {
		temp := append(set, arr[i])
		_combination_sum(arr, temp, res, target-arr[i], i+1)
		temp = temp[:len(temp)-1]
	}
}

//---------------------------brute force end ---------------------------

//-------------------------------top down dp-------------------------------
func top_down_combination_sum_IV(arr []int, sum int) int {
	memo := make(map[int]int)
	return _top_down_combination_sum_IV(arr, sum, &memo)
}

func _top_down_combination_sum_IV(arr []int, remain int, memo *map[int]int) int {
	if remain == 0 {
		return 1
	}

	if _, ok := (*memo)[remain]; ok {
		return (*memo)[remain]
	}

	res := 0
	for i := 0; i < len(arr); i++ {
		if remain-arr[i] >= 0 {
			res += _top_down_combination_sum_IV(arr, remain-arr[i], memo)
		}
	}
	(*memo)[remain] = res
	return res
}

//-------------------------------------------------------------------------

//-------------------------------bottom up dp-------------------------------
func bottom_up_combination_sum_IV(arr []int, sum int) int {
	dp := make([]int, sum+1)
	dp[0] = 1
	n := len(arr)
	for i := 1; i <= sum; i++ {
		for j := 0; j < n; j++ {
			if i-arr[j] >= 0 {
				dp[i] += dp[i-arr[j]]
			}
		}
	}
	return dp[sum]
}

//--------------------------------------------------------------------------

func main() {
	fmt.Println(bottom_up_combination_sum_IV([]int{10, 1, 2, 7, 6, 5}, 8))
}
