package main

import (
	"fmt"
	"math"
	"sort"
)

func minimum_diff_between_highest_and_lowest_k_Scores(nums []int, k int) int {
	sort.Ints(nums)
	n := len(nums)
	res := math.MaxInt64
	for i := 0; i <= n-k; i++ {
		res = min(res, nums[i+k-1]-nums[i])
	}
	return res
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	fmt.Println(minimum_diff_between_highest_and_lowest_k_Scores([]int{90}, 1))
}
