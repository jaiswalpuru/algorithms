package main

import (
	"fmt"
	"sort"
)

func min_arrows(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		if points[i][0] == points[j][0] {
			return points[i][1] < points[j][1]
		}
		return points[i][0] < points[j][0]
	})
	min_arrows := 1
	curr := points[0]
	for i := 1; i < len(points); i++ {
		if curr[1] >= points[i][0] {
			if curr[1] > points[i][1] {
				curr[1] = points[i][1]
			}
			continue
		}
		curr = points[i]
		min_arrows++
	}
	return min_arrows
}

func main() {
	fmt.Println(min_arrows([][]int{{10, 16}, {2, 8}, {1, 6}, {7, 12}}))
}
