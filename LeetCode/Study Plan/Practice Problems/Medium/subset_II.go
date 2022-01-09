/*
Given an integer array nums that may contain duplicates, return all possible subsets (the power set).

The solution set must not contain duplicate subsets. Return the solution in any order.
*/

package main

import (
	"fmt"
	"sort"
)

func subset(arr []int) [][]int {
	res := [][]int{}
	sort.Ints(arr)
	_subset(arr, &res, []int{}, 0)
	return res
}

func _subset(arr []int, res *[][]int, set []int, start int) {
	*res = append(*res, append([]int{}, set...))

	for i := start; i < len(arr); i++ {
		if i > start && arr[i] == arr[i-1] {
			continue
		}
		temp := append(set, arr[i])
		_subset(arr, res, temp, i+1)
		temp = temp[:len(temp)-1]
	}
}

func main() {
	fmt.Println(subset([]int{1, 2, 2}))
}
