package main

import (
	"fmt"
	"sort"
)

//--------------------brute force-------------------------
func remove_covered_intervals_brute(arr [][]int) int {
	n := len(arr)
	cnt := 0
	mp := make(map[int]bool)
	for i := 0; i < n; i++ {
		curr := arr[i]
		for j := 0; j < n; j++ {
			if i != j && !mp[j] {
				if curr[0] <= arr[j][0] && curr[1] >= arr[j][1] {
					mp[j] = true
					cnt++
				}
			}
		}
	}
	return n - cnt
}

//--------------------brute force-------------------------

//--------------------sorting-------------------------
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

//--------------------sorting-------------------------

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
