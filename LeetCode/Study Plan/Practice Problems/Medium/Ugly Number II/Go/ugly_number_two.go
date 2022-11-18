package main

import "fmt"

//----------brute force---------------
func is_ugly(n int) bool {
	if n == 0 {
		return false
	}
	for n%2 == 0 {
		n = n / 2
	}
	for n%3 == 0 {
		n = n / 3
	}
	for n%5 == 0 {
		n = n / 5
	}
	return n == 1
}

func ugly_number_two(n int) int {
	cnt := 1
	i := 1
	for cnt <= n {
		if is_ugly(i) {
			cnt++
		}
		i++
	}
	return i - 1
}

//----------brute force---------------

//----------efficient approach---------------
func ugly_number_two_eff(n int) int {
	dp := make([]int, n)
	dp[0] = 1
	count := 1
	idx_2, idx_3, idx_5 := 0, 0, 0
	for count < n {
		v1, v2, v3 := dp[idx_2]*2, dp[idx_3]*3, dp[idx_5]*5
		val := min(v1, min(v2, v3))
		dp[count] = val
		count++
		if val == v1 {
			idx_2++
		}
		if val == v2 {
			idx_3++
		}
		if val == v3 {
			idx_5++
		}
	}
	return dp[n-1]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

//----------efficientapproach---------------

func main() {
	fmt.Println(ugly_number_two_eff(10))
}
