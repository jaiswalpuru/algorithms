package main

import (
	"fmt"
	"strconv"
)

func nonDecreasingSubsequences(nums []int) [][]int {
	res := [][]int{}
	visited := make(map[string]bool)
	recur(nums, 0, []int{}, &res, &visited)
	return res
}

func toString(nums []int) string {
	res := ""
	for i := 0; i < len(nums); i++ {
		if i == len(nums)-1 {
			res += strconv.Itoa(nums[i])
			continue
		}
		res += strconv.Itoa(nums[i]) + "->"
	}
	return res
}

func recur(nums []int, ind int, set []int, res *[][]int, visited *map[string]bool) {
	if ind == len(nums) {
		str := toString(set)
		if len(set) > 1 && !(*visited)[str] {
			(*visited)[str] = true
			*res = append(*res, append([]int{}, set...))
		}
		return
	}
	recur(nums, ind+1, set, res, visited)
	if len(set) == 0 || nums[ind] >= set[len(set)-1] {
		set = append(set, nums[ind])
		recur(nums, ind+1, set, res, visited)
		set = set[:len(set)-1]
	}
}

func main() {
	fmt.Println(nonDecreasingSubsequences([]int{4, 6, 7, 7}))
}
