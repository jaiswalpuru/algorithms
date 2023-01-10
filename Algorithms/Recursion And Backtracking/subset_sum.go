/*
Given a list of n integers print all the sums of subsets in it. Output should be printed in increasing order of sums
*/

package main

import (
	"fmt"
	"sort"
)

func subset_sum(arr []int) []int {
	res := []int{}
	_subset_sum(arr, &res, 0, 0)
	sort.Ints(res)
	return res
}

func _subset_sum(arr []int, res *[]int, start int, sum int) {
	if start >= len(arr) {
		*res = append(*res, sum)
		return
	}

	_subset_sum(arr, res, start+1, sum+arr[start])
	_subset_sum(arr, res, start+1, sum)

}

func main() {
	fmt.Println(subset_sum([]int{3, 1, 2}))
}
