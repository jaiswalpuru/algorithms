package main

import (
	"fmt"
	"sort"
)

func smallest_range_one(arr []int, k int) int {
	sort.Ints(arr)
	min_val := arr[0] + k
	max_val := arr[len(arr)-1] - k
	return max(0, max_val-min_val)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(smallest_range_one([]int{1}, 0))
}
