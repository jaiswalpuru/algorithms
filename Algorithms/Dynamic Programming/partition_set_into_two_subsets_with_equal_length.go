/*
Partition the set into two subsets such that differnce between the sum of both the subsets is minimized
Concept used is subset sum question
Say if one subset sum is x then other subset sum will definitely be total_set_sum-x
*/

package main

import (
	"fmt"
	"math"
)

func subset_sum(arr []int, sum int) [][]bool {
	n := len(arr)
	dp := make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, sum+1)
		dp[i][0] = true
	}

	if arr[0] <= sum {
		dp[0][arr[0]] = true
	}

	for i := 1; i < n; i++ {
		for j := 1; j <= sum; j++ {
			not_take := dp[i-1][j]
			take := false
			if arr[i] <= j {
				take = dp[i-1][j-arr[i]]
			}
			dp[i][j] = take || not_take
		}
	}
	return dp
}

func subset_sum_with_equal_partition(arr []int) int {
	sum := 0
	n := len(arr)
	for i := 0; i < n; i++ {
		sum += arr[i]
	}

	min_val := math.MaxInt64
	dp := subset_sum(arr, sum)

	for i := 0; i <= sum/2; i++ {
		if dp[n-1][i] == true {
			min_val = min(min_val, abs((sum-i)-i))
		}
	}
	return min_val
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -1 * a
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	fmt.Println(subset_sum_with_equal_partition([]int{3, 2, 7}))
}
