package main

import (
	"fmt"
	"math"
)

func buy_sell(arr []int) int {
	buy := 0
	sell := 0
	n := len(arr)
	profit := math.MinInt64
	for i := 1; i < n; i++ {
		if i > buy && arr[i] < arr[buy] {
			buy = i
		}

		if i > buy && arr[i] > arr[buy] {
			sell = i
		}

		if sell > buy {
			profit = max(profit, arr[sell]-arr[buy])
		}
	}

	return profit
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(buy_sell([]int{7, 1, 5, 3, 6, 4}))
}
