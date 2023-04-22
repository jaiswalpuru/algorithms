package main

import "fmt"

//----------------------------brute force------------------------------
func best_time_to_buy_and_sell_stock_with_cooldown(prices []int) int {
	return recur(prices, 0, true)
}

func recur(price []int, ind int, buy bool) int {
	if ind >= len(price) {
		return 0
	}
	if buy {
		b := recur(price, ind+1, false) - price[ind]
		db := recur(price, ind+1, true)
		return max(b, db)
	} else {
		s := recur(price, ind+2, true) + price[ind]
		ds := recur(price, ind+1, false)
		return max(s, ds)
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//----------------------------brute force------------------------------

//----------------------------memo------------------------------
func best_time_to_buy_and_sell_stock_with_cooldown_memo(price []int) int {
	memo := make([][]int, len(price))
	for i := 0; i < len(price); i++ {
		memo[i] = make([]int, 2)
		memo[i][0] = -1
		memo[i][1] = -1
	}
	return _recur(price, 0, 1, &memo)
}

func _recur(price []int, ind, buy int, memo *[][]int) int {
	if ind >= len(price) {
		return 0
	}
	if (*memo)[ind][buy] != -1 {
		return (*memo)[ind][buy]
	}
	profit := 0
	if buy == 1 {
		b := _recur(price, ind+1, 0, memo) - price[ind]
		db := _recur(price, ind+1, 1, memo)
		profit = max(b, db)
	} else {
		s := _recur(price, ind+2, 1, memo) + price[ind]
		ds := _recur(price, ind+1, 0, memo)
		profit = max(s, ds)
	}
	(*memo)[ind][buy] = profit
	return (*memo)[ind][buy]
}

//----------------------------memo------------------------------

func main() {
	fmt.Println(best_time_to_buy_and_sell_stock_with_cooldown_memo([]int{1, 2, 3, 0, 2}))
}
