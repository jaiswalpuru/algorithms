package main

import (
	"fmt"
	"math"
)

const (
	MIN = int(math.MinInt64)
)

//------------------------------Brute Force------------------------------
func target_sum(arr []int, target int) int {
	cnt := 0
	_target_sum(arr, &cnt, 0, target, 0)
	return cnt
}

//brute force try every possibility
func _target_sum(arr []int, res *int, sum, target int, start int) {
	if start == len(arr) {
		if target == sum {
			*res += 1
		}
		return
	}
	_target_sum(arr, res, sum+arr[start], target, start+1)
	_target_sum(arr, res, sum-arr[start], target, start+1)
}

//------------------------------------------------------------------------------------------

//------------------------------with memoziation------------------------------
func target_sum_dp(arr []int, target int) int {
	sum := 0
	n := len(arr)
	for i := 0; i < n; i++ {
		sum += arr[i]
	}

	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, 2*sum+1)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < 2*sum+1; j++ {
			dp[i][j] = MIN
		}
	}

	return memo_target_sum(arr, &dp, 0, 0, sum, target)
}

func memo_target_sum(arr []int, dp *[][]int, start, sum, total, target int) int {
	if start == len(arr) {
		if sum == target {
			return 1
		}
		return 0
	}
	if (*dp)[start][sum+total] != MIN {
		return (*dp)[start][sum+total]
	}
	add := memo_target_sum(arr, dp, start+1, sum+arr[start], total, target)
	subtract := memo_target_sum(arr, dp, start+1, sum-arr[start], total, target)
	(*dp)[start][sum+total] = add + subtract
	return (*dp)[start][sum+total]
}

//------------------------------------------------------------------------------------------

//---------------------------------DP 2d---------------------------------

func target_sum_dp_2(arr []int, target int) int {
	total := 0
	n := len(arr)
	for i := 0; i < n; i++ {
		total += arr[i]
	}

	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, 2*total+1)
	}
	dp[0][arr[0]+total] = 1
	dp[0][-arr[0]+total] = 1

	for i := 1; i < n; i++ {
		for sum := -total; sum <= total; sum++ {
			if dp[i-1][sum+total] > 0 {
				dp[i][arr[i]+sum+total] += dp[i-1][sum+total]
				dp[i][-arr[i]+sum+total] += dp[i-1][sum+total]
			}
		}
	}
	if int(math.Abs(float64(target))) > total {
		return 0
	}
	return dp[len(dp)-1][total+target]

}

//------------------------------------------------------------------

func main() {
	fmt.Println(target_sum_dp_2([]int{2, 1}, 3))
}
