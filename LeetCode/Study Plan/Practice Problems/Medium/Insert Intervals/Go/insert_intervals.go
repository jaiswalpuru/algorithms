package main

import (
	"fmt"
	"sort"
)

func insert_intervals(interval [][]int, new_interval []int) [][]int {
	interval = append(interval, new_interval)
	sort.Slice(interval, func(i, j int) bool {
		if interval[i][0] == interval[j][0] {
			return interval[i][1] < interval[j][1]
		}
		return interval[i][0] < interval[j][0]
	})
	n := len(interval)
	res := [][]int{}
	curr := interval[0]
	for i := 1; i < n; i++ {
		if interval[i][0] > curr[1] {
			res = append(res, curr)
			curr = interval[i]
		} else {
			curr[1] = max(curr[1], interval[i][1])
		}
	}
	res = append(res, curr)
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(insert_intervals([][]int{
		{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16},
	}, []int{4, 8}))
}
