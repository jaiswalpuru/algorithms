package main

import "fmt"

var mod = int(1e9 + 7)

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

//--------------------brute force ----------------
func k_inverse_pairs(n, k int) int {
	if n == 0 {
		return 0
	}

	if k == 0 {
		return 1
	}

	inv := 0
	for i := 0; i <= min(k, n-1); i++ {
		inv = (inv + k_inverse_pairs(n-1, k-i)) % mod
	}

	return inv
}

//--------------------brute force ----------------

//--------------------recursion with memoization still TLE----------------
func k_inverse_pairs_eff(n, k int) int {
	memo := make([][]int, 1001)
	for i := 0; i < 1001; i++ {
		memo[i] = make([]int, 1001)
		for j := 0; j < 1001; j++ {
			memo[i][j] = -1
		}
	}
	return _memo(n, k, &memo)
}

func _memo(n, k int, memo *[][]int) int {
	if n == 0 {
		return 0
	}
	if k == 0 {
		return 1
	}

	if (*memo)[n][k] != -1 {
		return (*memo)[n][k]
	}

	inv := 0
	for i := 0; i <= min(n-1, k); i++ {
		inv = (inv + _memo(n-1, k-i, memo)) % mod
	}
	(*memo)[n][k] = inv
	return (*memo)[n][k]
}

//--------------------recursion with memoization still TLE----------------

//---------------------dp with cumulative sum---------------------------
func k_inverse_pairs_dp(n, k int) int {
	dp := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = make([]int, k+1)
	}
	for i := 1; i <= n; i++ {
		for j := 0; j <= k; j++ {
			if j == 0 {
				dp[i][j] = 1
			} else {
				val := (dp[i-1][j] + mod) % mod
				if j-i >= 0 {
					val = (val - dp[i-1][j-i]) % mod
				} else {
					val %= mod
				}
				dp[i][j] = (dp[i][j-1] + val) % mod
			}
		}
	}
	res := dp[n][k] + mod
	if k > 0 {
		res = (res - dp[n][k-1]) % mod
	} else {
		res %= mod
	}
	return res
}

//---------------------dp with cumulative sum---------------------------

func main() {
	fmt.Println(k_inverse_pairs_dp(3, 0))
}
