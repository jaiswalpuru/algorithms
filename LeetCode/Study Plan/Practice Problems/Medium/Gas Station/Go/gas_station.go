package main

import "fmt"

func gas_station(gas []int, cost []int) int {
	curr_tank, total_tank, start := 0, 0, 0
	for i := 0; i < len(gas); i++ {
		total_tank += gas[i] - cost[i]
		curr_tank += gas[i] - cost[i]
		if curr_tank < 0 {
			start = i + 1
			curr_tank = 0
		}
	}
	if total_tank < 0 {
		return -1
	}
	return start
}

func main() {
	fmt.Println(gas_station([]int{1, 2, 3, 4, 5}, []int{3, 4, 5, 1, 2}))
}
