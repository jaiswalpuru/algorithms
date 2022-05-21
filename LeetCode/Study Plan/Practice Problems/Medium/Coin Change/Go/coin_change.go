package main

import "fmt"
import "math"

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

//--------------------Brute force recursion---------------------------
func coin_change_brute(coins []int, amt int) int {
	l := math.MaxInt64
	_coin_change(coins, amt, []int{}, &l, 0)
	if l == math.MaxInt64 {
		return -1
	}
	return l
}

func _coin_change(coins []int, amt int, set []int, l *int, ind int) {

	if amt < 0 || ind >= len(coins) {
		return
	}

	if amt == 0 {
		*l = min(*l, len(set))
		return
	}

	if amt >= coins[ind] {
		temp := append(set, coins[ind])
		_coin_change(coins, amt-coins[ind], temp, l, ind)
		temp = temp[:len(temp)-1]
	}
	_coin_change(coins, amt, set, l, ind+1)
}

//--------------------Brute force recursion---------------------------

//---------------------------Memo recursion---------------------------
func coin_change(coins []int, amt int) int {
	if amt < 1 {
		return 0
	}
	memo := make([]int, amt+1)
	return _coin_change_memo(coins, amt, &memo)
}

func _coin_change_memo(coins []int, amt int, memo *[]int) int {
	if amt < 0 {
		return -1
	}

	if amt == 0 {
		return 0
	}

	if (*memo)[amt-1] != 0 {
		return (*memo)[amt-1]
	}

	min_val := math.MaxInt64
	for i := 0; i < len(coins); i++ {
		take := _coin_change_memo(coins, amt-coins[i], memo)
		if take >= 0 && take < min_val {
			min_val = take + 1
		}
	}

	if min_val == math.MaxInt64 {
		(*memo)[amt-1] = -1
	} else {
		(*memo)[amt-1] = min_val
	}
	return (*memo)[amt-1]
}

//---------------------------Memo recursion---------------------------

func main() {
	fmt.Println(coin_change([]int{1, 2, 5}, 11))
}
