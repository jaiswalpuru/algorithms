/*
Count the number of subsets with sum equal to k
*/
package main

import "fmt"

func subset_k(arr []int, target int) int {
	return _count_sum_k_subsets(arr, target, len(arr)-1)
}

func _count_sum_k_subsets(arr []int, target int, ind int) int {
	if target == 0 {
		return 1
	}

	if ind == 0 {
		if arr[0] == target {
			return 1
		} else {
			return 0
		}
	}

	not_take := _count_sum_k_subsets(arr, target, ind-1)
	take := 0
	if arr[ind] <= target {
		take = _count_sum_k_subsets(arr, target-arr[ind], ind-1)
	}
	return not_take + take

}

func subset_k_memo(arr []int, target int) int {
	n := len(arr)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, target+1)
		for j := 0; j <= target; j++ {
			dp[i][j] = -1
		}
	}

	return _count_sum_k_subsets_memo(arr, &dp, target, n-1)
}

func _count_sum_k_subsets_memo(arr []int, dp *[][]int, target int, ind int) int {
	if target == 0 {
		return 1
	}
	if ind == 0 {
		if arr[0] == target {
			return 1
		} else {
			return 0
		}
	}

	if (*dp)[ind][target] != -1 {
		return (*dp)[ind][target]
	}

	not_take := _count_sum_k_subsets_memo(arr, dp, target, ind-1)
	take := 0
	if arr[ind] <= target {
		take = _count_sum_k_subsets_memo(arr, dp, target-arr[ind], ind-1)
	}
	(*dp)[ind][target] = take + not_take
	return (*dp)[ind][target]
}

func subset_k_dp(arr []int, target int) int {
	n := len(arr)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, target+1)
		dp[i][0] = 1
	}

	if arr[0] <= target {
		dp[0][arr[0]] = 1
	}

	for i := 1; i < n; i++ {
		for j := 1; j <= target; j++ {
			not_take := dp[i-1][j]
			take := 0
			if j >= arr[i] {
				take = dp[i-1][j-arr[j]]
			}
			dp[i][j] = take + not_take
		}
	}
	return dp[n-1][target]
}

func main() {
	fmt.Println(subset_k_dp([]int{1, 2, 2, 3}, 3))
}
