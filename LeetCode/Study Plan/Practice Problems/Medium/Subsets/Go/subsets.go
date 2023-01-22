package main

import "fmt"

func subsets(nums []int) [][]int {
	res := [][]int{}
	backtrack(nums, 0, &res, []int{})
	return res
}

func backtrack(nums []int, ind int, res *[][]int, set []int) {
	if ind >= len(nums) {
		*res = append(*res, append([]int{}, set...))
		return
	}
	temp := append(set, nums[ind])
	backtrack(nums, ind+1, res, temp)
	temp = temp[:len(temp)-1]
	backtrack(nums, ind+1, res, temp)
}

func main() {
	fmt.Println(subsets([]int{1, 2, 3}))
}
