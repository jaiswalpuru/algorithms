package main

import "fmt"

func combinationSum(candidates []int, target int) [][]int {
	res := [][]int{}
	backtrack(candidates, target, 0, &res, []int{})
	return res
}

func backtrack(candidates []int, target, currIndex int, res *[][]int, temp []int) {
	if target <= 0 {
		if target == 0 {
			*res = append(*res, append([]int{}, temp...))
		}
		return
	}
	for i := currIndex; i < len(candidates); i++ {
		if candidates[i] > target {
			continue
		}
		temp = append(temp, candidates[i])
		backtrack(candidates, target-candidates[i], i, res, temp)
		temp = temp[:len(temp)-1]
	}
}

func main() {
	fmt.Println(combinationSum([]int{2, 3, 6, 7}, 7))
}
