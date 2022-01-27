package main

import (
	"fmt"
	"math"
)

//------------------------------brute force------------------------------
func max_subarray_after_one_operation_bf(arr []int) int {
	arr_temp := arr
	n := len(arr_temp)

	ans := math.MinInt64

	for i := 0; i < n; i++ {
		temp := arr_temp[i]
		arr_temp[i] = temp * temp
		ans = max(ans, max_sum(arr_temp))
		arr_temp[i] = temp
	}

	return ans
}

func max_sum(arr []int) int {
	n := len(arr)

	curr_max := arr[0]
	global_max := arr[0]
	for i := 1; i < n; i++ {
		curr_max = max(arr[i], arr[i]+curr_max)
		global_max = max(global_max, curr_max)
	}
	return global_max
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//------------------------------------------------------------------------

//-----------------------dp with memoization-------------------------------------------------
const (
	false = 0
	true  = 1
)

func max_subarray(arr []int) int {
	n := len(arr)

	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, 2)
		dp[i][0] = -1
		dp[i][1] = -1
	}
	res := 0
	memoization(arr, 0, 1, &res, &dp)
	return res
}

func memoization(arr []int, curr, consider int, res *int, dp *[][]int) int {
	if curr >= len(arr) {
		return 0
	}

	if (*dp)[curr][consider] > -1 {
		return (*dp)[curr][consider]
	}

	(*dp)[curr][false] = max(arr[curr], arr[curr]+memoization(arr, curr+1, false, res, dp))

	(*dp)[curr][true] = max(arr[curr]*arr[curr], max(arr[curr]+memoization(arr, curr+1, true, res, dp),
		arr[curr]*arr[curr]+memoization(arr, curr+1, false, res, dp)))

	*res = max(*res, (*dp)[curr][1])

	return (*dp)[curr][consider]
}

//----------------------------------------------------------------------------------------------

func dp_max_subarray_without_memo(arr []int) int {
	n := len(arr)
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, 2)
	}
	res := math.MinInt64
	for i := 0; i < n; i++ {
		dp[i+1][1] = max(dp[i][0]+arr[i]*arr[i], max(arr[i]*arr[i], dp[i][1]+arr[i]))
		dp[i+1][0] = max(dp[i][0]+arr[i], arr[i])
		res = max(res, max(dp[i+1][0], dp[i+1][1]))
	}

	return res
}

func main() {
	fmt.Println(dp_max_subarray_without_memo([]int{2, -1, -4, -3}))
}
