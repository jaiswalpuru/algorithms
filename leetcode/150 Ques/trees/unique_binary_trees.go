/*
Given an integer n, return the number of structurally unique BST's (binary search trees) which has
exactly n nodes of unique values from 1 to n.

Example 1:
Input: n = 3
Output: 5

Example 2:
Input: n = 1
Output: 1
*/

package main

import "fmt"

func num_trees(n int) int {
	dp := make([]int, n+1)

	dp[0] = 1
	dp[1] = 1

	for i := 2; i <= n; i++ {
		for j := 1; j <= i; j++ {
			dp[i] = dp[i] + (dp[i-j] * dp[j-1])
		}
	}
	return dp[n]
}

func main() {
	fmt.Println(num_trees(5))
	fmt.Println(num_trees(1))
}
