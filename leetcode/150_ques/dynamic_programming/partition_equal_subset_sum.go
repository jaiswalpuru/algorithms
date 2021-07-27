/*
Given a non-empty array nums containing only positive integers,
find if the array can be partitioned into two subsets such that the sum of elements in both subsets is equal.

Example 1:
Input: nums = [1,5,11,5]
Output: true
Explanation: The array can be partitioned as [1, 5, 5] and [11].

Example 2:
Input: nums = [1,2,3,5]
Output: false
Explanation: The array cannot be partitioned into equal sum subsets.
*/

package main

import "fmt"

func _partition_sum(arr []int, total int) int {
	n := len(arr)
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, total+1)
	}

	for i := 0; i <= n; i++ {
		dp[i][0] = 0
	}
	for i := 1; i <= total; i++ {
		dp[0][i] = 0
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= total; j++ {
			if arr[i-1] > j {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = max(arr[i-1]+dp[i-1][j-arr[i-1]], dp[i-1][j])
			}
		}
	}
	for i := 0; i < len(dp); i++ {
		fmt.Println(dp[i])
	}
	return dp[n][total]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func partition_sum(arr []int) bool {
	n := len(arr)
	sum := 0
	for i := 0; i < n; i++ {
		sum += arr[i]
	}

	if sum%2 != 0 {
		return false
	}
	val := _partition_sum(arr, sum/2)
	if val == sum/2 {
		return true
	}
	return false

}

func main() {
	fmt.Println(partition_sum([]int{1, 5, 11, 5}))
	fmt.Println(partition_sum([]int{1, 2, 3, 5}))
}
