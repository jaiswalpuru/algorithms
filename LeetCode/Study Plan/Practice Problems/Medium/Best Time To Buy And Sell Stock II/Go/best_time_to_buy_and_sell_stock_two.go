package main

import (
	"fmt"
	"math"
)

//----------------------recursion(TLE)-------------------------
func best_time_to_buy_and_sell_stock_two(arr []int) int {
	return recur(arr, 0, true)
}

func recur(arr []int, ind int, buy bool) int {
	if ind == len(arr) {
		return 0
	}
	profit := 0
	if buy {
		buy := recur(arr, ind+1, false) - arr[ind]
		dont_buy := recur(arr, ind+1, true)
		profit = max(buy, dont_buy)
	} else {
		sell := recur(arr, ind+1, true) + arr[ind]
		dont_sell := recur(arr, ind+1, false)
		profit = max(sell, dont_sell)
	}
	return profit
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//----------------------recursion(TLE)-------------------------

//---------------------memoization----------------------------
func best_time_to_buy_and_sell_stock_two_memo(arr []int) int {
	memo := make([][]int, len(arr))
	for i := 0; i < len(arr); i++ {
		memo[i] = make([]int, 2)
		memo[i][0] = -1
		memo[i][1] = -1
	}
	return recur_memo(arr, 0, 1, &memo)
}

func recur_memo(arr []int, ind int, buy int, memo *[][]int) int {
	if ind == len(arr) {
		return 0
	}

	if (*memo)[ind][buy] != -1 {
		return (*memo)[ind][buy]
	}

	profit := math.MinInt64
	if buy == 1 {
		buy := recur_memo(arr, ind+1, 0, memo) - arr[ind]
		dont_buy := recur_memo(arr, ind+1, 1, memo)
		profit = max(buy, dont_buy)
	} else {
		sell := recur_memo(arr, ind+1, 1, memo) + arr[ind]
		dont_sell := recur_memo(arr, ind+1, 0, memo)
		profit = max(sell, dont_sell)
	}
	(*memo)[ind][buy] = profit
	return (*memo)[ind][buy]
}

//---------------------memoization----------------------------

//---------------------greedy----------------------------
func best_time_to_buy_and_sell_stock_two_greedy(arr []int) int {
	profit := 0
	buy := 0
	n := len(arr)
	for i := 0; i < n; i++ {
		if arr[i] > arr[buy] {
			profit += arr[i] - arr[buy]
			buy = i
		} else {
			buy = i
		}
	}
	return profit
}

//---------------------greedy----------------------------

func main() {
	fmt.Println(best_time_to_buy_and_sell_stock_two_greedy([]int{7, 1, 5, 3, 6, 4}))
}
