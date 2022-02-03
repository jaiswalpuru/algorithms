/*
You are given an array, of n positive integers and a target k, your task is to check if the there exist a subset in array
such that the sum equal k
*/

package main

import "fmt"

//---------------------recursion---------------------
func subset_sum(arr []int, target int) bool {
	n := len(arr)
	return _subset_sum(arr, target, n-1)
}

func _subset_sum(arr []int, target, ind int) bool {

	if target == 0 {
		return true
	}

	if ind == 0 {
		return target == arr[0]
	}

	not_take := _subset_sum(arr, target, ind-1)
	take := false
	if arr[ind] <= target {
		take = _subset_sum(arr, target-arr[ind], ind-1)
	}
	return not_take || take
}

//--------------------------------------------------------

//-----------------------memoization-----------------------
func subset_sum_memo(arr []int, target int) bool {
	n := len(arr)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, target+1)
		for j := 0; j <= target; j++ {
			dp[i][j] = -1
		}
	}

	return subset_memo(arr, &dp, target, n-1)
}

func subset_memo(arr []int, dp *[][]int, target, ind int) bool {
	if target == 0 {
		return true
	}

	if ind == 0 {
		return target == arr[0]
	}

	if (*dp)[ind][target] != -1 {
		if (*dp)[ind][target] == 0 {
			return true
		} else {
			return false
		}
	}

	not_take := subset_memo(arr, dp, target, ind-1)
	take := false
	if arr[ind] <= target {
		take = subset_memo(arr, dp, target-arr[ind], ind-1)
	}

	if take || not_take {
		(*dp)[ind][target] = 1
		return true
	} else {
		(*dp)[ind][target] = 0
		return false
	}
}

//-----------------------memoization-----------------------

//-----------------------dp-----------------------
func subset_sum_dp(arr []int, target int) bool {
	n := len(arr)
	dp := make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, target+1)
		dp[i][0] = true
	}

	for i := 1; i < n; i++ {
		for tgt := 1; tgt <= target; tgt++ {
			not_take := dp[i-1][tgt]
			take := false
			if arr[i] <= target {
				take = dp[i-1][target-arr[i]]
			}
			dp[i][tgt] = take || not_take
		}
	}
	return dp[n-1][target]
}

//-----------------------dp-----------------------

func main() {
	fmt.Println(subset_sum_dp([]int{1, 2, 4}, 4))
}
