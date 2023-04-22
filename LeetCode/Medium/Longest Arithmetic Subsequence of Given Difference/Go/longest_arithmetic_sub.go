package main

import (
	"fmt"
	"math"
)

//------------------Brute force-----------------------(won't work TLE)
func longest_sub(arr []int, diff int) int {
	n := len(arr)
	res := int(math.MinInt64)
	for i := 0; i < n; i++ {
		val := arr[i]
		l := 1
		for j := i + 1; j < n; j++ {
			if arr[j]-val == diff {
				val = arr[j]
				l++
			}
		}
		res = max(res, l)
	}
	return res
}

//---------------------------------------------------------------------

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func longest_sub_eff(arr []int, diff int) int {
	dp := make(map[int]int)

	res := 1
	for i := range arr {
		dp[arr[i]] = dp[arr[i]-diff] + 1
		res = max(res, dp[arr[i]])
	}
	fmt.Println(dp)
	return res
}

func main() {
	fmt.Println(longest_sub_eff([]int{3, 4, -3, -2, -4}, -5))
}
