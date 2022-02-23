package main

import (
	"fmt"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func house_robber_II(arr []int) int {
	n := len(arr)
	return max(_house_robber_II(arr[:n-1], n-2), _house_robber_II(arr[1:], n-2))
}

func _house_robber_II(arr []int, start int) int {
	if start < 0 {
		return 0
	}
	if start == 0 {
		return arr[start]
	}

	rob := arr[start] + _house_robber_II(arr, start-2)
	dont_rob := _house_robber_II(arr, start-1)
	return max(rob, dont_rob)
}

func house_robber_II_memo(arr []int) int {
	n := len(arr)
	dp, dp_1 := make([]int, n-1), make([]int, n-1)
	for i := 0; i < n-1; i++ {
		dp[i] = -1
		dp_1[i] = -1
	}
	return max(_house_robber_II_memo(arr[1:], n-2, &dp), _house_robber_II_memo(arr[:n-1], n-2, &dp_1))
}

func _house_robber_II_memo(arr []int, start int, dp *[]int) int {
	if start < 0 {
		return 0
	}

	if start == 0 {
		return arr[start]
	}

	if (*dp)[start] != -1 {
		return (*dp)[start]
	}

	rob := _house_robber_II_memo(arr, start-2, dp) + arr[start]
	dont_rob := _house_robber_II_memo(arr, start-1, dp)
	(*dp)[start] = max(rob, dont_rob)
	return (*dp)[start]

}

func house_robber_dp(arr []int) int {
	n := len(arr)
	if n == 1 {
		return arr[0]
	}
	dp := make([]int, n-1)
	return max(_house_robber_dp(arr[1:], n-1, dp), _house_robber_dp(arr[:n-1], n-1, dp))
}

func _house_robber_dp(arr []int, end int, dp []int) int {
	if end == 1 {
		return arr[0]
	}
	dp[0] = arr[0]
	dp[1] = max(arr[0], arr[1])

	for i := 2; i < end; i++ {
		dont_rob := dp[i-1]
		rob := arr[i] + dp[i-2]
		dp[i] = max(rob, dont_rob)
	}
	return dp[end-1]
}

func main() {
	fmt.Println(house_robber_dp([]int{1, 3, 1, 3, 100}))
}
