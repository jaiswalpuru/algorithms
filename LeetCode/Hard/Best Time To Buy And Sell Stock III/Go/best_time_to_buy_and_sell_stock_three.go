package main

import (
	"fmt"
	"math"
)

//---------------------------recursion---------------------------
func best_time_to_buy_and_sell_stock_three(arr []int) int {
	return recur(arr, 0, 2, true)
}

func recur(arr []int, ind int, k int, buy bool) int {
	if ind == len(arr) || k == 0 {
		return 0
	}
	p := math.MinInt64
	if buy {
		take := recur(arr, ind+1, k, !buy) - arr[ind]
		dont_take := recur(arr, ind+1, k, buy)
		p = max(take, dont_take)
	} else {
		sell := recur(arr, ind+1, k-1, !buy) + arr[ind]
		dont_sell := recur(arr, ind+1, k, buy)
		p = max(sell, dont_sell)
	}
	return p
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//---------------------------recursion---------------------------

//---------------------------memoization---------------------------
func best_time_to_buy_and_sell_stock_three_memo(arr []int) int {
	memo := make([][][]int, len(arr))
	for i := 0; i < len(memo); i++ {
		memo[i] = make([][]int, 2)
		memo[i][0] = make([]int, 3)
		memo[i][1] = make([]int, 3)
		for j := 0; j < 3; j++ {
			memo[i][0][j] = -1
			memo[i][1][j] = -1
		}
	}
	return _memo(arr, 0, 2, 1, &memo)
}

func _memo(arr []int, ind int, k int, buy int, memo *[][][]int) int {
	if ind == len(arr) || k == 0 {
		return 0
	}
	if (*memo)[ind][buy][k] != -1 {
		return (*memo)[ind][buy][k]
	}
	p := math.MinInt64
	if buy == 1 {
		take := _memo(arr, ind+1, k, 0, memo) - arr[ind]
		dont_take := _memo(arr, ind+1, k, 1, memo)
		p = max(take, dont_take)
	} else {
		sell := _memo(arr, ind+1, k-1, 1, memo) + arr[ind]
		dont_sell := _memo(arr, ind+1, k, 0, memo)
		p = max(sell, dont_sell)
	}
	(*memo)[ind][buy][k] = p
	return p
}

//---------------------------memoization---------------------------

func main() {
	fmt.Println(best_time_to_buy_and_sell_stock_three_memo([]int{3, 3, 5, 0, 0, 3, 1, 4}))
}
