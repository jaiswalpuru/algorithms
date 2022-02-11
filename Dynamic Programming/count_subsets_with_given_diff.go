/*
Given an array partition it into two subsets(possibly empty) such that their union is the original array.
Let the sum of the elements of the two subsets be S1 and S2.
Given the difference D, count the number of partitions in which S1>=S2 and the difference between them is D.

S1 - S2 = D (i)

S1 can also be represented as S1 = TotalSum - S2
replacing the value of S1 in (i)

TotalSum - S2 - S2 = D
TotalSum - 2 S2 = D

TotalSum - D = 2 * S2
S2 = (TotalSum - D)/2
*/

package main

import "fmt"

const (
	mod = 1000000007
)

//------------------------------using recursion-------------------------------
func partition_with_given_difference(arr []int, diff int) int {
	//explanation is given above we need to find the modified target

	sum := 0
	n := len(arr)
	for i := 0; i < n; i++ {
		sum += arr[i]
	}

	target := (sum - diff) / 2
	return _partition(arr, target, n-1)
}

func _partition(arr []int, target int, ind int) int {

	if ind == 0 {
		if target == 0 && arr[ind] == 0 {
			return 2
		}
		if target == 0 || target == arr[ind] {
			return 1
		}
		return 0
	}

	not_take := _partition(arr, target, ind-1)
	take := 0
	if arr[ind] <= target {
		take = _partition(arr, target-arr[ind], ind-1)
	}
	return (not_take + take) % mod
}

//------------------------------using recursion-------------------------------

//------------------------------using memoization-------------------------------
func partition_memo(arr []int, diff int) int {
	n := len(arr)

	sum := 0
	for i := 0; i < n; i++ {
		sum += arr[i]
	}
	target := (sum - diff) / 2

	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, target+1)
		for j := 0; j <= target; j++ {
			dp[i][j] = -1
		}
	}

	return _partition_memo(arr, target, n-1, &dp)
}

func _partition_memo(arr []int, target int, ind int, dp *[][]int) int {
	if ind == 0 {
		if target == 0 && arr[ind] == 0 {
			return 2
		}
		if target == 0 || arr[ind] == target {
			return 1
		}
		return 0
	}
	if (*dp)[ind][target] != -1 {
		return (*dp)[ind][target]
	}

	not_take := _partition_memo(arr, target, ind-1, dp)
	take := 0
	if arr[ind] <= target {
		take = _partition_memo(arr, target-arr[ind], ind-1, dp)
	}
	(*dp)[ind][target] = not_take + take
	return (*dp)[ind][target]
}

//------------------------------using memoization-------------------------------

//------------------------------using dp bottom up-------------------------------
func partition_dp(arr []int, diff int) int {
	n := len(arr)
	sum := 0
	for i := 0; i < n; i++ {
		sum += arr[i]
	}

	target := (sum - diff) / 2
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, target+1)
	}

	if arr[0] == 0 {
		dp[0][0] = 2
	} else {
		dp[0][0] = 1
	}

	if arr[0] <= target {
		dp[0][arr[0]] = 1
	}

	for ind := 1; ind < n; ind++ {
		for sum := 0; sum <= target; sum++ {
			not_take := dp[ind-1][sum]
			take := 0
			if arr[ind] <= sum {
				take = dp[ind-1][sum-arr[ind]]
			}
			dp[ind][sum] = not_take + take
		}
	}
	return dp[n-1][target]
}

//------------------------------using dp bottom up------------------------------

func main() {
	fmt.Println(partition_dp([]int{1, 1, 1, 1}, 0))
}
