package main

import (
	"fmt"
)

func maxProfit(prices []int) int {
	size := len(prices)
	buy := 0
	profit := 0

	for i := 0; i < size; i++ {
		if prices[i] < prices[buy] {
			buy = i
		}
		if prices[i] > prices[buy] {
			profit = max(profit, prices[i]-prices[buy])
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
