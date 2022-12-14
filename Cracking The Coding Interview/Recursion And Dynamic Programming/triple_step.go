package main

import "fmt"

func triple_step(n int) int {
	if n == 0 || n == 1 || n == 2 {
		return n
	}
	if n == 3 {
		return 4
	}
	return triple_step(n-1) + triple_step(n-2) + triple_step(n-3)
}

func triple_step_eff(n int) int {
	memo := make(map[int]int)
	memo[0] = 0
	memo[1] = 1
	memo[2] = 2
	memo[3] = 4
	return _memo(n, &memo)
}

func _memo(n int, memo *map[int]int) int {
	if n == 0 || n == 1 || n == 2 {
		return n
	}
	if n == 3 {
		return 4
	}
	if v, ok := (*memo)[n]; ok {
		return v
	}
	(*memo)[n] = _memo(n-1, memo) + _memo(n-2, memo) + _memo(n-3, memo)
	return (*memo)[n]
}

func main() {
	fmt.Println(triple_step_eff(4))
}
