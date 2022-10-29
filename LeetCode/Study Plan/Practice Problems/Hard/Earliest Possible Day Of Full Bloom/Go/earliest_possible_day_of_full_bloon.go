package main

import (
	"fmt"
	"sort"
)

func earliest_possible_day_of_full_bloon(plant, grow []int) int {
	n := len(plant)
	arr := make([][]int, n)
	for i := 0; i < n; i++ {
		arr[i] = make([]int, 2)
		arr[i][0] = plant[i]
		arr[i][1] = grow[i]
	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i][1] > arr[j][1]
	})
	plant_sum := 0
	max_val := 0
	for i := 0; i < n; i++ {
		plant_sum += arr[i][0]
		max_val = max(max_val, plant_sum+arr[i][1])
	}
	return max_val
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(earliest_possible_day_of_full_bloon([]int{1, 4, 3}, []int{2, 3, 1}))
}
