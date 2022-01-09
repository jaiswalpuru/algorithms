/*
Given an array nums of distinct integers, return all the possible permutations. You can return the answer in any order.
*/

package main

import "fmt"

func permutations(arr []int) [][]int {
	res := [][]int{}
	mp := make(map[int]struct{})
	_permutations(arr, &res, []int{}, &mp)
	return res
}

func _permutations(arr []int, res *[][]int, set []int, visited *map[int]struct{}) {
	if len(set) == len(arr) {
		*res = append(*res, append([]int{}, set...))
		return
	}

	for i := 0; i < len(arr); i++ {
		if _, ok := (*visited)[i]; !ok {
			(*visited)[i] = struct{}{}
			temp := append(set, arr[i])
			_permutations(arr, res, temp, visited)
			delete(*visited, i)
			temp = temp[:len(temp)-1]
		}
	}
}

func main() {
	fmt.Println(permutations([]int{1, 2, 3}))
}
