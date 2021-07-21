/*
You are given an integer array coins representing coins of different denominations and an
integer amount representing a total amount of money.

Return the fewest number of coins that you need to make up that amount.
If that amount of money cannot be made up by any combination of the coins, return -1.

You may assume that you have an infinite number of each kind of coin.

Example 1:
Input: coins = [1,2,5], amount = 11
Output: 3
Explanation: 11 = 5 + 5 + 1

Example 2:
Input: coins = [2], amount = 3
Output: -1

Example 3:
Input: coins = [1], amount = 0
Output: 0

Example 4:
Input: coins = [1], amount = 1
Output: 1

Example 5:
Input: coins = [1], amount = 2
Output: 2
*/

package main

import (
	"fmt"
	"math"
)

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func coin_change(coins []int, amt int) int {
	dp := make([]int, amt+1, amt+1)
	for i := 1; i <= amt; i++ {
		dp[i] = math.MaxInt32
		for _, v := range coins {
			if i-v >= 0 {
				dp[i] = min(dp[i-v]+1, dp[i])
			}
		}
	}
	//fmt.Println(dp)
	if dp[amt] == math.MaxInt32 {
		return -1
	}
	return dp[amt]
}

func main() {
	fmt.Println(coin_change([]int{1, 2, 5}, 11))
	fmt.Println(coin_change([]int{2}, 3))
	fmt.Println(coin_change([]int{1}, 1))
	fmt.Println(coin_change([]int{1}, 2))
	//fmt.Println(all_possile([]int{1, 2, 5}, 11))
}

/*
Some experiments were done here

func _all(arr []int, amt int, res *[][]int, temp []int, ind int) {

	if amt < 0 {
		return
	}

	if amt == 0 {
		*res = append(*res, append([]int{}, temp...))
		return
	}

	for i := ind; i < len(arr); i++ {
		temp = append(temp, arr[i])
		_all(arr, amt-arr[i], res, temp, i)
		temp = temp[:len(temp)-1]
	}
}

func all_possile(arr []int, amt int) [][]int {
	res := [][]int{}
	temp := []int{}
	_all(arr, amt, &res, temp, 0)
	return res
}
*/
