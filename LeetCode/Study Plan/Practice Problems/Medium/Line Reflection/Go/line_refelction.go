package main

import (
	"fmt"
	"math"
)

func line_refelction(points [][]int) bool {
	reflection := make(map[[2]int]bool)
	x_max, x_min := math.MinInt64, math.MaxInt64
	for i := 0; i < len(points); i++ {
		x_max = max(x_max, points[i][0])
		x_min = min(x_min, points[i][0])
		reflection[[2]int{points[i][0], points[i][1]}] = true
	}
	sum := x_max + x_min
	for i := 0; i < len(points); i++ {
		if !reflection[[2]int{sum - points[i][0], points[i][1]}] {
			return false
		}
	}
	return true
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(line_refelction([][]int{
		{1, 1}, {-1, 1},
	}))
}
