/*
Given an integer array nums of unique elements, return all possible subsets (the power set).

The solution set must not contain duplicate subsets. Return the solution in any order.
*/

package main

import "fmt"

func subsets(arr []int) [][]int {
	res := [][]int{}
	_subset(arr, []int{}, 0, &res)
	return res
}

func _subset(arr []int, subset []int, i int, res *[][]int) {
	*res = append(*res, append([]int{}, subset...))
	for j := i; j < len(arr); j++ {
		temp := append(subset, arr[j])
		_subset(arr, temp, j+1, res)
		temp = temp[:len(temp)-1]
	}
}

func main() {
	fmt.Println(subsets([]int{1, 2, 3}))
}
