/*
A wall consists of several rows of bricks of various integer lengths and uniform height.
Your goal is to find a vertical line going from the top to the bottom of the edge between two bricks.
If the line goes through the edge between two bricks, this doesn't count as cut.

*/

package main

import (
	"fmt"
	"math"
)

func max_(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func max(mp map[int]int) int {
	val := int(math.MinInt32)
	for _, v := range mp {
		val = max_(val, v)
	}

	return val
}

func main() {

	bricks := [][]int{
		{3, 5, 1, 1},
		{2, 3, 3, 2},
		{5, 5},
		{4, 4, 2},
		{1, 3, 3, 3},
		{1, 1, 6, 1, 1},
	}
	mp := make(map[int]int)

	for i := 0; i < len(bricks); i++ {
		l := 0
		for j := 0; j < len(bricks[i])-1; j++ {
			l += bricks[i][j]
			mp[l]++
		}
	}
	fmt.Println(len(bricks) - max(mp))
}
