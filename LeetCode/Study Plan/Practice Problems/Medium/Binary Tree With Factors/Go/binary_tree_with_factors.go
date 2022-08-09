package main

import (
	"fmt"
	"sort"
)

var mod = int(1e9 + 7)

func binary_tree_with_factors(arr []int) int {
	mp := make(map[int]int)
	n := len(arr)
	sort.Ints(arr)
	for i := 0; i < n; i++ {
		mp[arr[i]] = i
	}
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = 1
	}
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			if arr[i]%arr[j] == 0 {
				right := arr[i] / arr[j]
				if _, ok := mp[right]; ok {
					dp[i] = (dp[i] + (dp[j] * dp[mp[right]])) % mod
				}
			}
		}
	}
	res := 0
	for i := 0; i < n; i++ {
		res = (res + dp[i]) % mod
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(binary_tree_with_factors([]int{2, 4}))
}
