package main

import "fmt"

//recursion
func climbing_stairs(n int) int {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return 1
	}

	return climbing_stairs(n-1) + climbing_stairs(n-2)
}

//memo
func climbing_stairs_memo(n int) int {
	if n <= 1 {
		return 1
	}
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	return _memo(n, &dp)
}

func _memo(n int, dp *[]int) int {
	if (*dp)[n] != 0 {
		return (*dp)[n]
	}
	(*dp)[n] = _memo(n-1, dp) + _memo(n-2, dp)
	return (*dp)[n]
}

//bottom up

func climbing_stairs_dp(n int) int {
	if n <= 1 {
		return 1
	}
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

func main() {
	fmt.Println(climbing_stairs_dp(3))
}
