package main

import (
	"fmt"
)

//------------------------------------------------backtracking------------------------------------------------
func house_robber(arr []int) int {
	return _house_robber(arr, len(arr)-1)
}

func _house_robber(arr []int, ind int) int {
	if ind < 0 {
		return 0
	}
	if ind == 0 {
		return arr[ind]
	}

	pick := arr[ind] + _house_robber(arr, ind-2)
	not_pick := _house_robber(arr, ind-1)

	return max(pick, not_pick)
}

//------------------------------------------------------------------------------------------------

//dp with memoization
func house_robber_memo(arr []int) int {
	n := len(arr)
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = -1
	}

	res := _house_robber_(arr, &dp, n-1)
	fmt.Println(dp)
	return res
}

func _house_robber_(arr []int, dp *[]int, ind int) int {

	if ind < 0 {
		return 0
	}

	if ind == 0 {
		return arr[ind]
	}

	if (*dp)[ind] > -1 {
		return (*dp)[ind]
	}
	pick := arr[ind] + _house_robber_(arr, dp, ind-2)
	not_pick := _house_robber_(arr, dp, ind-1)

	(*dp)[ind] = max(pick, not_pick)
	return (*dp)[ind]
}

//dp with memo end

//dp bottom up
func house_robber_dp(arr []int) int {
	n := len(arr)
	dp := make([]int, n)

	dp[0] = arr[0]

	var take, dnt_take int
	for i := 1; i < n; i++ {
		take = arr[i]
		if i > 1 {
			take += dp[i-2]
		}
		dnt_take = dp[i-1]
		dp[i] = max(take, dnt_take)
	}
	return dp[n-1]
}

//dp bottom up end

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(house_robber_dp([]int{1, 2}))
}
