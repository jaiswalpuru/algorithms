package main

import "fmt"

func permuteUnique(nums []int) [][]int {
	res := [][]int{}
	backtrack(&nums, &res, 0)
	return res
}

func backtrack(nums *[]int, res *[][]int, start int) {
	if start == len(*nums) {
		*res = append(*res, append([]int{}, *nums...))
		return
	}
	visited := make(map[int]bool)
	for i := start; i < len(*nums); i++ {
		if visited[(*nums)[i]] {
			continue
		}
		(*nums)[i], (*nums)[start] = (*nums)[start], (*nums)[i]
		backtrack(nums, res, start+1)
		(*nums)[i], (*nums)[start] = (*nums)[start], (*nums)[i]
		visited[(*nums)[i]] = true
	}
}

func main() {
	fmt.Println(permuteUnique([]int{1, 1, 2}))
}
