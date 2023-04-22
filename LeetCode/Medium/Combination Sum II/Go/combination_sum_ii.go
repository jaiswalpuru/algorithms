package main

import (
	"fmt"
	"sort"
)

func combinationSum2(candidates []int, target int) [][]int {
	res := [][]int{}
	sort.Ints(candidates)
	backtrack(candidates, target, &res, []int{}, 0)
	return res
}

func backtrack(candidates []int, target int, res *[][]int, temp []int, start int) {
	if target == 0 {
		*res = append(*res, append([]int{}, temp...))
		return
	}
	for i := start; i < len(candidates); i++ {
		if i > start && candidates[i] == candidates[i-1] {
			continue
		}
		if candidates[i] > target {
			continue
		}
		temp = append(temp, candidates[i])
		backtrack(candidates, target-candidates[i], res, temp, i+1)
		temp = temp[:len(temp)-1]
	}
}

func main() {
	fmt.Println(combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8))
}
