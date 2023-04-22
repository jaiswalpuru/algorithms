/*
You are given an m x n integer matrix mat and an integer target.

Choose one integer from each row in the matrix such that the absolute difference between target and the
sum of the chosen elements is minimized.

Return the minimum absolute difference.

The absolute difference between two numbers a and b is the absolute value of a - b.
*/

package main

import (
	"fmt"
	"math"
)

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func minimize(arr [][]int, target int) int {
	dp := make([][]int, 5001)
	for i := 0; i < 5001; i++ {
		dp[i] = make([]int, 71)
		for j := 0; j < 71; j++ {
			dp[i][j] = -1
		}
	}

	return mini(arr, 0, 0, &target, &dp)
}

func mini(arr [][]int, i, sum int, target *int, dp *[][]int) int {
	if i >= len(arr) {
		return int(math.Abs(float64(*target - sum)))
	}

	if (*dp)[sum][i] != -1 {
		return (*dp)[sum][i]
	}

	ans := int(math.MaxInt64)
	for j := 0; j < len((arr)[i]); j++ {
		ans = min(ans, mini(arr, i+1, sum+arr[i][j], target, dp))
	}
	(*dp)[sum][i] = ans
	return ans
}

func main() {
	fmt.Println(minimize([][]int{
		{1, 2, 3}, {4, 5, 6}, {7, 8, 9},
	}, 13))

	// fmt.Println(minimize([][]int{
	// 	{1}, {2}, {3},
	// }, 100))

	// fmt.Println(minimize([][]int{
	// 	{1, 2, 9, 8, 7},
	// }, 6))
}
