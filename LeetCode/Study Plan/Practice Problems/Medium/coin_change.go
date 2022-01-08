/*
You are given an integer array coins representing coins of different denominations and an integer amount
representing a total amount of money.

Return the fewest number of coins that you need to make up that amount. If that amount of money cannot be
made up by any combination of the coins, return -1.

You may assume that you have an infinite number of each kind of coin
*/

package main

import (
	"fmt"
	"math"
)

//memoization will give tle
func coin_change(arr []int, amt int) int {
	return _coin_change(0, arr, amt)
}

func _coin_change(idx int, arr []int, amt int) int {
	if amt == 0 {
		return 0
	}
	if idx < len(arr) && amt > 0 {
		max_val := amt / arr[idx]
		min_cost := int(math.MaxInt64)
		for i := 0; i <= max_val; i++ {
			res := _coin_change(idx+1, arr, amt-i*arr[idx])
			if res != -1 {
				min_cost = min(min_cost, res+i)
			}
		}
		if min_cost == int(math.MaxInt64) {
			return -1
		}
		return min_cost
	}
	return -1
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

//dp
func coin_change_dp(arr []int, amt int) int {
	if amt < 0 {
		return -1
	}
	count := make([]int, amt)
	return _coin_change_dp(amt, arr, count)
}

func _coin_change_dp(amt int, arr, count []int) int {
	if amt < 0 {
		return -1
	}
	if amt == 0 {
		return 0
	}
	if count[amt-1] != 0 {
		return count[amt-1]
	}

	min_val := int(math.MaxInt64)
	for i := 0; i < len(arr); i++ {
		res := _coin_change_dp(amt-arr[i], arr, count)
		if res >= 0 && res < min_val {
			min_val = res + 1
		}
	}
	if min_val == int(math.MaxInt64) {
		count[amt-1] = -1
	} else {
		count[amt-1] = min_val
	}
	return count[amt-1]
}

func main() {
	fmt.Println(coin_change_dp([]int{1, 2, 5}, 11))
}
