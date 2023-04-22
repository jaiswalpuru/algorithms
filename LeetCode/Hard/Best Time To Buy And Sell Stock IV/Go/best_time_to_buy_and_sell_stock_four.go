package main

import (
	"fmt"
	"math"
)

//-------------------------memoization-----------------------------
func best_time_to_buy_and_sell_stock_four(k int, prices []int) int {
	memo := make([][][]int, len(prices))
	for i := 0; i < len(memo); i++ {
		memo[i] = make([][]int, 2)
		memo[i][0] = make([]int, k+1)
		memo[i][1] = make([]int, k+1)
		for j := 0; j <= k; j++ {
			memo[i][0][j] = -1
			memo[i][1][j] = -1
		}
	}
	return recur(prices, 0, k, 1, &memo)
}

func recur(arr []int, ind int, k int, buy int, memo *[][][]int) int {
	if ind == len(arr) || k == 0 {
		return 0
	}
	if (*memo)[ind][buy][k] != -1 {
		return (*memo)[ind][buy][k]
	}
	p := math.MinInt64
	if buy == 1 {
		take := recur(arr, ind+1, k, 0, memo) - arr[ind]
		dont_take := recur(arr, ind+1, k, 1, memo)
		p = max(take, dont_take)
	} else {
		sell := recur(arr, ind+1, k-1, 1, memo) + arr[ind]
		dont_sell := recur(arr, ind+1, k, 0, memo)
		p = max(sell, dont_sell)
	}
	(*memo)[ind][buy][k] = p
	return p
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//-------------------------memoization-----------------------------

func main() {
	fmt.Println(best_time_to_buy_and_sell_stock_four(2, []int{2, 4, 1}))
}
