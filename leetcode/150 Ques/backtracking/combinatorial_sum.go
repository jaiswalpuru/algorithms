/*
Given an array of distinct integers candidates and a target integer target,
return a list of all unique combinations of candidates where the chosen numbers sum to target.
You may return the combinations in any order.

The same number may be chosen from candidates an unlimited number of times.
Two combinations are unique if the frequency of at least one of the chosen numbers is different.

It is guaranteed that the number of unique combinations that sum up to target is less than 150 combinations
for the given input.

Example 1:
Input: candidates = [2,3,6,7], target = 7
Output: [[2,2,3],[7]]
Explanation:
2 and 3 are candidates, and 2 + 2 + 3 = 7. Note that 2 can be used multiple times.
7 is a candidate, and 7 = 7.
These are the only two combinations.

Example 2:
Input: candidates = [2,3,5], target = 8
Output: [[2,2,2,2],[2,3,3],[3,5]]

Example 3:
Input: candidates = [2], target = 1
Output: []

Example 4:
Input: candidates = [1], target = 1
Output: [[1]]

Example 5:
Input: candidates = [1], target = 2
Output: [[1,1]]
*/

package main

import (
	"fmt"
	"sort"
	"strconv"
)

func _combinatorial_sum(arr []int, res *[][]int, target, start int, temp []int, str string) {
	//fmt.Println(target, str, temp, start)
	if target < 0 {
		return
	}
	if target == 0 {
		*res = append(*res, append([]int{}, temp...))
		return
	}

	for i := start; i < len(arr); i++ {
		temp = append(temp, arr[i])
		_combinatorial_sum(arr, res, target-arr[i], i, temp, "i="+strconv.Itoa(i)+" 2nd call")
		temp = temp[:len(temp)-1]
	}

}

func combinatorial_sum(arr []int, target int) [][]int {
	sort.Ints(arr)
	res := [][]int{}
	_combinatorial_sum(arr, &res, target, 0, []int{}, "1st call")
	return res
}

func main() {
	fmt.Println(combinatorial_sum([]int{2, 3, 6, 7}, 7))
	fmt.Println(combinatorial_sum([]int{2, 3, 5}, 8))
	fmt.Println(combinatorial_sum([]int{2}, 1))
	fmt.Println(combinatorial_sum([]int{1}, 1))
	fmt.Println(combinatorial_sum([]int{1}, 2))
}