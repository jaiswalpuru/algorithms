package main

import (
	"fmt"
	"math"
	"sort"
)

//--------------------recursion----------------------------
func non_overlapping_intervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] < intervals[j][1]
		}
		return intervals[i][0] < intervals[j][0]
	})
	return recur(intervals, 0, 1)
}

func recur(intervals [][]int, i, j int) int {
	if j == len(intervals) {
		return 0
	}
	t1, t2 := 0, 0
	if intervals[i][1] > intervals[j][0] {
		choose := recur(intervals, i, j+1)
		dont_choose := recur(intervals, j, j+1)
		t1 = 1 + min(choose, dont_choose)
	} else {
		t2 = recur(intervals, j, j+1)
	}
	return t1 + t2
}

//--------------------recursion----------------------------

//--------------------memoization----------------------------
func non_overlapping_intervals_memo(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] < intervals[j][1]
		}
		return intervals[i][0] < intervals[j][0]
	})
	n := len(intervals)
	memo := make([]int, n)
	for i := 0; i < n; i++ {
		memo[i] = -1
	}
	return _memo(intervals, 0, &memo)
}

func _memo(intervals [][]int, start int, memo *[]int) int {
	if start == len(intervals) {
		return 0
	}
	if (*memo)[start] != -1 {
		return (*memo)[start]
	}
	ind := start + 1
	max_val := intervals[start][1]
	for ind < len(intervals) && intervals[ind][0] < max_val {
		ind++
	}
	take := math.MaxInt64
	//skip all the overlapping intervals
	take = (ind - start - 1) + _memo(intervals, ind, memo)
	//skip only the current overlapping interval
	if ind > start+1 {
		take = min(take, _memo(intervals, start+1, memo)+1)
	}
	(*memo)[start] = take
	return (*memo)[start]
}

//--------------------memoization----------------------------

//--------------------memoization greedy----------------------------
func non_overlapping_intervals_greedy(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool { return intervals[i][1] < intervals[j][1] })
	prev := 0
	cnt := 0
	for i := 1; i < len(intervals); i++ {
		if intervals[prev][1] > intervals[i][0] {
			if intervals[prev][1] > intervals[i][1] {
				prev = i
			}
			cnt++
		} else {
			prev = i
		}
	}
	return cnt
}

//--------------------memoization greedy----------------------------

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	fmt.Println(non_overlapping_intervals_greedy([][]int{
		{1, 2}, {2, 3}, {3, 4}, {1, 3},
	}))
}
