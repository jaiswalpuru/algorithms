/*
Given a number in the form of a list of digits, return all possible permutations.

For example, given [1,2,3], return [[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]].
*/

package main

import "fmt"

func _permutation(arr []int, res *[][]int, start int) {

	if start == len(arr) {
		*res = append(*res, append([]int{}, arr...))
		return
	}

	for i := start; i < len(arr); i++ {
		arr[i], arr[start] = arr[start], arr[i]
		_permutation(arr, res, start+1)
		arr[i], arr[start] = arr[start], arr[i]
	}
}

func permutations(arr []int) [][]int {
	res := [][]int{}
	_permutation(arr, &res, 0)
	return res
}

func main() {
	fmt.Println(permutations([]int{1, 2, 3}))
}
