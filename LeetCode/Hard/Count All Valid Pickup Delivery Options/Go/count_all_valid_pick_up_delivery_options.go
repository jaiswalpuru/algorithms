package main

import "fmt"

var mod = int(1e9 + 7)

//------------------------recurison (TLE)-----------------------------------
func count_all_valid_pickup_delivery_options_recursion(n int) int {
	return _count_all_valid_pickup_delivery_options_recursion(n, n)
}

func _count_all_valid_pickup_delivery_options_recursion(unpicked, undelivered int) int {
	if unpicked == 0 && undelivered == 0 {
		return 1
	}

	if undelivered < 0 || unpicked < 0 || undelivered < unpicked {
		return 0
	}

	ans := 0
	ans += unpicked * _count_all_valid_pickup_delivery_options_recursion(unpicked-1, undelivered)
	ans %= mod
	ans += (undelivered - unpicked) * _count_all_valid_pickup_delivery_options_recursion(unpicked, undelivered-1)
	ans %= mod
	return ans
}

//------------------------------------------------------------------------

//------------------------------------using memoization------------------------------------
func count_all_valid_pickup_delivery_options_memo(n int) int {
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, n+1)
		for j := 0; j <= n; j++ {
			dp[i][j] = -1
		}
	}

	return _count_all_valid_pickup_delivery_options_memo(n, n, &dp)
}

func _count_all_valid_pickup_delivery_options_memo(unpicked, undelivered int, dp *[][]int) int {
	if undelivered == 0 && unpicked == 0 {
		return 1
	}

	if undelivered < 0 || unpicked < 0 || undelivered < unpicked {
		return 0
	}

	if (*dp)[unpicked][undelivered] != -1 {
		return (*dp)[unpicked][undelivered]
	}

	count := 0
	count += unpicked * _count_all_valid_pickup_delivery_options_memo(unpicked-1, undelivered, dp)
	count %= mod
	count += (undelivered - unpicked) * _count_all_valid_pickup_delivery_options_memo(unpicked, undelivered-1, dp)
	count %= mod
	(*dp)[unpicked][undelivered] = count
	return (*dp)[unpicked][undelivered]
}

//------------------------------------------------------------------------

//-----------------------------using bottom up approach-------------------------------------
func count_all_valid_pickup_delivery_options_dp(n int) int {
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, n+1)
	}

	for unpic := 0; unpic <= n; unpic++ {
		for undev := unpic; undev <= n; undev++ {
			if unpic == 0 && undev == 0 {
				dp[unpic][undev] = 1
			} else {
				if unpic > 0 {
					dp[unpic][undev] += unpic * dp[unpic-1][undev]
					dp[unpic][undev] %= mod
				}
				if undev > unpic {
					dp[unpic][undev] += (undev - unpic) * dp[unpic][undev-1]
					dp[unpic][undev] %= mod
				}
			}
		}
	}

	return dp[n][n]
}

//-------------------------------------------------------------------------------------------

//----------------------------------using permutation and combination ---------------------------(Best approach)
func count_all_valid_pickup_delivery_options_permute(n int) int {
	ans := 1
	for i := 1; i <= n; i++ {
		ans *= i
		ans *= (2*i - 1)
		ans %= mod
	}
	return ans
}

//----------------------------------using permutation and combination ---------------------------

func main() {
	fmt.Println(count_all_valid_pickup_delivery_options_permute(3))
}
