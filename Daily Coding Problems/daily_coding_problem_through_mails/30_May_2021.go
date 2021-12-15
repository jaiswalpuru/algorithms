/*
	Given a array of numbers representing the stock prices of a company in chronological order,
	write a function that calculates the maximum profit you could have made from buying and selling that stock once.
	You must buy before you can sell it.

	For example, given [9, 11, 8, 5, 7, 10], you should return 5, since you could buy the stock at 5 dollars
	and sell it at 10 dollars.
*/

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

func max_profit(arr []int) int {
	max_profit := 0
	buy := 0
	for i := 1; i < len(arr); i++ {
		if arr[i] > arr[buy] {
			max_profit = max(max_profit, arr[i]-arr[buy])
		}
		if arr[buy] > arr[i] {
			buy = i
		}
	}
	return max_profit
}

func main() {
	arr := []int{9, 11, 8, 5, 7, 10}
	fmt.Println(max_profit(arr))
}
