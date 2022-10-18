package main

import (
	"fmt"
	"sort"
)

func merge_intervals(interval [][]int) [][]int {
	sort.Slice(interval, func(i, j int) bool {
		if interval[i][0] == interval[j][1] {
			return interval[i][1] < interval[j][1]
		}
		return interval[i][0] < interval[j][0]
	})
	n := len(interval)
	res := [][]int{}
	curr := interval[0]
	for i := 1; i < n; i++ {
		if interval[i][0] <= curr[1] {
			curr[1] = max(curr[1], interval[i][1])
			continue
		} else {
			res = append(res, curr)
			curr = interval[i]
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
	fmt.Println(merge_intervals([][]int{
		{1, 3}, {2, 6}, {8, 10}, {15, 18},
	}))
}
