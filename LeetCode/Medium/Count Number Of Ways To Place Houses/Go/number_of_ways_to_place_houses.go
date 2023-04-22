package main

import "fmt"

var mod = int(1e9 + 7)

//------------------------brute force-----------------------
func count_number_of_ways_to_place_houses(n int) int {
	ans := _recur(n, 0) + _recur(n, 1)
	return (ans * ans) % mod
}

func _recur(n int, can_place int) int {
	if n == 1 {
		return 1
	}

	tot_ways := 0
	if can_place == 0 {
		tot_ways = _recur(n-1, 1) % mod
	}
	tot_ways += _recur(n-1, 0) % mod

	return tot_ways
}

//------------------------brute force-----------------------

//------------------------memoization-----------------------
func count_number_of_ways_to_place_houses_eff(n int) int {
	memo := make([][2]int, n+1)
	for i := 0; i <= n; i++ {
		memo[i][0] = -1
		memo[i][1] = -1
	}
	res := _memo(n, 0, &memo) + _memo(n, 1, &memo)
	return (res * res) % mod
}

func _memo(n, can_place int, memo *[][2]int) int {
	if n == 1 {
		return 1
	}
	if (*memo)[n][can_place] != -1 {
		return (*memo)[n][can_place]
	}
	tot_ways := 0
	if can_place == 0 {
		tot_ways = _memo(n-1, 1, memo)
	}
	tot_ways += _memo(n-1, 0, memo) % mod
	(*memo)[n][can_place] = tot_ways
	return (*memo)[n][can_place]
}

//------------------------memoization-----------------------

func main() {
	fmt.Println(count_number_of_ways_to_place_houses_eff(1))
}
