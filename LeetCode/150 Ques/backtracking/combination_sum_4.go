/*
Given an array of distinct integers nums and a target integer target, return the number of possible combinations that add up to target.

The answer is guaranteed to fit in a 32-bit integer.

Example 1:
Input: nums = [1,2,3], target = 4
Output: 7
Explanation:
The possible combination ways are:
(1, 1, 1, 1)
(1, 1, 2)
(1, 2, 1)
(1, 3)
(2, 1, 1)
(2, 2)
(3, 1)
Note that different sequences are counted as different combinations.

Example 2:
Input: nums = [9], target = 3
Output: 0
*/

package main

import "fmt"

func _combination_sum_4(arr []int, res *[][]int, temp []int, target, ind int) {

	if target < 0 {
		return
	}

	if target == 0 {
		*res = append(*res, append([]int{}, temp...))
		return
	}

	for i := 0; i < len(arr); i++ {
		temp = append(temp, arr[i])
		_combination_sum_4(arr, res, temp, target-arr[i], i+1)
		temp = temp[:len(temp)-1]
	}

}

func combination_sum_4(arr []int, target int) [][]int {
	res := [][]int{}
	temp := []int{}
	_combination_sum_4(arr, &res, temp, target, 0)
	return res
}

func main() {
	fmt.Println(combination_sum_4([]int{1, 2, 3}, 4))
	fmt.Println(combination_sum_4([]int{9}, 3))
}
