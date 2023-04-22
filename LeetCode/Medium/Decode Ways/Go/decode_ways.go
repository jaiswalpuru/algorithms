package main

import (
	"fmt"
	"strconv"
)

//--------------------------------brute force bad one don't do--------------------------------
func decode_ways(str string) int {
	n := len(str)
	return _decode_ways(str, 0, n)
}

func _decode_ways(str string, ind int, n int) int {
	if ind == n {
		return 1
	}

	if str[ind] == '0' {
		return 0
	}

	not_take := _decode_ways(str, ind+1, n)
	take := 0
	if ind+2 <= n && str[ind] != '0' {
		val, _ := strconv.Atoi(str[ind : ind+2])
		if val <= 26 {
			take = _decode_ways(str, ind+2, n)
		}
	}

	return take + not_take

}

//----------------------------------------------------------------------------

//--------------------------------memoization--------------------------------
func decode_ways_memo(str string) int {
	n := len(str)
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = -1
	}
	return _decode_ways_memo(str, 0, n, &dp)
}

func _decode_ways_memo(str string, ind int, n int, dp *[]int) int {
	if ind == n {
		return 1
	}

	if str[ind] == '0' {
		return 0
	}

	if (*dp)[ind] != -1 {
		return (*dp)[ind]
	}

	not_take := _decode_ways_memo(str, ind+1, n, dp)
	take := 0
	if ind+2 <= n && str[ind] != '0' {
		val, _ := strconv.Atoi(str[ind : ind+2])
		if val <= 26 {
			take = _decode_ways_memo(str, ind+2, n, dp)
		}
	}
	(*dp)[ind] = not_take + take
	return (*dp)[ind]
}

//---------------------------------------------------------------------------

//----------------------------------dp bottom up-----------------------------
func decode_ways_dp(str string) int {
	n := len(str)
	dp := make([]int, n+1)

	dp[n] = 1

	for i := n - 1; i >= 0; i-- {
		if str[i] == '0' {
			dp[i] = 0
		} else {
			not_take := dp[i+1]
			take := 0
			if i+2 <= n && str[i] != '0' {
				val, _ := strconv.Atoi(str[i : i+2])
				if val <= 26 {
					take = dp[i+2]
				}
			}
			dp[i] = not_take + take
		}
	}
	return dp[0]
}

//----------------------------------dp bottom up-----------------------------

func main() {
	fmt.Println(decode_ways_dp("111111111111111111111111111111111111111111111"))
}
