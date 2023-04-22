package main

import (
	"fmt"
	"sort"
)

func subsetsWithDup(nums []int) [][]int {
	res := [][]int{}
	sort.Ints(nums)
	backtrack(nums, 0, &res, []int{})
	return res
}

func backtrack(nums []int, start int, res *[][]int, temp []int) {
	*res = append(*res, append([]int{}, temp...))
	for i := start; i < len(nums); i++ {
		if i > start && nums[i] == nums[i-1] {
			continue
		}
		temp = append(temp, nums[i])
		backtrack(nums, i+1, res, temp)
		temp = temp[:len(temp)-1]
	}
}

func main() {
	fmt.Println(subsetsWithDup([]int{1, 2, 2}))
}
