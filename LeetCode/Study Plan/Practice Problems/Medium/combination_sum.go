/*
Given an array of distinct integers candidates and a target integer target, return a list of all unique combinations of
candidates where the chosen numbers sum to target. You may return the combinations in any order.

The same number may be chosen from candidates an unlimited number of times. Two combinations are unique if the frequency
of at least one of the chosen numbers is different.

It is guaranteed that the number of unique combinations that sum up to target is less than 150 combinations for the given input.
*/

package main

import "fmt"

func combination_sum(arr []int, target int) [][]int {
	res := [][]int{}
	_combination_sum(arr, &res, []int{}, 0, target)
	return res
}

func _combination_sum(arr []int, res *[][]int, set []int, start, target int) {
	if start == len(arr) {
		if target == 0 {
			*res = append(*res, append([]int{}, set...))
		}
		return
	}

	if arr[start] <= target {
		temp := append(set, arr[start])
		_combination_sum(arr, res, temp, start, target-arr[start])
		temp = temp[:len(temp)-1]
	}
	_combination_sum(arr, res, set, start+1, target)
}

func main() {
	fmt.Println(combination_sum([]int{2, 3, 6, 7}, 7))
	fmt.Println(combination_sum([]int{2, 3, 5}, 8))
}
