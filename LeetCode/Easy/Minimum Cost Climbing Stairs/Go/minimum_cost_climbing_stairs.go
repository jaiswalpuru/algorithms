package main

import (
	"fmt"
)

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

//----------------------brute force-------------------------
func min_cost_climbing_stairs(arr []int) int {
	return min(_recur(arr, 0), _recur(arr, 1))
}

func _recur(arr []int, ind int) int {
	if ind >= len(arr) {
		return 0
	}

	return arr[ind] + min(_recur(arr, ind+1), _recur(arr, ind+2))
}

//----------------------brute force-------------------------

//----------------------memoization-------------------------
func min_cost_climbing_stairs_memo(arr []int) int {
	dp := make([]int, len(arr))
	for i := 0; i < len(arr); i++ {
		dp[i] = -1
	}
	return min(_memo(arr, 0, &dp), _memo(arr, 1, &dp))
}

func _memo(arr []int, ind int, dp *[]int) int {
	if ind >= len(arr) {
		return 0
	}

	if (*dp)[ind] != -1 {
		return (*dp)[ind]
	}

	(*dp)[ind] = arr[ind] + min(_memo(arr, ind+1, dp), _memo(arr, ind+2, dp))
	return (*dp)[ind]
}

//----------------------memoization-------------------------

func main() {
	fmt.Println(min_cost_climbing_stairs_memo([]int{10, 15, 20}))
}
