/*
Given a collection of candidate numbers (candidates) and a target number (target), find all unique
combinations in candidates where the candidate numbers sum to target.

Each number in candidates may only be used once in the combination.

Note: The solution set must not contain duplicate combinations.
*/

package main

import (
	"fmt"
	"sort"
)

func combination_sum_II(arr []int, target int) [][]int {
	res := [][]int{}
	sort.Ints(arr)
	_combination_sum_II(arr, &res, []int{}, target, 0)
	return res
}

func _combination_sum_II(arr []int, res *[][]int, temp []int, target, start int) {

	if target == 0 {
		*res = append(*res, append([]int{}, temp...))
		return
	}

	for i := start; i < len(arr); i++ {
		if i > start && arr[i] == arr[i-1] {
			continue
		}
		if target-arr[i] < 0 {
			break
		}
		temp = append(temp, arr[i])
		_combination_sum_II(arr, res, temp, target-arr[i], i+1)
		temp = temp[:len(temp)-1]
	}
}

func main() {
	fmt.Println(combination_sum_II([]int{10, 1, 2, 7, 6, 1, 5}, 8))
}
