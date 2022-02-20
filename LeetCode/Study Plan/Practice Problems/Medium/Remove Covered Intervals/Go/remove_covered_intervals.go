package main

import (
	"fmt"
	"sort"
)

func remove_covered_intervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] > intervals[j][1]
		}
		return intervals[i][0] < intervals[j][0]
	})
	cnt := 0
	n := len(intervals)
	end, prev := 0, 0
	for i := 0; i < n; i++ {
		end = intervals[i][1]
		if prev < end {
			cnt++
			prev = end
		}
	}
	return cnt
}

func main() {
	fmt.Println(remove_covered_intervals([][]int{
		{3, 10}, {4, 10}, {5, 11},
	}))
	fmt.Println(remove_covered_intervals([][]int{
		{1, 2}, {1, 4}, {3, 4},
	}))
	fmt.Println(remove_covered_intervals([][]int{
		{2, 3}, {1, 4},
	}))
}
