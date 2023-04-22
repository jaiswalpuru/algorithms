package main

import (
	"fmt"
	"sort"
)

//-------------------without sorting-----------------------
func insert_intervals_without_sort(intervals [][]int, new_interval []int) [][]int {
	interval_mod := [][]int{}
	res := [][]int{}
	f := true
	for i := 0; i < len(intervals); i++ {
		if intervals[i][1] > new_interval[0] && f {
			interval_mod = append(interval_mod, new_interval)
			interval_mod = append(interval_mod, intervals[i])
			f = false
		} else {
			interval_mod = append(interval_mod, intervals[i])
		}
	}
	if f {
		interval_mod = append(interval_mod, new_interval)
	}
	start := interval_mod[0]
	for i := 1; i < len(interval_mod); i++ {
		if start[1] >= interval_mod[i][0] {
			start[0] = min(start[0], interval_mod[i][0])
			start[1] = max(start[1], interval_mod[i][1])
		} else {
			res = append(res, start)
			start = interval_mod[i]
		}
	}
	res = append(res, start)
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

//-------------------without sorting-----------------------

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
