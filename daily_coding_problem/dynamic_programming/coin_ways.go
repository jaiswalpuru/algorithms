/*
Number of ways it is possible to lay pennies and nickels in a line on a table such that they sum to a dollar
*/

package main

import "fmt"

var (
	cache = make(map[int]int)
)

//memoization
func coins_top_down(n int) int {

	if _, ok := cache[n]; ok {
		return cache[n]
	}

	if n < 0 {
		return 0
	}

	cache[n] = coins_top_down(n-1) + coins_top_down(n-5)
	return cache[n]
}

func coin_bottom_up(n int) int {
	in_mem := make(map[int]int)
	in_mem[0] = 1

	for i := 1; i < n+1; i++ {
		if _, ok := in_mem[i-1]; ok {
			in_mem[i] += in_mem[i-1]
		}
		if _, ok := in_mem[i-5]; ok {
			in_mem[i] += in_mem[i-5]
		}
	}

	return in_mem[n]
}

func main() {
	cache[0] = 1
	fmt.Println(coins_top_down(100))
	fmt.Println(coin_bottom_up(100))
}
