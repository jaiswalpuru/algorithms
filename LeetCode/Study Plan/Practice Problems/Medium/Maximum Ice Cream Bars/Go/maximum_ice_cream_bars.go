package main

import (
	"fmt"
	"sort"
)

func maximum_ice_cream_bars(costs []int, coins int) int {
	sort.Ints(costs)
	i := 0
	for i = 0; i < len(costs); i++ {
		if costs[i] <= coins {
			coins -= costs[i]
		} else {
			return i
		}
	}
	return i
}

func main() {
	fmt.Println(maximum_ice_cream_bars([]int{1, 3, 2, 4, 1}, 7))
}
