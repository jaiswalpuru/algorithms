/*
Given an integer array nums that may contain duplicates, return all possible subsets (the power set).

The solution set must not contain duplicate subsets. Return the solution in any order.
*/

package main

import (
	"fmt"
	"sort"
)

func subsets_II(arr []int) [][]int {
	res := make([][]int, 0)
	sort.Ints(arr)
	_subsets_II(arr, []int{}, 0, &res)
	return res
}

func _subsets_II(arr []int, subsets []int, i int, res *[][]int) {
	*res = append(*res, append([]int{}, subsets...))
	for j := i; j < len(arr); j++ {
		if j > i && arr[j] == arr[j-1] {
			continue
		}
		temp := append(subsets, arr[j])
		_subsets_II(arr, temp, j+1, res)
		temp = temp[:len(temp)-1]
	}
}

func main() {
	fmt.Println(subsets_II([]int{1, 2, 2}))
	fmt.Println(subsets_II([]int{0}))
}
