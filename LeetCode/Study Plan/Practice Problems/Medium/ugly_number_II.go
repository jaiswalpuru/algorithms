/*
An ugly number is a positive integer whose prime factors are limited to 2, 3, and 5.

Given an integer n, return the nth ugly number.
*/

package main

import (
	"fmt"
)

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func ugly_number(n int) int {
	if n == 1 {
		return 1
	}
	dp := make([]int, n)
	dp[0] = 1
	idx_2, idx_3, idx_5 := 0, 0, 0
	count := 1
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

func main() {
	fmt.Println(ugly_number(10))
}
