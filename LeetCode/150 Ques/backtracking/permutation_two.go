/*
Given a collection of numbers, nums, that might contain duplicates,
return all possible unique permutations in any order.

Example 1:
Input: nums = [1,1,2]
Output:
[[1,1,2],
 [1,2,1],
 [2,1,1]]

Example 2:
Input: nums = [1,2,3]
Output: [[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
*/

package main

import "fmt"

func _permute(arr []int, res *[][]int, ind int) {
	if ind == len(arr) {
		*res = append(*res, append([]int{}, arr...))
		return
	}
	for i := ind; i < len(arr); i++ {
		if i > ind && arr[i] == arr[ind] {
			continue
		}
		arr[i], arr[ind] = arr[ind], arr[i]
		_permute(arr, res, ind+1)
		arr[i], arr[ind] = arr[ind], arr[i]
	}
}

func permute(arr []int) [][]int {
	res := [][]int{}
	_permute(arr, &res, 0)
	return res
}

func main() {
	fmt.Println(permute([]int{1, 1, 2}))
	fmt.Println(permute([]int{1, 2, 3}))
}
