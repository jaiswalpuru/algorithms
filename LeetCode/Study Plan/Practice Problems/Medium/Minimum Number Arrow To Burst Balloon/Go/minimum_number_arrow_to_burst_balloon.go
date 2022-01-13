package main

import (
	"fmt"
	"sort"
)

type Item [][]int

func (vec Item) Len() int           { return len(vec) }
func (vec Item) Less(i, j int) bool { return vec[i][0] < vec[j][0] }
func (vec Item) Swap(i, j int)      { vec[i], vec[j] = vec[j], vec[i] }

func min_arrows(balloon Item) int {
	n := balloon.Len()
	if n == 1 {
		return 1
	}
	sort.Sort(balloon)

	max_end := balloon[0][1]
	min_arrow := 1
	for i := 1; i < n; i++ {
		if balloon[i][0] <= max_end {
			max_end = min(max_end, balloon[i][1])
		} else {
			min_arrow++
			max_end = balloon[i][1]
		}
	}
	return min_arrow
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	fmt.Println(min_arrows([][]int{{10, 16}, {2, 8}, {1, 6}, {7, 12}}))
}
