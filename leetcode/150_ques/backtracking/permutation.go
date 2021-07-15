/*
Given an array nums of distinct integers, return all the possible permutations.
You can return the answer in any order.

Example 1:
Input: nums = [1,2,3]
Output: [[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]

Example 2:
Input: nums = [0,1]
Output: [[0,1],[1,0]]

Example 3:
Input: nums = [1]
Output: [[1]]
*/

package main

import "fmt"

func _permute(arr []int, res *[][]int, ind int) {

	if ind == len(arr) {
		*res = append(*res, append([]int{}, arr...))
	}

	for i := ind; i < len(arr); i++ {
		arr[i], arr[ind] = arr[ind], arr[i]
		_permute(arr, res, ind+1)
		arr[i], arr[ind] = arr[ind], arr[i]
	}

}

func permutation(arr []int) [][]int {
	res := [][]int{}
	_permute(arr, &res, 0)
	return res
}

func main() {
	fmt.Println(permutation([]int{1, 2, 3}))
	fmt.Println(permutation([]int{0, 1}))
	fmt.Println(permutation([]int{1}))
}
