/*
Given a non-empty array nums containing only positive integers,
find if the array can be partitioned into two subsets such that the sum of elements in both subsets is equal.

*/
package main

import "fmt"

//--------------------------------------------bottom up--------------------------------------------

func partition_bu(arr []int) bool {
	n := len(arr)
	sum := 0
	for i := 0; i < n; i++ {
		sum += arr[i]
	}

	if sum%2 != 0 {
		return false
	}
	subset_sum := sum / 2
	dp := make([][]bool, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]bool, subset_sum+1)
	}

	dp[0][0] = true
	for i := 1; i <= n; i++ {
		curr := arr[i]
		for j := 1; j <= subset_sum; j++ {
			if j < curr {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j] || dp[i-1][j-curr]
			}
		}
	}
	return dp[n][subset_sum]
}

//----------------------------------------------------------------------------------------

//--------------------------------------------Using memoization--------------------------------------------
func _partition_dfs(arr []int, dp *[][]int, n, sum int) bool {
	if sum == 0 {
		return true
	}
	if n == 0 || sum < 0 {
		return false
	}

	if (*dp)[n][sum] != -1 {
		if (*dp)[n][sum] == 0 {
			return false
		} else {
			return true
		}
	}
	res := _partition_dfs(arr, dp, n-1, sum-arr[n]) || _partition_dfs(arr, dp, n-1, sum)
	if res == false {
		(*dp)[n][sum] = 0
	} else {
		(*dp)[n][sum] = 1
	}
	return res
}

func partition(arr []int) bool {
	total_sum := 0
	sum := 0
	n := len(arr)
	for i := 0; i < n; i++ {
		total_sum += arr[i]
	}

	if total_sum%2 != 0 {
		return false
	}
	sum = total_sum / 2
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, sum+1)
		for j := 0; j <= sum; j++ {
			dp[i][j] = -1
		}
	}
	return _partition_dfs(arr, &dp, n-1, sum)

}

//--------------------------------------------------------------------------------------------------------------

//-------------------------------------------------Approach 1 Brute force-----------------------------------------
//we can use dfs for this approach generate all possible subsets of given sum
func partition_bf(arr []int) bool {
	total_sum := 0
	n := len(arr)
	sum := 0
	for i := 0; i < n; i++ {
		total_sum += arr[i]
	}

	if total_sum%2 != 0 {
		return false
	}

	sum = total_sum / 2
	return dfs(arr, len(arr)-1, sum)
}

func dfs(arr []int, n int, sum int) bool {
	if sum == 0 {
		return true
	}
	if sum < 0 || n == 0 {
		return false
	}
	fmt.Println(arr[n-1])
	res := dfs(arr, n-1, sum-arr[n]) || dfs(arr, n-1, sum)
	return res
}

//--------------------------------------------------------------------------------------------------------------

func main() {
	fmt.Println(partition([]int{1, 5, 11, 5}))
}
