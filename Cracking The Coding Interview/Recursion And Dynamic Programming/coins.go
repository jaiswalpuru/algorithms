package main

import (
	"fmt"
)

//-------------------brute force------------------
func coins(amt int, change []int) int {
	return recur(amt, change, 0)
}

func recur(amt int, change []int, ind int) int {
	if ind >= len(change) || amt < 0 {
		return 0
	}
	if amt == 0 {
		return 1
	}
	ways := 0
	for i := 0; i <= amt; i += change[ind] {
		ways += recur(amt-i, change, ind+1)
	}
	return ways
}

//-------------------brute force------------------

//-------------------optimized------------------
func coins_memo(amt int, change []int) int {
	memo := make([][]int, amt+1)
	for i := 0; i <= amt; i++ {
		memo[i] = make([]int, len(change))
	}
	return _memo(amt, change, 0, &memo)
}

func _memo(amt int, change []int, ind int, memo *[][]int) int {
	if ind >= len(change) || amt < 0 {
		return 0
	}
	if amt == 0 {
		return 1
	}
	if (*memo)[amt][ind] > 0 {
		return (*memo)[amt][ind]
	}
	ways := 0
	for i := 0; i <= amt; i += change[ind] {
		ways += _memo(amt-i, change, ind+1, memo)
	}
	(*memo)[amt][ind] = ways
	return (*memo)[amt][ind]
}

//-------------------optimized------------------

func main() {
	fmt.Println(coins_memo(10, []int{25, 10, 5, 1}))
}
