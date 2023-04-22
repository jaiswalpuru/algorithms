package main

import "fmt"

func number_of_submatrices_that_sum_to_target(arr [][]int, target int) int {
	n, m := len(arr), len(arr[0])
	pre_sum := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		pre_sum[i] = make([]int, m+1)
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			pre_sum[i][j] = pre_sum[i-1][j] + pre_sum[i][j-1] - pre_sum[i-1][j-1] + arr[i-1][j-1]
		}
	}
	cnt, curr_sum := 0, 0
	for i := 1; i <= n; i++ {
		for j := i; j <= n; j++ {
			mp := make(map[int]int)
			mp[0] = 1
			for col := 1; col <= m; col++ {
				curr_sum = pre_sum[j][col] - pre_sum[i-1][col]
				cnt += mp[curr_sum-target]
				mp[curr_sum]++
			}
		}
	}
	return cnt
}

func main() {
	fmt.Println(number_of_submatrices_that_sum_to_target([][]int{
		{0, 1, 0}, {1, 1, 1}, {0, 1, 0},
	}, 0))
}
