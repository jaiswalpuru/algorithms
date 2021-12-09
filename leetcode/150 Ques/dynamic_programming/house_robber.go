/*
You are a professional robber planning to rob houses along a street.
Each house has a certain amount of money stashed, the only constraint stopping you
from robbing each of them is that adjacent houses have security systems connected
and it will automatically contact the police if two adjacent houses were broken into on the same night.

Given an integer array nums representing the amount of money of each house,
return the maximum amount of money you can rob tonight without alerting the police.

Example 1:
Input: nums = [1,2,3,1]
Output: 4
Explanation: Rob house 1 (money = 1) and then rob house 3 (money = 3).
Total amount you can rob = 1 + 3 = 4.

Example 2:
Input: nums = [2,7,9,3,1]
Output: 12
Explanation: Rob house 1 (money = 2), rob house 3 (money = 9) and rob house 5 (money = 1).
Total amount you can rob = 2 + 9 + 1 = 12.
*/

package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func robber(arr []int) int {
	n := len(arr)
	if n == 1 {
		return arr[0]
	}
	if n == 2 {
		return max(arr[0], arr[1])
	}
	dp := make([]int, n)

	dp[0] = arr[0]
	dp[1] = max(arr[0], arr[1])

	for i := 2; i < n; i++ {
		dp[i] = max(dp[i-1], arr[i]+dp[i-2])
	}

	return dp[n-1]
}

func _robber_recursion(arr []int, i int) int {
	if i < 0 {
		return 0
	}
	rob := arr[i] + _robber_recursion(arr, i-2)
	dnt_rob := _robber_recursion(arr, i-1)
	return max(rob, dnt_rob)
}

func robber_recursion(arr []int) int {
	return _robber_recursion(arr, len(arr)-1)
}

func main() {
	fmt.Println(robber_recursion([]int{1, 2, 3, 1}))
	fmt.Println(robber_recursion([]int{2, 7, 9, 3, 1}))
	fmt.Println(robber([]int{2, 1, 1, 2}))
	fmt.Println(robber([]int{1, 2, 1, 1}))
}
