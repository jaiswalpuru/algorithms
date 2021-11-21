/*
Given a collection of candidate numbers (candidates) and a target number (target),
find all unique combinations in candidates where the candidate numbers sum to target.

Each number in candidates may only be used once in the combination.

Note: The solution set must not contain duplicate combinations.

Example 1:
Input: candidates = [10,1,2,7,6,1,5], target = 8
Output:
[
[1,1,6],
[1,2,5],
[1,7],
[2,6]
]

Example 2:
Input: candidates = [2,5,2,1,2], target = 5
Output:
[
[1,2,2],
[5]
]
*/

package main

import (
	"fmt"
	"sort"
)

func _combination_sum2(arr []int, target int, temp []int, res *[][]int, ind int) {

	if target == 0 {
		*res = append(*res, append([]int{}, temp...))
		return
	}

	for i := ind; i < len(arr); i++ {
		if i > ind && arr[i] == arr[i-1] {
			//skip
			continue
		}

		//optimization : early stop
		if target-arr[i] < 0 {
			break
		}

		temp = append(temp, arr[i])
		_combination_sum2(arr, target-arr[i], temp, res, i+1)
		temp = temp[:len(temp)-1]
	}
}

func combination_sum2(arr []int, target int) [][]int {
	res := [][]int{}
	var temp []int
	sort.Ints(arr)
	_combination_sum2(arr, target, temp, &res, 0)
	return res
}

func main() {
	fmt.Println(combination_sum2([]int{10, 1, 2, 7, 6, 1, 5}, 8))
	fmt.Println(combination_sum2([]int{2, 5, 2, 1, 2}, 5))
}
