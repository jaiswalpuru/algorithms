/*
Given an integer array nums of unique elements, return all possible subsets (the power set).

The solution set must not contain duplicate subsets. Return the solution in any order.
*/

package main

import "fmt"

func subset(arr []int) [][]int {
	res := [][]int{}
	_subset(arr, &res, []int{}, 0)
	return res
}

func _subset(arr []int, res *[][]int, set []int, start int) {
	if start >= len(arr) {
		*res = append(*res, append([]int{}, set...))
		return
	}

	temp := append(set, arr[start])
	_subset(arr, res, temp, start+1)
	temp = temp[:len(temp)-1]
	_subset(arr, res, temp, start+1)
}

func main() {
	fmt.Println(subset([]int{1, 2, 3}))
}
