package main

import "fmt"

//----------------------recursion---------------------------
func partition_equal_subset_sum(arr []int) bool {
	sum := 0
	n := len(arr)
	for i := 0; i < n; i++ {
		sum += arr[i]
	}

	if sum%2 != 0 {
		return false // cannot partition odd sum
	}

	return _partition(arr, sum/2, n-1)
}

func _partition(arr []int, target int, ind int) bool {
	if target == 0 {
		return true
	}

	if ind == 0 {
		return arr[0] == target
	}

	not_take := _partition(arr, target, ind-1)
	take := false
	if arr[ind] <= target {
		take = _partition(arr, target-arr[ind], ind-1)
	}
	return take || not_take
}

//------------------------recursion---------------------------

//------------------------memoization---------------------------
func partition_equal_subset_sum_memo(arr []int) bool {
	sum := 0
	n := len(arr)
	for i := 0; i < n; i++ {
		sum += arr[i]
	}

	if sum%2 != 0 {
		return false
	}

	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, sum/2+1)
		for j := 0; j <= sum/2; j++ {
			dp[i][j] = -1
		}
	}

	return _partition_memo(arr, &dp, n-1, sum/2)
}

func _partition_memo(arr []int, dp *[][]int, ind int, target int) bool {
	if target == 0 {
		return true
	}

	if ind == 0 {
		return target == arr[0]
	}

	if (*dp)[ind][target] != -1 {
		if (*dp)[ind][target] == 0 {
			return false
		}
		return true
	}

	not_take := _partition_memo(arr, dp, ind-1, target)
	take := false
	if arr[ind] <= target {
		take = _partition_memo(arr, dp, ind-1, target-arr[ind])
	}
	if take || not_take {
		(*dp)[ind][target] = 1
		return true
	} else {
		(*dp)[ind][target] = 0
		return false
	}
}

//------------------------memoization---------------------------

//------------------------dp---------------------------
func partition_dp(arr []int) bool {
	n := len(arr)
	sum := 0

	for i := 0; i < n; i++ {
		sum += arr[i]
	}

	if sum%2 != 0 {
		return false
	}
	dp := make([][]bool, n)
	k := sum / 2
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, k+1)
		dp[i][0] = true
	}
	if arr[0] <= k {
		dp[0][arr[0]] = true
	}
	for i := 1; i < n; i++ {
		for j := 1; j <= k; j++ {
			not_take := dp[i-1][j]
			take := false
			if arr[i] <= j {
				take = dp[i-1][j-arr[i]]
			}
			dp[i][j] = take || not_take
		}
	}
	return dp[n-1][k]
}

//-------------------------dp---------------------------

func main() {
	fmt.Println(partition_dp([]int{1, 5, 11, 5}))
}
