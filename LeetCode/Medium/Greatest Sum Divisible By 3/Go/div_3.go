package main

import (
	"fmt"
	"math"
)

//-------------------------------Brute force-------------------------------
//generate all possible subsequnces and check their sum
func div_3(arr []int) int {
	res := 0
	_div_3(arr, 0, &res, 0)
	return res
}

func _div_3(arr []int, start int, res *int, sum int) {
	if start == len(arr) {
		if sum%3 == 0 {
			*res = max(*res, sum)
		}
		return
	}

	sum += arr[start]
	_div_3(arr, start+1, res, sum)
	sum -= arr[start]
	_div_3(arr, start+1, res, sum)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//---------------------------------------------------------------------------------

//-------------------------------Efficient DP solution-------------------------------
//keeping track of remainder at each state
func div_3_eff(arr []int) int {
	n := len(arr)
	dp := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = make([]int, 3)
	}
	dp[0][1] = int(math.MinInt32)
	dp[0][2] = int(math.MinInt32)

	for i := 1; i <= n; i++ {
		if arr[i-1]%3 == 0 {
			dp[i][0] = max(dp[i-1][0], dp[i-1][0]+arr[i-1])
			dp[i][1] = max(dp[i-1][1], dp[i-1][1]+arr[i-1])
			dp[i][2] = max(dp[i-1][2], dp[i-1][2]+arr[i-1])
		} else if arr[i-1]%3 == 1 {
			dp[i][0] = max(dp[i-1][0], dp[i-1][2]+arr[i-1])
			dp[i][1] = max(dp[i-1][1], dp[i-1][0]+arr[i-1])
			dp[i][2] = max(dp[i-1][2], dp[i-1][1]+arr[i-1])
		} else {
			dp[i][0] = max(dp[i-1][0], dp[i-1][1]+arr[i-1])
			dp[i][1] = max(dp[i-1][1], dp[i-1][2]+arr[i-1])
			dp[i][2] = max(dp[i-1][2], dp[i-1][0]+arr[i-1])
		}
	}
	return dp[n][0]
}

//--------------------------------------------------------------------------------------

func main() {
	fmt.Println(div_3_eff([]int{3, 6, 5, 1, 8}))
}
